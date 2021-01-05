package main

type stack struct{
	arr []int
}
func(s *stack) Push(node int){
	s.arr = append(s.arr,node)
}

func(s *stack) Pop() int{
	if len(s.arr) ==0 {
		return 0
	}
	re := s.arr[len(s.arr)-1]
	if len(s.arr)>1 {
		s.arr =s.arr[0:len(s.arr)-1]
	}else{
		s.arr = nil
	}
	return re
}

func(s *stack) IsEmpty() bool{
	return len(s.arr) ==0
}


type CQueue struct {
	A *stack
	B *stack
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
func Constructor() CQueue {
	return CQueue{A:&stack{arr:[]int{}},B:&stack{arr:[]int{}}}
}


func (this *CQueue) AppendTail(value int)  {
	this.A.Push(value)
}


func (this *CQueue) DeleteHead() int {
	// 当栈 B 不为空
	if !this.B.IsEmpty() {
		return this.B.Pop()
	}
	if this.A.IsEmpty() {
		return -1
	}
	for !this.A.IsEmpty() {
		this.B.Push(this.A.Pop())
	}
	return this.B.Pop()
}