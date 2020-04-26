/**
	学习接口
 */
package main
import "fmt"
func main() {
	// 编译器将把123的类型推断为内置类型int。
	var x interface{} = 123
	// 情形一：断言成功
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true
	fmt.Printf("n 的类型是%T \n",n) // int
	n = x.(int) // 断言成功
	fmt.Println(n) // 123

	// 情形二：断言失败
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false
	fmt.Printf("a 的类型是%T \n",a) // float64
	// 情形三：
	//a = x.(float64) // 没有第二个值来接收结果，将产生一个错误
}
/*
	总结：
	对于 T 是一个非接口类型的情况，如果断言值 varI 的动态类型存在并且此动态类型和 T 为同一类型，则此断言将成功；
否则，此断言失败。 当此断言成功时，此类型断言表达式的估值结果为断言值 varI 的动态值的一个复制。它的类型是T
	无论 T 是什么类型，如果 x 是 nil 接口值，类型断言都会失败
	一个失败的类型断言的估值结果为断言类型的零值。
 */



