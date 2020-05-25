/*
	如何翻转栈的所有元素 (有点难度)
	翻转（也称颠倒）栈的所有元素，例如输入栈{1， 2， 3， 4， 5}，其中，1处在栈顶，翻转之后的栈为{5， 4， 3， 2， 1}，其中，5处在栈顶。
 */
package main

import (
	"errors"
	"fmt"
)

// 栈定义
type SliceStack struct {
	Arr       []int
	StackSize int
}

// 压栈
func (p *SliceStack) Push(t int) {
	p.Arr = append(p.Arr, t)
	p.StackSize = p.StackSize + 1
}

// 判断是否是空栈
func (p *SliceStack) IsEmpty() bool {
	return p.StackSize == 0
}

// 弹栈
func (p *SliceStack) Pop() int {
	if p.StackSize > 0 {
		p.StackSize--
		ret := p.Arr[p.StackSize]
		p.Arr = p.Arr[:p.StackSize]
		return ret
	}
	panic(errors.New("栈已经为空.."))
}

/*
	创建一个测试的栈
 */
func CreateStack(list []int) *SliceStack{
	sliceStack := &SliceStack{Arr: make([]int, 0)} // 注意这里make第二个参数len需要是0
	for _,v := range list{
		sliceStack.Push(v)
	}
	return sliceStack
}

/*
	翻转栈顺序
 */
func ReverseStack(s *SliceStack){
	if s.IsEmpty() {
		return
	}
	// 把栈低元素移动到栈顶
	moveBottomToTop(s)
	top  := s.Pop()
	// 递归处理子栈
	ReverseStack(s)
	s.Push(top)

}

/*
	把底部元素迁移到顶部
 */
func moveBottomToTop(s *SliceStack){
	if s.IsEmpty() {
		return
	}
	var top1,top2 int
	top1 = s.Pop() // 弹出栈顶元素
	if !s.IsEmpty() {
		// 递归处理不包含栈顶元素的子栈
		moveBottomToTop(s)
		top2 = s.Pop()
		s.Push(top1)
		s.Push(top2)
	}else{
		s.Push(top1)
	}
}

func main() {
	stack := CreateStack([]int{5,4,3,2,1})
	fmt.Println(stack)
	ReverseStack(stack)
	fmt.Println(stack)
}
