package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 232, 34, 565, 656, 7}
	fmt.Println(arr)
	SelectSort(arr)
	fmt.Println(arr)
}

// 选择排序
func SelectSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		fmt.Println("数组不满足要求")
		return
	}
	// 从0开始查找
	for i :=0 ; i <=len(arr) - 1; i++ {
		flag := i
		for j := i + 1; j <= len(arr) - 1; j++ {
			if arr[flag] > arr[j] {
				flag  = j
			}
		}

		if flag != i {
			temp := arr[i]
			arr[i] = arr[flag]
			arr[flag] = temp
		}
		fmt.Println(arr) //为了看具体排序的过程

	}
}

