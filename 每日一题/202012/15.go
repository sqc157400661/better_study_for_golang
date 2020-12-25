package main

import (
	"fmt"
)

func main(){
	//a := []int{2, 3, 1, 0, 2, 5, 3}
	//fmt.Println(findRepeatNumber(a))

	b := []int{0, 3, 1, 0}
	fmt.Println(findRepeatNumber(b))
}

func findRepeatNumber(nums []int) int {
	var zern =0 // 对0特殊处理 统计数量
	for i:=0;i<len(nums);i++{
		index := nums[i]
		if index==0 {
			zern++
		}
		if zern > 1 {
			return 0
		}
		if index <0 {
			index *= -1
		}
		if nums[index] < 0 {
			return index
		}else{
			nums[index] *= -1
		}
	}
	return -1
}