package main

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
)

func main() {
	//http://127.0.0.1:8000/demo
	// 单独写回调函数
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
type UserData struct{
	Name string
	Age int
}
type Ret struct{
	Code int
	Msg string
	Data []UserData
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	cangJingKong := UserData{Name:"CangLaoShi",Age:35}
	ret := Ret{
		Code:200,
		Msg:"success",
		Data:[]UserData{cangJingKong},
	}
	jsonData,_ := json.Marshal(ret)
	go putFileContents("ok\n")
	// 回复
	w.Write(jsonData)
}
var Mutex sync.Mutex

func putFileContents(str string){
	var File,_ = os.OpenFile("/website/goreplay/godemo.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	Mutex.Lock()
	defer File.Close()
	defer Mutex.Unlock()
	File.WriteString(str) //写入文件(字节数组)
	File.Sync()
}


