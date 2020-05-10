/**
	学习接口
 */
package main

import (
	"fmt"
)

type Human struct {

}

func (s *Human) Eat(food string) error{
	fmt.Printf("The human  eat  %v \n",food)
	return nil
}
func (s *Human) Walk(speed int) (n int, err error) {
	fmt.Printf("The human  walk in speed %v \n",speed)
	return speed, nil
}
// 除了继承接口的方法也可以有自己的方法
func (s *Human) See(something string) string {
	fmt.Printf("The human see %v \n",something)
	return "The human see"+something
}

// 腿 接口
type ILeg interface {
	Walk(speed int) (n int, err error)
}
// 手接口
type IHand interface {
	Eat(food string) error
}

func Walk(obj ILeg,speed int) {
	obj.Walk(speed)
}

func Eat(obj IHand,food string) {
	obj.Eat(food)
}
func main() {
	// 实例化
	s := new(Human)
	Walk(s,10)
	Eat(s,"黄瓜")
}

/*
	总结：
	一个类型可以实现多个接口
 */



