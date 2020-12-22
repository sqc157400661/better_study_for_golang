package main

import "fmt"

type stack1 struct{
	arr []string
}
func(s *stack1) Push(str string){
	s.arr = append(s.arr,str)
}

func(s *stack1) Pop() string{
	if len(s.arr) ==0 {
		return ""
	}
	re := s.arr[len(s.arr)-1]
	if len(s.arr)>1 {
		s.arr =s.arr[0:len(s.arr)-1]
	}else{
		s.arr = nil
	}
	return re
}

func(s *stack1) Size() int{
	return len(s.arr)
}
/**
 *
 * @param s string字符串
 * @return bool布尔型
 */
func isValid( s string ) bool {
	validMap := map[string]string{"{":"}","(":")","[":"]"}
	stack := &stack1{arr:[]string{}}
	for i:=0;i<len(s);i++{
		key :=s[i:i+1]
		if _,ok:=validMap[key];ok{
			stack.Push(key)
		}else {
			tmp :=stack.Pop()
			if(validMap[tmp] !=key) {
				return false
			}
		}
	}
	return stack.Size() ==0
}

func main(){
	str :="[]{[]}"
	fmt.Println(isValid(str))
}