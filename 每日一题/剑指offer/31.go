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

func(s *stack) peek() int {
	return s.arr[len(s.arr)-1]
}

func(s *stack) IsEmpty() bool{
	return len(s.arr) ==0
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := &stack{arr:[]int{}}
	i := 0
	for _,v := range pushed {
		stack.Push(v)
		for !stack.IsEmpty() && stack.peek() == popped[i]{
			stack.Pop()
			i++
		}
	}
	return stack.IsEmpty()
}