/**
	学习切片-切片中删除元素
 */
package main

import "fmt"

func main() {

	a := []int{1, 2, 3,4,5,6}
	fmt.Printf("a:%v,地址:%p\n", a, a)
	a = a[:len(a)-1] // 删除尾部1个元素
	a = a[:len(a)-2] // 删除尾部2个元素
	fmt.Printf("a:%v,地址:%p\n", a, a)

	// 坑点
	//for i,v := range a{
	//	if v%2 ==1 {
	//		a = append(a[:i], a[i+1:]...)
	//	}
	//}
	//fmt.Printf("a:%v,地址:%p\n", a, a)


	seq := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("seq:%v,地址:%p\n", seq, seq)
	// 指定删除位置
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Printf("seq:%v,地址:%p\n", seq, seq)




}
/*
	总结：
	1、Go语言并没有对删除切片元素提供专用的语法或者接口，
需要使用切片本身的特性来删除元素。根据要删除元素的位置有三种情况：
从开头位置删除，从中间位置删除，从尾部删除。其中删除切片尾部的元素最快：
 */




