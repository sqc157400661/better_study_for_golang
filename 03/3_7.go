/**
	学习函数 - 函数的定义
 */
package main

import "fmt"

// 函数可以成为一个类型，叫函数类型 注意：不需要函数体 {} 可以赋值给变量，如 add := binOp 。
type binOp func(int, int) int

func main() {
	var name = "ziroom"
	//greeting(name)
	Greeting1(&name)
	fmt.Println(name)
	// 第二个返回值赋给了空白符 _，然后自动丢弃掉
	var a,_ = addAndMultiply(1,2)
	fmt.Println(a)
}
/*
按值传递
就是参数是单纯的变量的值，函数对其修改不会影响函数外变量的值
 */
func greeting(name string) {
	fmt.Println(name+" Hi!!!!!")
	name = "sqc"
}
/*
引用传递
函数的参数是指针（&变量）或者引用类型的变量（切片（slice）、字典（map）、接口（interface）、通道（channel）），
函数内修改，其实就是修改的变量本身
*/
func Greeting1(name *string) {
	fmt.Println(*name+" Hi!!!!!")
	*name = "sqc"
}

func addAndMultiply(x int,y int)(addRe int,mpRe int){
	addRe = x + y
	mpRe = x * y
	return
	//return 1,2
}

/*
	总结：
	1、当命名返回值时即使只有一个命名返回值，也需要使用 () 括起来
	2、尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。
	3、即使函数使用了命名返回值，你依旧可以无视它而返回明确的值(以return后面参数的值为优先)。

*/



