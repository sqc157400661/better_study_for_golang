/**
	学习通道- 多路复用
 */
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		default:
			fmt.Println(i)
		}
	}
}

/*
	解释：

 */

