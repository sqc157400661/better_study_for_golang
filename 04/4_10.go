/**
	学习复合数据类型-结构体
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time
		Position  string
		salary    int
		ManagerID int
		Pid *Employee  // 注意这里 Employee会报错
	}

	// dilbert 是值类型
	var dilbert Employee

	// 结构体变量的成员可以通过点操作符访问或者赋值
	dilbert.salary = 5000
	fmt.Println(dilbert.salary,dilbert)

	// 使用new关键词  new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针
	// 惯用方法是：dilbert1 := new(Employee) 等同于 var dilbert1 *Employee; dilbert1=new(Employee)
	dilbert1 := new(Employee)
	dilbert1.ID = 12
	fmt.Println(dilbert1.ID,dilbert1)

}
/*
总结：
	1、结构体变量的成员可以通过点操作符访问或者赋值
	2、无论变量是一个结构体类型还是一个结构体类型指针，都可以使用同样的 选择器符（selector-notation）【点操作符】 来引用结构体的字段
	3、如果结构体成员名字是以大写字母开头的，那么该成员就是导出的；这是Go语言导出规则决定的。一个结构体可能同时包含导出和未导出的成员
	4、一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身。（该限制同样适应于数组。）但是S类型的结构体可以包含*S指针类型的成员，这可以让我们创建递归的数据结构
	*/