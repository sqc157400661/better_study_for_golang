package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <=0 {
		return 0
	}
	var profit,cost int
	cost = prices[0]
	for i:=0;i<len(prices);i++{
		if prices[i] < cost{
			cost = prices[i]
		}
		profit = max(profit,prices[i] - cost)
	}
	return profit
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

