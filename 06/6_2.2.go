/**
	学习并发
	匿名函数创建 goroutine
 */

package main
import (
	"fmt"
	"time"
)
func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}
func main() {
	// 并发执行程序
	go func(){
		var times int
		// 构建一个无限循环
		for {
			times++
			fmt.Println("tick", times)
			// 延时1秒
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(5 * 1e9)
}




