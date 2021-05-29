package main

import (
	"fmt"
	"runtime"
	"time"
)

// Main Goroutine
func main() {
	// 模拟单核 CPU
	runtime.GOMAXPROCS(1)
	go func() {
		for true {
			fmt.Println("11212")
		}
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("shiqingchuang")
}
