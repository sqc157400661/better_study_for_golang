package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 232, 34, 565, 656, 7}
	fmt.Println(arr)
	InsetionSort(arr)
	fmt.Println(arr)
}

// 插入排序
func InsetionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		fmt.Println("数组不满足要求")
		return
	}
	// 0位置时，只有1个元素，我们认为他已经是有序的
	for i := 1; i <= len(arr)-1; i++ {
		// 若当前元素比前一个元素小，查看 i 位置的元素应该插入到 0-（i-1）位置的何处
		if arr[i] < arr[i-1] {
			temp := arr[i]
			for j := i; j >= 0; j-- {
				if j > 0 && arr[j-1] > temp {
					arr[j] = arr[j-1]
				} else {
					arr[j] = temp
					break
				}
			}
		}
	}
}

