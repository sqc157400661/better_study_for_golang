/**
	学习复合数据类型-结构体
 */
package main

import "fmt"

type Animal struct {
	mouth string
	leg int
	brain string
	face string
}

type Humen struct {
	Animal
	say string
	run string
	think string
}

type Man struct {
	Humen //内嵌结构体  结构体也是一种数据类型，所以它也可以作为一个匿名字段来使用
	sex int
	game string
	string //声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员
}

func main() {
	//匿名字段和内嵌结构体
	man := new(Man)
	man.sex = 1
	man.leg = 3 //匿名成员名字在点操作符中是可选的。我们在访问子成员的时候可以忽略任何匿名成员
	//但是 结构体字面值并没有简短表示匿名成员的语法，
	liupt := Man{Humen:Humen{},sex:1,game:"lol",string:"一个男人",}
	//liupt := Man{Humen{},sex:1,game:"lol",string:"一个男人",} //报错
	fmt.Println(liupt)
}

