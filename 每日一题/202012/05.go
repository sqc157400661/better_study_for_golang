package main

import "fmt"

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	// stack1 不为空 stack2为空 把stack1的内容压到stack2中  用来进行负负得正
	if len(stack2) == 0 {
		for len(stack1) > 0 {
			tmp := stack1[0]
			if len(stack1) > 1{
				stack1 = stack1[1:]
			}else{
				stack1 = []int{}
			}
			stack2 = append(stack2, tmp)
		}
	}
	var re int
	if len(stack2) <= 0 {
		return -1
	} else if len(stack2) == 1 {
		//清空stack2
		re = stack2[0]
		stack2 = []int{}
	} else {
		// pop stack2
		re = stack2[0]
		stack2 = stack2[1:]

	}
	return re
}

func main(){
	arr := [][]int{{1,1},{1,2},{1,3},{2},{2},{1,4},{2},{1,5},{2},{2}}
	for _,v:= range arr{
		if v[0] ==1 {
			Push(v[1])
		}else{
			fmt.Println(Pop())
		}
	}
}
