package main

import "fmt"

func search(nums []int, target int) int {
	i:= 0
	j:= len(nums) -1

	// 搜索右边界 right
	for i<=j {
		m := (i+j)/2
		if nums[m] <= target {
			i = m + 1
		}else{
			j = m - 1
		}
	}
	right := i
	// 若数组中无 target ，则提前返回
	if j >= 0 && nums[j] != target{
		return 0
	}
	// 搜索左边界 left
	i = 0
	for i<=j {
		m := (i+j)/2
		if nums[m] < target {
			i = m + 1
		}else{
			j = m - 1
		}
	}
	left := j
	return right - left - 1
}

func main(){
	nums :=[]int{5,7,7,8,8,10}
	fmt.Println(search(nums,8))
	fmt.Println(searchV2(nums,8))
}

func searchV2(nums []int, target int) int {
	if len(nums) <1 {
		return 0
	}
	i:= 0
	j:= len(nums) -1

	// 搜索右边界 right
	var m int
	for i<=j {
		m = (i+j)/2
		if nums[m] < target {
			i = m + 1
		}else if nums[m] > target{
			j = m - 1
		}else{
			break
		}
	}
	// 若数组中无 target ，则提前返回
	if nums[m] != target{
		return 0
	}
	right :=j
	left := i
	// 搜索右边最后一个出现的数
	for k:=m;k<=j;k++{
		if nums[k] !=target {
			break
		}
		right = k
	}
	// 搜索左边最后一个出现的数
	for h:=m;h>=i;h--{
		if nums[h] !=target {
			break
		}
		left = h
	}
	return right - left + 1
}