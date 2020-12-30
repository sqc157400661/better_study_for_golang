package main

import "fmt"

func missingNumber(nums []int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		m := (i + j) / 2
		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}

func main() {
	arr := []int{0, 1, 3, 4, 5, 6}
	fmt.Println(missingNumber(arr))
}
