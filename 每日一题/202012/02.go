package main

import "fmt"

// 使用冒泡法查找第k个元素，只需要冒泡k次
func FindSmallK(arr []int, k int) {
	for i := 0; i < k; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				t := arr[j]
				arr[j] = arr[min]
				arr[min] = t
			}
		}
		fmt.Println(arr[min])
	}
	fmt.Println(arr)
}

func main() {
	FindSmallK([]int{4, 5, 1, 6, 2, 7, 3, 8}, 3)
}
