/**
	学习通道- 协程定时器
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个打点器, 每秒触发一次
	tick := time.Tick(1 * time.Second)
	// 创建一个计时器, 5秒后触发
	boom := time.After(5 * time.Second)
	// 不断地检查通道情况
	for {
		select {
		case <-tick:
			// 打点器触发了
			fmt.Println("tick.")
		case <-boom:
			// 计时器到时了
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}

/*
	解释：
	1、time.Tick(d) 函数声明为 Tick(d Duration) <-chan Time，它会返回一个时间的通道：它以 d 为周期给返回的通道发送时间，d 是纳秒数
	2、time.After(d) 函数声明为 func After(d Duration) <-chan Time 类似 Tick()，但是它只发送一次时间，d 是纳秒数
 */

