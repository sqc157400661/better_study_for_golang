/**
	学习函数 - 函数玩法 - 闭包及应用
 */
package main

import "fmt"

func main() {
	/*
		Go语言中闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，
		即使已经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量，因此，简单的说：
		函数 + 引用环境 = 闭包
	 */
	// 准备一个字符串
	str := "hello world"
	// 创建一个匿名函数
	foo := func() {

		// 匿名函数中访问str
		str = "hello dude"
	}
	// 调用匿名函数
	foo()


	// 匿名函数作为返回值1
	add := add2(2) // 赋值给add这个变量的时候，函数b这个形参的值得以记忆
	fmt.Printf("the result is %v \n",add(2))
	fmt.Printf("the result is %v \n",add(4))
	// 匿名函数作为返回值2
	twoAdd := adder()
	fmt.Printf("the result of twoAdd is %v \n",twoAdd(3))
	fmt.Printf("the result of twoAdd is %v \n",twoAdd(3))
	fmt.Printf("the result of twoAdd is %v \n",twoAdd(3))
}


func add2(b int) func(a int) int {
	return func(a int) int {
		return a + b;
	}
}

/*
	细节说明
	1.adder 是一个函数，返回的数据类型是 fun (int) int
	2.返回的是一个匿名函数, 但是这个匿名函数引用到函数外的x ,因此这个匿名函数就和x 形成一 个整体，构成闭包。
	3.大家可以这样理解: 闭包是类, 函数是操作，x 是字段。函数和它使用到 x  构成闭包
	4.当我们反复的调用 twoAdd 函数时，因为 x 是初始化一次，因此每调用一次就进行累计。
	5.我们要搞清楚闭包的关键，就是要分析出返回的函数它使用(引用)到哪些变量，因为函数和它引
	用到的变量共同构成闭包。
	6.对上面代码的一个修改，加深对闭包的理解
 */
func adder() func(a int) int {
	var x int
	return func(delta  int) int {
		x += delta
		return x;
	}
}
/*
	总结：
	1、闭包其实就是一个函数和与其相关的引用环境组合的一个整体(实体)
	2、注意：匿名函数外部变量 如例子中x变量。是会被保留的
 */





