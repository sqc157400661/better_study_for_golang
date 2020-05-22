/*
	如何实现栈 - 数组实现
	实现一个栈的数据结构，使其具有以下方法：压栈、弹栈、取栈顶元素、判断栈是否为空以及获取栈中元素个数。
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

// 判断是否是空栈
func (p *SliceStack) IsEmpty() bool {
	return p.StackSize == 0
}

// 获取栈的大小
func (p *SliceStack) Size() int {
	return p.StackSize
}

// 取栈顶元素
func (p *SliceStack) Top() int {
	if p.IsEmpty() {
		panic(errors.New("栈已经为空."))
	}
	fmt.Println(p.Arr)
	return p.Arr[p.StackSize-1]
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

// 压栈
func (p *SliceStack) Push(t int) {
	p.Arr = append(p.Arr, t)
	p.StackSize = p.StackSize + 1
}
func SliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Slice构建栈结构")
	sliceStack := &SliceStack{Arr: make([]int, 0)} // 注意这里make第二个参数len需要是0
	sliceStack.Push(1)
	sliceStack.Push(2)
	sliceStack.Push(3)
	fmt.Println("栈顶元素为：", sliceStack.Top())
	fmt.Println("栈大小为：", sliceStack.Size())
	sliceStack.Pop()
	fmt.Println("弹栈成功：", sliceStack.Size())
	sliceStack.Pop()
}
func main() {
	SliceMode()
}
