/**
	学习控制结构-for 结构
 */
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
	//[Go 语言没有while 和 do...while语法 可以用条件判断迭代来实现]
	//基本形式：for 条件语句 {} 可以当成初始化语句; 修饰语句省略了
	var count int  = 1
	for count < 20 {
		fmt.Printf("hello world %v\n", count)
		count++
	}

	// 循环体内使用使用 break 或 return 来退出循环,关键字 continue 忽略剩余的循环体而直接进入下一次循环
	var count1 int
	for {
		fmt.Println("Hello world %v\n",count1)
		count1++
		if count1>10 {
			break
		}
	}
}
/*
	总结：
	1、基本形式： for 初始化语句; 条件语句; 修饰语句 {}
	2、循环开始前会执行且仅会执行一次初始化语句,然后判断条件【false的话直接退出】，最后执行修饰语句
	3、左花括号 { 必须和 for 语句在同一行，计数器的生命周期在遇到右花括号 } 时便终止

*/



