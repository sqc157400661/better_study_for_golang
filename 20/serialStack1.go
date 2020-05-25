/*
	如何根据入栈序列判断可能的出栈序列
	输入两个整数序列，其中一个序列表示栈的push（入）顺序，判断另一个序列有没有可能是对应的pop（出）顺序。
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

// 取栈顶元素
func (p *SliceStack) Top() int {
	if p.IsEmpty() {
		panic(errors.New("栈已经为空."))
	}
	return p.Arr[p.StackSize-1]
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

// 判断push序列和pop序列是否能匹配
func IsPopSerial(push []int, pop []int) bool {
	var pushLen, popLen int
	pushLen = len(push)
	popLen = len(pop)
	// 做基础的判断
	if pushLen == 0 || popLen == 0 || pushLen != popLen {
		return false
	}
	pushIndex := 0
	popIndex := 0
	stack := &SliceStack{Arr: make([]int, 0)}
	for pushIndex < pushLen {
		// 把push序列依次入栈，直到栈顶元素等于pop序列的第一个元素
		stack.Push(push[pushIndex])
		pushIndex++
		//栈顶元素出栈，pop序列移动到下一个元素
		for !stack.IsEmpty() && stack.Top() == pop[popIndex] {
			stack.Pop()
			popIndex++
		}
	}
	if stack.IsEmpty() && popIndex == popLen {
		return true
	}
	return false
}

func main() {
	push := []int{1, 2, 3, 4, 5}; // push 顺序
	pop := []int{3, 2, 5, 4, 1};  // pop 顺序
	if IsPopSerial(push, pop) {
		fmt.Println(pop, "是", push, "的一个pop序列")
	} else {
		fmt.Println(pop, "不是", push, "的一个pop序列")
	}
}
