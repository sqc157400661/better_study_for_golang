package main

import "errors"

func main(){
	test()
}

func test() error{
	return errors.New("文件名错误")
}