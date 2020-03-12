/**
	学习复合数据类型-结构体
 */
package main

import "fmt"

type Cat struct {
	Color string "我是成员Color的可选标签哦" // 反射的时候可以取到这个标签
	Name  string
}
func main() {
	//结构体工厂方法
	a := NewCat("sqc","blue")
	fmt.Println(a)

}
func NewCat(color string,name string,) *Cat {
	return &Cat{
		Name: name,
		Color:color,// }换行这里要加一个,
	}
}
/*
总结：
	1、同一个包下结构体名字不能重复
	2、&结构体{成员值列表}  是一种简写，底层仍然会调用 new ()  即表达式 new(Type) 和 &Type{} 是等价的。
	3、可以用结构体值或者指针变量 使用点对成员进行赋值，如果是结构体指针 也可以通过指针的方式来设置值
*/