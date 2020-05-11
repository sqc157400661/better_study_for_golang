/**
Go性能调优 - Pprof
runtime/pprof：采集工具型应用运行数据进行分析
net/http/pprof：采集服务型应用运行时数据进行分析
*/
package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// 一段有问题的代码
func do() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("我是有问题的那一行，通道无法读取到值：%v", v)
		default:

		}
	}
}

func main() {
	// 创建分析文件
	file, err := os.Create("./cpu.prof")
	if err != nil {
		fmt.Printf("创建采集文件失败, err:%v\n", err)
		return
	}
	// 进行cpu数据的获取
	if err = pprof.StartCPUProfile(file); err == nil {
		defer pprof.StopCPUProfile()
	}
	// 执行一段有问题的代码
	for i := 0; i < 4; i++ {
		go do()
	}
	time.Sleep(10 * time.Second)
}
