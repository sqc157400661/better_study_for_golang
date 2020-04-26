/**
	学习接口
 */
package main

import "fmt"

type Shaper interface {
	// 要继承该接口，1、需要实现Area方法，2、有一个参数，类型是float32 3、有一个float32的返回值
	// 如果有多个方法，则都需要实现
	Area(ext float32) float32
}

type Square struct {
	side float32
}

func (sq *Square) Area(ext float32) float32 {
	return (sq.side * sq.side + ext)
}

func main() {
	// 实例化
	sq1 := new(Square)
	sq1.side = 5

	// 定义一个接口类型的变量  然后将实例化的子类赋值给变量 这是指针指向了实例化的子类
	//var areaIntf Shaper
	//areaIntf = sq1

	// 简短方式，无需声明接口类型的变量 接口是个动态的类型
	areaIntf := Shaper(sq1)

	// 直接使用子类调用
	//areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area(3))
}
/*
	总结：
	接口被实现的条件一：接口的方法与实现接口的类型方法格式一致【方法名称、参数列表、返回参数列表需要格式一致】
	接口被实现的条件二：接口中所有方法均被实现
	类型不需要显式声明它实现了某个接口：接口被隐式地实现。
	只要类 (结构体) 实现对应的接口，那么根据该类创建的对象，可以赋值给对应的接口类型。
 */



