/**
	学习复合数据类型-结构体
 */
package main

import "fmt"

func main() {
	type StructDemo struct {
		DemoMember1 int
		DemoMember2 string
		DemoMember3 float64
		DemoMember4 bool
		DemoMember5 interface{}
	}
	// demo1 := StructDemo{1,"12"} 会报错 这个写法必须按照字段顺序来初始化 而且要初始化所有字段值
	// 添加字段值的写法 不需要遵循顺序，而且可以值赋值一个字段 如下
	demo1 := StructDemo{DemoMember1:1,DemoMember2:"12"}
	fmt.Println(demo1)

	demo2 := &StructDemo{DemoMember1:1,DemoMember2:"12"}
	(*demo2).DemoMember2 = "666"
	demo2.DemoMember4 = false
	fmt.Println(demo2)

	//结构体也是可以比较的，两个结构体将可以使用==或!=运算符进行比较
	demo3 := StructDemo{DemoMember1:1,DemoMember2:"12"}
	fmt.Println(demo1 == demo3)
}
/*
总结：
	1、同一个包下结构体名字不能重复
	2、&结构体{成员值列表}  是一种简写，底层仍然会调用 new ()  即表达式 new(Type) 和 &Type{} 是等价的。
	3、可以用结构体值或者指针变量 使用点对成员进行赋值，如果是结构体指针 也可以通过指针的方式来设置值
	4、支持 == 与 !=比较运算符，但不支持 > 或 <
*/