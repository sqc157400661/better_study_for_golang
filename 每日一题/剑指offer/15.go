package main

import "fmt"

//  方法一：逐位判断
func hammingWeightV1(num uint32) int {
	var res uint32
	for num >0{
		res += num&1
		num >>=1
	}
	return int(res)
}

//  方法二：巧用 n&(n−1)
func hammingWeightV2(num uint32) int {
	var res uint32
	for num >0{
		res += 1
		num &= num -1
	}
	return int(res)
}

func main(){
	fmt.Println(hammingWeightV1(00000000000000000000000000001011))
	fmt.Println(hammingWeightV2(00000000000000000000000000001011))
}