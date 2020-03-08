/**
	学习函数 - 函数玩法，将函数作为参数
 */
package main

import "fmt"

func main() {
	// 将函数作为参数
	callback(3,Add)
	// 匿名函数
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}




