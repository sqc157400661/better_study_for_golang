/**
	学习函数 - 关键词defer
 */
package main

import "fmt"

func main() {
	fmt.Printf("main\n")
	df()
}

func df() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Printf("defer\n")
	panic("123123")
}
/*
	总结：
	1、关键词defer可以在函数结束时运行某个语句或函数(可看做一种延时机制)
	2、当 go 执行到一个 defer 时，不会立即执行 defer 后的语句，而是将 defer 后的语句压入到一个栈中[为了理解，暂时称该栈为 defer 栈],然后继续执行函数下一个语句。

	经常使用defer进行程序或者函数的收尾工作如：
	1、关闭文件流 defer file.Close()
	2、解锁一个加锁的资源 defer mu.Unlock()
	3、打印最终报告 defer printFooter()
	4、关闭数据库链接 defer disconnectFromDB()
	5、打印日志和程序追踪

*/



