/*
	如何实现队列 - 数组实现
	实现一个队列的数据结构，使其具有入队列、出队列、查看队列首尾元素、查看队列大小等功能
 */
package main

import (
	"errors"
	"fmt"
)

// 队列定义
type SliceQueue struct {
	Arr       []int
	front int // 队列头
	rear  int //  队列尾
}

// 判断队列是否是空
func (p *SliceQueue) IsEmpty() bool {
	return p.front == p.rear
}

// 获取队列的大小
func (p *SliceQueue) Size() int {
	return p.rear - p.front
}

// 获取队列首元素
func (p *SliceQueue) GetFront() int {
	if p.IsEmpty() {
		panic(errors.New("队列已经为空."))
	}
	return p.Arr[p.front]
}

// 获取队列尾元素
func (p *SliceQueue) GetBack() int {
	if p.IsEmpty() {
		panic(errors.New("队列已经为空."))
	}
	return p.Arr[p.rear-1]
}

// 删除队列头元素
func (p *SliceQueue) DeQueue()  {
	if p.rear > p.front {
		p.rear--
		p.Arr = p.Arr[1:]
	}else {
		panic(errors.New("队列已经为空.."))
	}
}
// 把新元素加入队列尾
func (p *SliceQueue) EnQueue(t int) {
	p.Arr = append(p.Arr, t)
	p.rear++
}


func SliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Slice构建队列结构")
	sliceQueue := &SliceQueue{Arr: make([]int, 0)} // 注意这里make第二个参数len需要是0
	sliceQueue.EnQueue(1)
	sliceQueue.EnQueue(2)
	sliceQueue.EnQueue(3)
	fmt.Println("队列头元素为：", sliceQueue.GetFront())
	fmt.Println("队列尾元素为：", sliceQueue.GetBack())
	fmt.Println("队列大小为：", sliceQueue.Size())
}
func main() {
	SliceMode()
}
