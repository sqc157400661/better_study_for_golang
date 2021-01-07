package main

import "fmt"

func add(a int, b int) int {
	c := 0
	// 进位为0时,加法计算完成,退出循环
	for b != 0 {
		c = a & b << 1 // 进位
		a = a ^ b      //非进位和  并赋值给a 进而进入下一次循环使用
		b = c // 进位赋值给b,进而进入下一次循环使用
	}
	return a
}

func main() {
	fmt.Println(add(5, 3))
	fmt.Println(add(5, 10))
	fmt.Println(add(5, -10))
}
