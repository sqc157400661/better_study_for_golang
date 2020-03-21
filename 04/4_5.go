/**
	学习切片-切片追加2
 */
package main

import "fmt"

func main() {
	// 在切片的开头添加元素
	var a = []int{1,2,3}
	fmt.Printf("a:%v,地址:%p\n", a, a)
	a = append([]int{0}, a...) // 在开头添加1个元素
	a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片
	fmt.Printf("a:%v,地址:%p\n", a, a)

	var b = []int{5,6,8,9,66}
	b = append(b[:2], append([]int{111}, b[2:]...)...) // 在第2个位置插入x
	fmt.Printf("b:%v,地址:%p\n", b, b)
	//每个添加操作中的第二个 append 调用都会创建一个临时切片
	b = append(b[:3], append([]int{1,2,3}, b[3:]...)...) // 在第3个位置插入切片
	fmt.Printf("b:%v,地址:%p\n", b, b)


}
/*
	总结：
	1、可以在切片的开头添加元素，但是导致内存的重新分，性能会降低
	2、 append 函数返回新的切片，也就是它支持链式操作，可以实现在切片中间插入元素
 */




