/**
	学习通道- select 切换协程
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}


/*
	总结：
	1、信号量是实现互斥锁（排外锁）常见的同步机制，限制对资源的访问，解决读写问题
	2、带缓冲的通道可以实现：（1）带缓冲通道的容量和要同步的资源容量相同 （2）通道的长度（当前存放的元素个数）与当前资源被使用的数量相同
（3）容量减去通道的长度就是未处理的资源个数（标准信号量的整数值）
 */
