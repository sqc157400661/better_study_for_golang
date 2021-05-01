package main

import (
	"fmt"
	"runtime"
)

// Main Goroutine
func main() {
	// 模拟单核 CPU
	runtime.GOMAXPROCS(1)

	// 模拟 Goroutine 死循环
	go func() {
		for true {

		}
	}()

	fmt.Println("shiqingchuang")
}
