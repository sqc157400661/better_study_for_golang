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

func (p *ShowColor) show() {
	fmt.Println("showColor func")
}

type ShowPic struct {
}

func (p *ShowPic) show() {
	fmt.Println("showPic func")
}

func main() {
	var basei ShowBaseI = new(ShowPic)
	switch t := basei.(type) {
	case *ShowPic:
		fmt.Printf("Type ShowPic %T with value %v\n", t, t)
	case *ShowColor:
		fmt.Printf("Type ShowColor %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
	/*
		说明：
		1、type-switch 不允许有 fallthrough
		2、例子中 变量 t 得到了 areaIntf 的值和类型 ,不用它的值，那么就可以不需要赋值语句 switch  basei.(type) { ....}
	 */
}

/*
	总结：
	接口变量的类型也可以使用一种特殊形式的 switch 来检测：type-switch
 */



