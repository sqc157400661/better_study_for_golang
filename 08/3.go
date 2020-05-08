/**
	下面的代码会输出什么，并说明原因
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/**
	B:  9
	A:  10
	A:  10
	A:  10
	A:  10
	A:  10
	A:  10
	A:  10
	A:  10
	A:  10
	A:  10
	B:  0
	B:  1
	B:  2
	B:  3
	B:  4
	B:  5
	B:  6
	B:  7
	B:  8

	解答：
	考点：go执行的随机性和闭包
	解答： 谁也不知道执行后打印的顺序是什么样的，所以只能说是随机数字。
	其中A:输出完全随机，取决于goroutine执行时i的值是多少；

	而B:一定输出为0~9，但顺序不定。

	第一个go func中i是外部for的一个变量，地址不变化，但是值都在改变。

	第二个go func中i是函数参数，与外部for中的i完全是两个变量。
	尾部(i)将发生值拷贝，go func内部指向值拷贝地址。

	所以在使用goroutine在处理闭包的时候，避免发生类似第一个go func中的问题。

 */