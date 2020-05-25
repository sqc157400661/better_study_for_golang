/*
	如何实现队列 - 链表实现
	实现一个队列的数据结构，使其具有入队列、出队列、查看队列首尾元素、查看队列大小等功能
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

// 队列定义
type LinkedQueue struct {
	head *LNode  // 队列头
	end *LNode // 队列尾
}

// 判断队列是否为空
func (p *LinkedQueue) IsEmpty() bool {
	return p.head == nil
}

// 获取队列的大小
func (p *LinkedQueue) Size() int {
	size :=0
	node := p.head
	for node != nil{
		node = node.Next
		size++
	}
	return size
}

// 入队列 把新元素加到队尾
func (p *LinkedQueue) EnQueue(t int)  {
	node := &LNode{Data:t}
	if p.head == nil {
		p.head = node
		p.end = node
	}else{
		p.end.Next = node
		p.end = node
	}
}

// 出队列，删除队列的首元素
func (p *LinkedQueue) DeQueue() {
	if p.head ==nil {
		panic(errors.New("队列已经为空"))
	}
	p.head = p.head.Next
	if p.head == nil {
		p.end = nil
	}
}

// 获取队列首元素
func (p *LinkedQueue) GetFront() int {
	if p.IsEmpty() {
		panic(errors.New("队列已经为空."))
	}
	return p.head.Data.(int)
}

// 获取队列尾元素
func (p *LinkedQueue) GetBack() int {
	if p.IsEmpty() {
		panic(errors.New("队列已经为空."))
	}
	return p.end.Data.(int)
}



func LinkedMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("链表构建队列结构")
	linkedQueue := &LinkedQueue{head: nil,end:nil} // 注意这里make第二个参数len需要是0
	linkedQueue.EnQueue(1)
	linkedQueue.EnQueue(2)
	linkedQueue.EnQueue(3)
	fmt.Println("队列头元素为：", linkedQueue.GetFront())
	fmt.Println("队列尾元素为：", linkedQueue.GetBack())
	fmt.Println("队列大小为：", linkedQueue.Size())
}
func main() {
	LinkedMode()
}
