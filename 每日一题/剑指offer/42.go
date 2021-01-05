package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += max(nums[i-1], 0)
		res = max(res, nums[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
