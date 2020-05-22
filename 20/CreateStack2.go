/*
	如何实现栈 - 链表实现
	实现一个栈的数据结构，使其具有以下方法：压栈、弹栈、取栈顶元素、判断栈是否为空以及获取栈中元素个数。
 */
package main

import (
	"errors"
	"fmt"
)

//链表定义
type LNode struct {
	Data interface{} // 数据域
	Next *LNode      //  对下一个节点的作用
}

// 栈定义
type LinkedStack struct {
	head *LNode
}

// 判断是否是空栈
func (p *LinkedStack) IsEmpty() bool {
	return p.head.Next == nil
}

// 获取栈的大小
func (p *LinkedStack) Size() int {
	size := 0
	node := p.head.Next
	for node != nil {
		node = node.Next
		size++
	}
	return size
}

// 取栈顶元素
func (p *LinkedStack) Top() int {
	if p.IsEmpty() {
		panic(errors.New("栈已经为空."))
	}
	return p.head.Next.Data.(int)
}

// 弹栈
func (p *LinkedStack) Pop() int {
	tmp := p.head.Next
	if tmp != nil {
		p.head.Next = tmp.Next
		return tmp.Data.(int)
	}
	panic(errors.New("栈已经为空.."))
}

// 压栈
func (p *LinkedStack) Push(t int) {
	node := &LNode{Data: t, Next: p.head.Next}
	p.head.Next = node
}
func LinkedMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("链表构建栈结构")
	linkedStack := &LinkedStack{head: &LNode{}} // 注意这里make第二个参数len需要是0
	linkedStack.Push(1)
	linkedStack.Push(2)
	linkedStack.Push(3)
	fmt.Println("栈顶元素为：", linkedStack.Top())
	fmt.Println("栈大小为：", linkedStack.Size())
	linkedStack.Pop()
	fmt.Println("弹栈成功：", linkedStack.Size())
	linkedStack.Pop()
}
func main() {
	LinkedMode()
}
