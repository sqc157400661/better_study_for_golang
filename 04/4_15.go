/**
	学习方法
 */
package main
import "fmt"

// 车
type Car struct {
	brand string // 品牌
	speed int // 油门
	tap int // 档位
}
func main() {
	c := Car{"大众",2,3}
	fmt.Printf("%+v\n", c)

	speed := c.moveSpeed()
	fmt.Printf("%+v\n", speed)

	c.stop()
	fmt.Printf("%+v\n", c)

	//在来个非结构体的例子：
	fmt.Println(IntVector{1, 2, 3}.Sum()) // 输出是6
}

func (car * Car) moveSpeed() int{
	return car.speed * car.tap
}

//在接收者是指针时，方法可以改变接收者的值（或状态）
func (car * Car) stop() *Car{
	car.speed  = 0
	car.tap = 0
	return car
}

//在来个非结构体的例子：

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

/*
	函数和方法的区别
	1、调用方式，函数Function1(参数列表)；方法调用：recv.Method1(参数列表)
	2、方法有接收者，接收者必须有一个显式的名字，这个名字必须在方法中被使用
	3、receiver_type 叫做 （接收者）基本类型，这个类型必须在和方法同样的包中被声明。
 */


