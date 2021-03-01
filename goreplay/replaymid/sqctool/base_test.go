package sqctool

import (
	"encoding/hex"
	"reflect"
	"testing"
)

var passby = make(map[string]int)

func init() {
	passby["counter"] = 0
}

func incrCounter(gor *sqctool.Gor, msg *sqctool.GorMessage, kwargs ...interface{}) *sqctool.GorMessage {
	passbyReadonly, _ := kwargs[0].(map[string]int)
	increase, _ := kwargs[1].(int)
	passby["counter"] += passbyReadonly["counter"] + increase
	return msg
}

func TestMessageLogic(t *testing.T) {
	gor := sqctool.CreateGor()
	gor.On("message", incrCounter, "", &passby, 1)
	gor.On("request", incrCounter, "", &passby, 2)
	gor.On("response", incrCounter, "2", &passby, 3)
	if len(gor.retainQueue) != 2 {
		t.Errorf("gor retain queue length %d != 2", len(gor.retainQueue))
	}
	if len(gor.tempQueue) != 1 {
		t.Errorf("gor temp queue length %d != 1", len(gor.tempQueue))
	}
	req, err := gor.ParseMessage(hex.EncodeToString([]byte("1 2 3\nGET / HTTP/1.1\r\n\r\n")))
	if err != nil {
		t.Error(err.Error())
	}
	resp, err := gor.ParseMessage(hex.EncodeToString([]byte("2 2 3\nHTTP/1.1 200 OK\r\n\r\n")))
	if err != nil {
		t.Error(err.Error())
	}
	resp2, err := gor.ParseMessage(hex.EncodeToString([]byte("2 3 3\nHTTP/1.1 200 OK\r\n\r\n")))
	if err != nil {
		t.Error(err.Error())
	}
	gor.Emit(req)
	gor.Emit(resp)
	gor.Emit(resp2)
	if passby["counter"] != 8 {
		t.Errorf("passby counter %d != 8", passby["counter"])
	}
}

func TestParseMessage(t *testing.T) {
	gor := sqctool.CreateGor()
	payload := hex.EncodeToString([]byte("1 2 3\nGET / HTTP/1.1\r\n\r\n"))
	if msg, err := gor.ParseMessage(payload); err != nil {
		t.Error(err.Error())
	} else {
		if msg.ID != "2" {
			t.Errorf("invalid msg id %s != 2", msg.ID)
		}
		if msg.Type != "1" {
			t.Errorf("invalid msg type %s != 1", msg.Type)
		}
		meta := [][]byte{[]byte("1"), []byte("2"), []byte("3")}
		if !reflect.DeepEqual(msg.Meta, meta) {
			t.Errorf("invalid msg meta")
		}
		if !reflect.DeepEqual(msg.HTTP, []byte("GET / HTTP/1.1\r\n\r\n")) {
			t.Errorf("invalid msg http")
		}
	}
}
