package main

import "fmt"

// 约瑟夫环问题
func lastRemaining(n int, m int) int {
	pos :=0 // 最终活下来那个人的初始位置
	for i:=2;i<=n;i++{
		pos = (pos + m)%i // 每次循环右移
	}
	return pos
}

func main(){
	fmt.Println(lastRemaining(5,3))
}