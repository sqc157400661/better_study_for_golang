/**
	学习通道- 信号量模式
 */
package main

import (
	"fmt"
)

func sumxy(x, y int, c chan int) {
	c <- x + y
}

func main() {
	ch := make(chan int)
	go sumxy(66, 88, ch)
	fmt.Println(<-ch)
}

/*
	总结：
	1、协程通过在通道 ch 中放置一个值来处理结束的信号。main 协程等待 <-ch 直到从中获取到值。
 */
