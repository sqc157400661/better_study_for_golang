package main

import "fmt"

// 贪婪算法
func maxProfit(prices []int) int {
	var tmp,res,start int
	for i:=0;i<len(prices);i++{
		tmp = prices[i] -  prices[start]
		if tmp <= 0 {
			start = i
			tmp = 0
		}
		res = max(res,tmp)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}
