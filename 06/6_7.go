/**
	学习通道- 协程超时机制
 */
package main
import (
	"fmt"
	"time"
)
func main() {
	ch1 := make(chan int)
	quit := make(chan bool)
	//开一个协程
	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):  // 时间可能不是准确的3秒
				fmt.Println("超时")
				quit <- true
			}
		}
	}()
	for i := 0; i < 5; i++ {
		ch1 <- i
		time.Sleep(time.Second)
	}
	<-quit // 利用通道阻塞控制信号量
	fmt.Println("程序结束")
}

