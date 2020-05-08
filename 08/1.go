/**
	写出下面代码输出内容。
 */
package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

/**
	打印后
	打印中
	打印前
	panic: 触发异常

	defer执行顺序 解答：
	defer 是后进先出。
	协程遇到panic时，遍历本协程的defer链表，并执行defer。
	在执行defer过程中，遇到recover则停止panic，返回recover处继续往下执行。
	如果没有遇到recover，遍历完本协程的defer链表后，向stderr抛出panic信息。
	从执行顺序上来看，实际上是按照先进后出的顺序执行defer
 */