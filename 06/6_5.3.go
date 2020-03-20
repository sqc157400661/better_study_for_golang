/**
	学习通道- 信号量模式
 */
package main

import (
	"fmt"
	"time"
)

type Empty interface {}
type semaphore chan Empty

// acquire n resources
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resouces
func (s semaphore) V(n int) {
	for i:= 0; i < n; i++{
		<- s
	}
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}


func main() {
	var handle semaphore  = make(semaphore,1)

	go handle.Wait(3)
	for i:= 0;i<10;i++{
		go func() {
			handle.Signal() // 这里只会通过3次 其他协程都在等待 直到main协程结束 从而达到控制协程并发执行的数量
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 6)
}



/*
	总结：
	1、信号量是实现互斥锁（排外锁）常见的同步机制，限制对资源的访问，解决读写问题
	2、带缓冲的通道可以实现：（1）带缓冲通道的容量和要同步的资源容量相同 （2）通道的长度（当前存放的元素个数）与当前资源被使用的数量相同
（3）容量减去通道的长度就是未处理的资源个数（标准信号量的整数值）
 */
