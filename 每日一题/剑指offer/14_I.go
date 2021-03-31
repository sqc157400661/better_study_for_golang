package main

import "fmt"

func cuttingRope(n int) int {
	if n < 4 {
		return n-1
	}
	res :=1
	for n>4 {
		res *=3
		n -= 3
	}
	tmp := int64(res * n)
	return int(tmp % 1000000007)// 乘以剩余的数
}

func main(){
	fmt.Println(cuttingRope(1))
	fmt.Println(cuttingRope(2))
	fmt.Println(cuttingRope(3))
	fmt.Println(cuttingRope(4))
	fmt.Println(cuttingRope(5))
	fmt.Println(cuttingRope(6))
}