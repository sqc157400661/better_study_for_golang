package main

import "fmt"

func twoSum(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		s := nums[i] + nums[j]
		if s > target {
			j -= 1
		} else if s < target {
			i += 1
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return nil
}

// 获取所有组合
func twoSumAll(nums []int, target int) [][]int {
	i := 0
	j := len(nums) - 1
	res := [][]int{}
	for i < j {
		s := nums[i] + nums[j]
		if s > target {
			j -= 1
		} else if s < target {
			i += 1
		} else {
			res = append(res, []int{nums[i], nums[j]})
			i += 1
			j -= 1
		}
	}
	return res
}

func main() {
	arr := []int{1, 2, 3, 5, 6, 7, 8, 9}
	fmt.Println(twoSum(arr, 8))
}
