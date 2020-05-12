/**
实现链表的逆序
方法一：就地逆序
*/
package main

import "fmt"

type Student struct {
	Name string
}

func remove(stu *Student) {
	fmt.Printf("2----%v\n", stu)
	fmt.Printf("2'----%p\n", stu)
	fmt.Printf("2''----%p\n", &stu)
	stu = nil
	fmt.Printf("3----%v\n", stu)
	fmt.Printf("3'----%p\n", stu)
	fmt.Printf("3''----%p\n", &stu)
}
func main() {
	stu := &Student{"中国"}
	fmt.Printf("1----%v\n", stu)
	fmt.Printf("1'----%p\n", stu)
	fmt.Printf("1''----%p\n", &stu)
	remove(stu)
	fmt.Printf("4----%v\n", stu)
	fmt.Printf("4'----%p\n", stu)
	fmt.Printf("4''----%p\n", &stu)
}

/**
	fmt.Printf中%p是打印指针里存储的地址，不是指印指针本身所在的地址。
 */