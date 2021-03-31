package main

import "fmt"

//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
func cuttingRope(n int) int {
	if n < 4 {
		return n-1
	}
	var res,p int64
	p = 1000000007
	res =1
	for n>4 {
		res =(3 * res) % p
		n -= 3
	}

	tmp := res * int64(n) % p
	return int(tmp)// 乘以剩余的数
}



func main(){
	fmt.Println(cuttingRope(1))
	fmt.Println(cuttingRope(2))
	fmt.Println(cuttingRope(3))
	fmt.Println(cuttingRope(4))
	fmt.Println(cuttingRope(5))
	fmt.Println(cuttingRope(6))
}