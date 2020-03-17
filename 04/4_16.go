/**
	学习方法
	指针或值作为接收者
*/
package main

import "fmt"

// 车
type Test struct {
	num int
}
func main() {
	b := Test{1}  //b := &Test{1}  //指针方法和值方法都可以在指针或非指针上被调用
	// 传递Test类型的指针  会修改接收者的值和状态
	// b会被自动转换为 &b
	b.addOne()
	b.show()
	b.addOne()
	b.show()

	// 传递值  不会修改接收者的值和状态  进行值拷贝
	b.add()
	b.showR()
	b.add()
	b.showR()
}

func (b *Test) addOne(){
	b.num +=1
}

func (b *Test) show(){
	fmt.Println(b)
}

func (b Test) add(){
	b.num +=1
}
func (b Test) showR(){
	fmt.Println(b)
}

/*

 */


