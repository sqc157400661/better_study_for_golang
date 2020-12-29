package main

import "fmt"

// 首尾双指针
func exchangeV1(nums []int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		// 左边遇到奇数
		if (nums[left] & 1) != 0 {
			left++
			continue
		}

		// 右边边遇到偶数
		if (nums[right] & 1) == 0 {
			right--
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}

// 快慢双指针
func exchangeV2(nums []int) []int {
	low := 0
	fast := 0
	for fast < len(nums) {
		if (nums[fast] & 1) != 0 {
			nums[low], nums[fast] = nums[fast], nums[low]
			low++
		}
		fast++
	}
	return nums
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(exchangeV1(arr))
	fmt.Println(exchangeV2(arr))
}
