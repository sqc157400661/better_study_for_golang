/**
	学习切片-切片追加
 */
package main

import "fmt"

func main() {
	var a []int
	a = append(a, 1) // 追加1个元素
	a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
	a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
	fmt.Println(a)

	var numbers []int
	// 循环向 numbers 切片添加10个数。
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		// 打印输出切片的长度、容量和指针变化。使用 len() 函数查看切片拥有的元素个数，使用 cap() 函数查看切片的容量情况。
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

}
/*
	总结：
	1、内置的append函数用于向slice追加元素：func append(s[]T, x ...T) []T
	2、append函数将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片
	3、append函数追加时元素必须和原切片的元素同类型
	4、如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储（2 倍数扩充）
	5、append后返回的切片可能已经指向一个不同的相关数组了。append 方法总是返回成功，除非系统内存耗尽了
 */




