package main

import (
	"fmt"
	"math"
)

// 一个数字出现一次，其他都出现了三次
func singleNumbers(nums []int) int {
	res := 0
	for i:=0; i<32;i++  {
		cnt :=0 // 记录当前 bit 有多少个1
		bit := 1<<i //  记录当前要操作的 bit
		for i := 0; i < len(nums); i++ {
			if nums[i] & bit !=0 {
				cnt++
			}
		}
		// 不等于0说明唯一出现的数字在这个 bit 上是1
		if cnt % 3 != 0{
			res |= bit
		}
	}
	// 负数处理
	if res > int(math.Pow(2,31) - 1) {
		return res - int(math.Pow(2,32))
	}else{
		return res
	}

}

func main(){
	arr := []int{1,2,1,1,2,2,4}
	fmt.Println(singleNumbers(arr))
	fmt.Println(singleNumbers([]int{1,2,1,1,2,2,-6}))
}