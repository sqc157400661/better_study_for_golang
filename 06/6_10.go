/**
	学习通道- 多核并行化
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

// 分配给每个CPU的计算任务
func DoSome(c chan int) {
	a :=1
	for i:=0; i < 1000000000; i++ {
		a += i
	}
	c <- 1 // 发信号告诉任务管理者我已经计算完成了
}

func DoAll(Cpu int) {
	c := make(chan int, Cpu) // 用于接收每个CPU的任务完成信号

	for i := 0; i < Cpu; i++ {
		go DoSome(c)
	}
	// 等待所有CPU的任务完成
	for i := 0; i < Cpu; i++ {
		<-c // 获取到一个数据，表示一个CPU计算完成了
	}
	// 到这里表示所有计算已经结束
}
func main() {
	start := time.Now()
	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	runtime.GOMAXPROCS(10) //设置需要用到的cpu数量
	fmt.Println("cpu核心数:", cpuNum)
	DoAll(cpuNum)
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)
	//runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量
}

/*
	总结：
可以先通过设置环境变量 GOMAXPROCS 的值来控制使用多少个 CPU 核心，或者使用语句runtime.GOMAXPROCS(cpuNum)
 */

