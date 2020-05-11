/**
	学习并发
	普通函数创建 goroutine
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
	go running()

	// main程序sleep 5s后退出  main主协程结束，其他协程也会结束
	time.Sleep(5 * 1e9)
}




