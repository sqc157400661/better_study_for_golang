package main

import "fmt"

/*
我们只要用一个变量记录一个历史最低价格 minprice，那么我们在第 i 天卖出股票能得到的利润就是 prices[i] - minprice。

因此，我们只需要遍历价格数组一遍，记录历史最低点，
然后在每一天考虑这么一个问题：如果我是在历史最低点买进的，那么我今天卖出能赚多少钱？当考虑完所有天数之时，我们就得到了最好的答案。

 */


/**
 *
 * @param prices int整型一维数组
 * @return int整型
 */
func maxProfit( prices []int ) int {
	minPrice := prices[0] // 记录最低价格
	maxProfitV := 0 // 记录最大收益
	for _,v := range prices{
		if v<minPrice{
			minPrice = v
		}else{
			tmp := v - minPrice
			if tmp > maxProfitV {
				maxProfitV = tmp
			}
		}
	}
	return maxProfitV
}

func main(){
	arr := []int{1,2,3,4,5}
	fmt.Println(maxProfit(arr))
}