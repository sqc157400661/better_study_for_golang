package main

import "fmt"

func movingCount(m int, n int, k int) int {

	return 0
}

func main(){
	fmt.Println(sums(21))
}

func sums(x int) int {
	s:=0
	for x !=0 {
		s += x % 10
		x = x / 10
	}
	return s
}