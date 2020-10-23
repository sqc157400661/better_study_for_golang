package main

import (
	"encoding/json"
	"net/http"
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
	// 回复
	w.Write(jsonData)
}


