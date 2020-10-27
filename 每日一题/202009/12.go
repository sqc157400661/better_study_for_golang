
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//defer func() {
	//	if e:=recover();e !=nil{
	//		fmt.Println(e)
	//	}
	//}()
	go get()
	go put()
	time.Sleep(time.Second * 10)
	fmt.Println(111111)
}

func get(){
	defer func() {
		if e:=recover();e !=nil{
			//fmt.Println(e)
			var buf [4096]byte
			n := runtime.Stack(buf[:],false)
			fmt.Println(e,string(buf[:n]))
			//const size = 64 << 10
			//buf := make([]byte, size)
			//buf = buf[:runtime.Stack(buf, false)]
			//fmt.Println(string(buf))
		}
	}()
	panic(111)
}

func put(){
	fmt.Println(3333)
}