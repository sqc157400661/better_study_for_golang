package main

import (
	"fmt"
	"time"
)

func Process(ch chan int){
	// Do some work....
	time.Sleep(time.Second)
	ch <- 1  //管道中写入一个元素表示当前协程已结束
}

func main(){
	//创建一个10个元素的切片，元素类型为channel
	channels := make([]chan int,10)

	for i :=0;i<10;i++{
		//切片中放入一个channel
		channels[i] = make(chan int)
		//启动协程，传一个管道用于通信
		go Process(channels[i])
	}
	for i,ch := range channels {
		//遍历切片，等待子协程结束
		<-ch
		fmt.Println("Routine ", i, " quit!")
	}
}