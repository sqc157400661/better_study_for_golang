package main

import (
	"fmt"
	"reflect"
)

func main(){
	buf := "test"
	aa :=reflect.ValueOf(buf)
	aa.Call()
}

func test(){
	fmt.Println(1111)
}