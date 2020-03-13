/**
学习复合数据类型-结构体
*/
package main

import "fmt"

// 车轮
type Wheel struct {
	Size int
}

type Car struct {
	Wheel
	// 引擎
	// 将结构体直接定义在嵌入的结构体中,结构体的定义不会被外部引用到
	Engine struct {
		Power int    // 功率
		Type  string // 类型
	}
}


func main() {
	//匿名字段和内嵌结构体
	car := new(Car)
	car.Size = 1
	//car.Power = 1
	car.Engine.Power = 1
	car.Engine.Type = "德国制造"
	fmt.Println(car)

	// 初始化
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},
		// 初始化引擎
		// 内嵌定义的结构体 ,就需要再次声明结构才能赋予数据
		Engine: struct {   // 这里Engine: 可以省略 但是Wheel: 也需要同时省略
			Power int
			Type  string
		}{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)
}

