/**
	学习通道- 通道阻塞
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	// go getData(ch) // 1.0逻辑
	getData3(ch) // 2.0
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch) // 2.0
}

// 1.0逻辑
func getData(ch chan string) {
	var input string
	//time.Sleep(1e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

// 改造后的 2.0逻辑
func getData2(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}

// 使用 for-range 语句来读取通道是更好的办法，因为这会自动检测通道是否关闭
func getData3(ch chan string) {
	for input := range ch {
		fmt.Printf("%s ", input)
	}
}


