/**
	学习通道- 通道阻塞
 */
package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	//ch1 <- 111 // 没有接收操作，这里将会死锁
	//<- ch1 // 没有写操作 这里将会死锁

	go push(ch1)
	fmt.Println("send")
	fmt.Println(<-ch1)
	fmt.Println("sended")

	//go get(ch1)
	//time.Sleep(1e9) // 防止main函数终止,main终止，所有的协程也终止
}

func push(ch chan int) {
	fmt.Println("sending")
	for i := 0; i<3; i++ {
		ch <- i
	}
}


func get(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

/*
	总结：
	1、数据往通道中发送时，如果接收方一直都没有接收，那么发送操作将持续阻塞
	2、接收通道中数据时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止
 */