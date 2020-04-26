/**
	学习接口
 */
package main

import "fmt"

// 定义展示的接口
type ShowBaseI interface {
	show()
}
type ShowColor struct {

}
func (p * ShowColor) show(){
	fmt.Println("showColor func")
}
type ShowPic struct {

}
func (p * ShowPic) show(){
	fmt.Println("showPic func")
}

func main() {
	var basei ShowBaseI

	//ShowColor既实现了ShowBaseI，也实现了interface{}。
	//var showColor interface{} = new(ShowColor)
	var showColor ShowBaseI = new(ShowColor)

	// 检测是否是实现了ShowBaseI接口
	if t, ok := showColor.(ShowBaseI); ok {
		fmt.Printf("The type of showColor is: %T\n", t)
	}else{
		fmt.Printf("%T\n ", t)
	}

	// 检测是否是实现了interface{}接口
	if t, ok := showColor.(interface{}); ok {
		fmt.Printf("The type of showColor2 is: %T\n", t)
	}else{
		fmt.Printf("%v\n", t)
	}

	// 检测是否是ShowPic类型
	basei = new(ShowPic)
	if t, ok := basei.(*ShowPic); ok {
		fmt.Printf("The type of ShowPic is: %T\n", t)
	}else{
		fmt.Printf("%v\n", t)
	}

}
/*
	总结：
	对于 T 是一个接口类型的情况，当断言值 i 的动态类型存在并且此动态类型实现了接口类型 T，则此断言将成功；否则，此断言失败。
	一个失败的类型断言的估值结果为断言类型的零值。
 */



