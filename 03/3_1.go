/**
	学习控制结构-if else
 */
package main

import "fmt"

func main() {
	var age int = 40
	//分支控制
	if age > 20 {
		fmt.Println("age>=20")
	}

	age3 :=10
	//golang在if判断中直接定义变量
	if age3 := 40; age3 > 20 {
		fmt.Println("age3>=20")
	}
	// if语句里age3会把外部的隐藏
	fmt.Println(age3)


	//双分支 if else
	// {}不能省略  else 不能换行
	// if 后面的括号不需要

	if age > 10 {
		fmt.Println("age>=10")
	} else if age > 20 {
		fmt.Println("age>=20")
	} else {
		fmt.Println("age=", age)
	}

}
/*
	总结：
	1、关键字 if 和 else 之后的左大括号 { 必须和关键字在同一行，如果你使用了 else-if 结构，则前段代码块的右大括号 } 必须
和 else-if 关键字在同一行。这两条规则都是被编译器强制规定的。
    2、条件语句两侧的括号()是可以被省略的；当条件比较复杂时，则可以使用括号让代码更易读,你可以使用括号来提升某个表达式的运算优先级，并提高代码的可读性
	3、在if中 用简短方式 := 声明的变量的作用域只存在于 if 结构中，会隐藏同名的外部变量
 */



