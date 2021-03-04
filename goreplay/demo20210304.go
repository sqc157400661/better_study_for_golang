package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)


const (
	MaxIdleConns        int = 300
	MaxIdleConnsPerHost int = 200
	IdleConnTimeout     int = 300 //keepalive 5 min
)

var httpClient = &http.Client{
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   3000 * time.Millisecond,
			KeepAlive: 300 * time.Second,
		}).DialContext,
		MaxIdleConns:        MaxIdleConns,
		MaxIdleConnsPerHost: MaxIdleConnsPerHost,
		IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
	},
	Timeout: 3000 * time.Millisecond,
}
func main() {
	//http://127.0.0.1:8000/demo
	// 单独写回调函数
	http.HandleFunc("/test", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	req,errRe := http.NewRequest(http.MethodGet,"http://127.0.0.1:8001/",strings.NewReader(""))
	if errRe != nil{
		return
	}
	resp,errResp:= httpClient.Do(req)
	if errResp !=nil{
		var body []byte
		if resp !=nil{
			body,_ = ioutil.ReadAll(resp.Body)
		}
		str := string(body) +"\n" +errResp.Error() + " \n"
		go putFileContents(str)
	}

	// 回复
	//w.Write([]byte("111"))
}

var Mutex sync.Mutex

func putFileContents(str string){
	var File,_ = os.OpenFile("/website/godemo20210304.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	Mutex.Lock()
	defer File.Close()
	defer Mutex.Unlock()
	File.WriteString(str) //写入文件(字节数组)
	File.Sync()
}


