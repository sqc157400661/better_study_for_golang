/**
	学习方法
	内嵌类型的方法和继承
 */
package main
import "fmt"
// 可飞行的
type Flying struct{}
func (f *Flying) Fly() {
	fmt.Println("can fly")
}
// 可行走的
type Walkable struct{}
func (f *Walkable) Walk() {
	fmt.Println("can walk")
}
// 人类
type Human struct {
	Walkable // 人类能行走
}
func (f *Human) Walk() {
	fmt.Println("Human  walk")
}
// 鸟类 多重继承
type Bird struct {
	Walkable // 鸟类能行走
	Flying   // 鸟类能飞行
}
func main() {
	// 实例化鸟类
	b := new(Bird)
	fmt.Println("Bird: ")
	b.Fly()
	b.Walk()
	// 实例化人类
	h := new(Human)
	fmt.Println("Human: ")
	h.Walk()
}
/*
	可以使用内嵌结构体来仿照其他语言的继承
	可以覆写方法，具有同样名字的外层类型的方法会覆写内嵌类型对应的方法
	可以多重继承，一个结构体嵌入多个匿名类型来实现
 */


