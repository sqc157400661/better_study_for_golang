/*
学习常量定义和赋值
常量的定义格式：const identifier [type] = value
可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型
 */

package main

// 引入第三方包
import "fmt"

//全局常量
const Pi = 3.14159

func main() {

	//显式类型定义：
	const a string = "abc"
	// 隐式类型定义：
	const b = "acd"

	//常量也允许使用并行赋值的形式
	const beef, two, c = "eat", 2, "veg"
	const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
	//const (
	//Monday, Tuesday, Wednesday = 1, 2, 3
	//Thursday, Friday, Saturday = 4, 5, 6
	//)
	fmt.Println(Pi,a,b,c,beef,two)
	fmt.Println(Monday,Tuesday,Wednesday,Thursday,Friday,Saturday)
	getE()
	//常量的值必须是能够在编译时就能够确定的；你可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期间就能获得。
	//const c1 = getE() //const initializer getE() is not a constant
	test()
}

func getE() int{
	fmt.Println(Pi)
	return 2
}

func test(){
	//const (
	//	RED  = 3 // 0
	//	ORANGE // 1
	//	YELLOW // 2
	//	GREEN // ..
	//	BLUE
	//	INDIGO
	//	VIOLET // 6
	//)

	//iota 可以被用作枚举值
	const (
		RED  = iota // 0
		ORANGE // 1
		YELLOW // 2
		GREEN // ..
		BLUE
		INDIGO
		VIOLET // 6
	)

	fmt.Println(RED,ORANGE,YELLOW,GREEN,BLUE,INDIGO,VIOLET)
}
