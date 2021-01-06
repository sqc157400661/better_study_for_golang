/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
package main

type stack struct {
	arr []int
}

func (s *stack) Push(node int) {
	s.arr = append(s.arr, node)
}

func (s *stack) Pop() int {
	if len(s.arr) == 0 {
		return 0
	}
	re := s.arr[len(s.arr)-1]
	if len(s.arr) > 1 {
		s.arr = s.arr[0 : len(s.arr)-1]
	} else {
		s.arr = nil
	}
	return re
}

func (s *stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *stack) Peek() int {
	if len(s.arr) == 0 {
		return 0
	}
	return s.arr[len(s.arr)-1]
}

type MinStack struct {
	A *stack
	B *stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{A: &stack{arr: []int{}}, B: &stack{arr: []int{}}}
}

func (this *MinStack) Push(x int) {
	this.A.Push(x)
	if this.B.IsEmpty() || this.B.Peek() >= x{
		this.B.Push(x)
	}
}

func (this *MinStack) Pop() {
	re := this.A.Pop()
	if re == this.B.Peek() {
		this.B.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.A.Peek()
}

func (this *MinStack) Min() int {
	return this.B.Peek()
}
