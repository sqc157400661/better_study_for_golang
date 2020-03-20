/**
	学习通道- 单向通道
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	sendChan := make(chan int)
	reciveChan := make(chan int)
	go processChannel(sendChan, reciveChan)
	go generate(sendChan)
	go printChannel(reciveChan)
	time.Sleep(1e9)
}
func generate(sendChan chan int){
	for i := 0; i<3; i++ {
		sendChan <- i // Send 'i' to channel 'ch'.
	}
}

func processChannel(in <-chan int, out chan<- int) {
	for inValue := range in {
		result := inValue*2 /// processing inValue
		out <- result
	}
}
func printChannel(out chan int) {
	for value := range out {
		fmt.Println(value)
	}
}

