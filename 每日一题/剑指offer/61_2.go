package main

import "sort"

func isStraight(nums []int) bool {
	sort.Ints(nums)
	joker := 0

	for i:=0;i<4;i++{

		if nums[i] == 0 {
			// 统计大小王数量
			joker++
		}else if nums[i] == nums[i + 1]{
			//若有重复，提前返回 false
			return false
		}

	}
	return nums[4] - nums[joker] < 5 // 最大牌 - 最小牌 < 5 则可构成顺子
}
