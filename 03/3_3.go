/**
	学习控制结构-switch 结构
 */
package main

import "fmt"

func main() {
	//Type Switch:switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际指向的 变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是 %T", i) // 标识没有赋值情况下的interface{}类型
		//fallthrough
	case int:
		fmt.Printf("x的类型是int")
	case float64:
		fmt.Printf("x的类型是float64")
	case func(int) float64:
		fmt.Printf("x的类型是func(int)")
	case bool, string:
		fmt.Printf("x的类型是bool或者string")
	default:
		fmt.Printf("未知类型")
		//switch 和 if 的比较
		//如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几种类型。建议使用 swtich 语句，简洁高效。
		//其他情况:对区间判断和结果为 bool 类型的判断，使用 if，if 的使用范围更广。
	}
}
/*
	总结：
	1、在 Type Switch 语句的 case 子句中不能使用fallthrough
	2、Type Switch 是 Go 语言中一种特殊的 switch 语句，它比较的是类型而不是具体的值。它判断某个接口变量的类型【type-switch 只用于判断interface】

*/



