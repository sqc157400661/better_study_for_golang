package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	//设置计数器，数值即为goroutine的个数
	wg.Add(2)
	go func() {
		//Do some work
		time.Sleep(1*time.Second)

		fmt.Println("Goroutine 1 finished!")
		wg.Done()
	}()

	go func() {
		//Do some work
		time.Sleep(2*time.Second)

		fmt.Println("Goroutine 2 finished!")
		wg.Done() //goroutine执行结束后将计数器减1
	}()

	wg.Wait()
	fmt.Printf("All Goroutine finished!")
}
