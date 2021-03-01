package sqctool

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

// GorMessage stores data and parsed information in a incoming request
type GorMessage struct {
	ID      string
	Type    string
	Meta    [][]byte // Meta is an array size of 4, containing: request type, uuid, timestamp, latency
	RawMeta []byte   //
	HTTP    []byte   // Raw HTTP payload
}

// InterFunc stores event callback function, args and create time
type InterFunc struct {
	created time.Time
	fn      func(*Gor, *GorMessage, ...interface{}) *GorMessage
	args    []interface{}
}

// Gor is the middleware itself
type Gor struct {
	retainQueue map[string]([]*InterFunc) // retainQueue stores immutable registered event, such as "request", "response" etc
	tempQueue   map[string]([]*InterFunc) // tempQueue stores mutable event associated with request ID
	lock        *sync.RWMutex
	input       chan string
	parsed      chan *GorMessage
	stderr      io.Writer
}

// CreateGor creates a Gor object
func CreateGor() *Gor {
	gor := &Gor{
		retainQueue: make(map[string]([]*InterFunc)),
		tempQueue:   make(map[string]([]*InterFunc)),
		lock:        new(sync.RWMutex),
		input:       make(chan string),
		parsed:      make(chan *GorMessage),
		stderr:      os.Stderr,
	}
	return gor
}

// On func 注册事件和回调函数
func (gor *Gor) On(
	channel string, fn func(*Gor, *GorMessage, ...interface{}) *GorMessage,
	idx string, args ...interface{}) {

	if idx != "" {
		channel = channel + "#" + idx
	}
	inmsg := &InterFunc{
		created: time.Now(),
		fn:      fn,
		args:    args,
	}
	gor.lock.Lock()
	if idx != "" {
		if c, ok := gor.tempQueue[channel]; ok {
			c = append(c, inmsg)
		} else {
			newChan := make([]*InterFunc, 0)
			newChan = append(newChan, inmsg)
			gor.tempQueue[channel] = newChan
		}
	} else {
		if c, ok := gor.retainQueue[channel]; ok {
			c = append(c, inmsg)
		} else {
			newChan := make([]*InterFunc, 0)
			newChan = append(newChan, inmsg)
			gor.retainQueue[channel] = newChan
		}
	}
	gor.lock.Unlock()
}

// Emit triggers the registerd event callback when receiving certain GorMessage
func (gor *Gor) Emit(msg *GorMessage) error {
	chanPrefix, ok := ChanPrefixMap[msg.Type]
	if !ok {
		return fmt.Errorf("invalid message type: %s", msg.Type)
	}
	chanIDs := [2]string{"message", chanPrefix}
	resp := msg
	for _, chanID := range chanIDs {
		if funcs, ok := gor.retainQueue[chanID]; ok {
			for _, f := range funcs {
				r := f.fn(gor, msg, f.args...)
				if r != nil {
					resp = r
				}
			}
		}
	}

	// lazy remove registered events in gor.cleanOldChannel goroutine
	tempChanID := fmt.Sprintf("%s#%s", chanPrefix, msg.ID)
	if funcs, ok := gor.tempQueue[tempChanID]; ok {
		var f *InterFunc
		tmp := make([]*InterFunc, 0)
		gor.lock.Lock()
		tmp, funcs = funcs[:], funcs[len(funcs):]
		gor.lock.Unlock()
		for len(tmp) > 0 {
			f, tmp = tmp[0], tmp[1:]
			r := f.fn(gor, msg, f.args...)
			if r != nil {
				resp = r
			}
		}
	}
	if resp != nil {
		fmt.Printf("%s", gor.HexData(resp))
	}
	return nil
}

// HexData 解析数据到GorMessage
func (gor *Gor) HexData(msg *GorMessage) string {
	encodeList := [3][]byte{msg.RawMeta, []byte("\n"), msg.HTTP}
	encodedList := make([]string, 3)
	for i, val := range encodeList {
		encodedList[i] = hex.EncodeToString(val)
	}
	encodedList = append(encodedList, "\n")
	return strings.Join(encodedList, "")
}

// ParseMessage parses string middleware dataflow into a GorMessage
func (gor *Gor) ParseMessage(line string) (*GorMessage, error) {
	payload, err := hex.DecodeString(strings.TrimSpace(line))
	if err != nil {
		return nil, err
	}


	metaPos := bytes.Index(payload, []byte("\n"))
	metaRaw := payload[:metaPos]
	metaArr := bytes.Split(metaRaw, []byte(" "))
	ptype, pid := metaArr[0], string(metaArr[1])
	if err != nil {
		return nil, err
	}
	httpPayload := payload[metaPos+1:]

	return &GorMessage{
		ID:      pid,
		Type:    string(ptype),
		Meta:    metaArr,
		RawMeta: metaRaw,
		HTTP:    httpPayload,
	}, nil
}

func (gor *Gor) cleanOldChannel(interval int) {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		gor.lock.Lock()
		for channel, funcs := range gor.tempQueue {
			for i := len(funcs) - 1; i >= 0; i-- {
				if time.Since(funcs[i].created) > time.Duration(interval) {
					funcs = append(funcs[:i], funcs[i+1:]...)
				}
			}
			if len(funcs) == 0 {
				delete(gor.tempQueue, channel)
			}
		}
		gor.lock.Unlock()
	}
}

func (gor *Gor) preProcessor() {
	for {
		line := <-gor.input
		if msg, err := gor.ParseMessage(line); err != nil {
			gor.stderr.Write([]byte(err.Error()))
		} else {
			gor.parsed <- msg
		}
	}
}

func (gor *Gor) receiver() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintf(os.Stderr, "%s\n", scanner.Text())
		gor.input <- scanner.Text()
	}
}

func (gor *Gor) processor() {
	for {
		msg := <-gor.parsed
		gor.Emit(msg)
	}
}

// 根据信号结束服务 todo
func (gor *Gor) shutdown() {
}

func (gor *Gor) handleSignal(sigChan chan os.Signal) {
	for {
		s := <-sigChan
		gor.stderr.Write([]byte(fmt.Sprintf("receive a signal %s\n", s.String())))
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			gor.shutdown()
			return
		default:
			return
		}
	}
}

// Run is entrypoint of Gor
func (gor *Gor) Run() {
	go gor.receiver()
	go gor.preProcessor()
	go gor.processor()
	go gor.cleanOldChannel(60)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM,
		syscall.SIGINT, syscall.SIGSTOP)
	gor.handleSignal(c)
}
