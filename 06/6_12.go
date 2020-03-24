/**
	学习基于共享变量的并发- 竞争条件
 */
package main

import (
	"runtime"
	"time"
)

func main() {
	var x []int
	runtime.GOMAXPROCS(2)
	go func() {
		time.Sleep(5*time.Millisecond)
		x = make([]int, 10)
	}()
	go func() {
		time.Sleep(2*time.Millisecond)
		x = make([]int, 1000000)
	}()
	time.Sleep(10*time.Millisecond)
	x[999999] = 1 // NOTE: undefined behavior; memory corruption possible!
}

