package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 232, 34, 565, 656, 7}
	fmt.Println(arr)
	BubbleSortOptimize1(arr)
	fmt.Println(arr)
}

// 冒泡
func BubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		fmt.Println("数组不满足要求")
		return
	}
	n := 0 // 计数器
	// 外层循环：确定扫描的次数
	for i := 1; i <= len(arr)-1; i++ {
		// 内层循环：一轮扫描内，两两比较，进行交换
		for j := 0; j <= len(arr)-1-i; j++ { // - i 的原因是后面的元素已经被排序过
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
			n++
		}
	}
	fmt.Println("总的循环次数", n)
}

// 冒泡的优化1
func BubbleSortOptimize1(arr []int) {
	if arr == nil || len(arr) < 2 {
		fmt.Println("数组不满足要求")
		return
	}
	var sorted bool
	n := 0 // 计数器
	// 外层循环：确定扫描的次数
	for i := 1; i <= len(arr)-1; i++ {
		sorted = true
		// 内层循环：一轮扫描内，两两比较，进行交换
		for j := 0; j <= len(arr)-1-i; j++ { // - i 的原因是后面的元素已经被排序过
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				sorted = false
			}
			n++
		}
		if sorted {
			break
		}
	}
	fmt.Println("总的循环次数", n)
}

// 冒泡的优化2
func BubbleSortOptimize2(arr []int) {
	if arr == nil || len(arr) < 2 {
		fmt.Println("数组不满足要求")
		return
	}
	var sorted bool
	sortIndex := len(arr) - 1 - 1
	lastIndex := 0 // 记录最后一次交换的位置
	n := 0         // 计数器
	// 外层循环：确定扫描的次数
	for i := 1; i <= len(arr)-1; i++ {
		sorted = true
		// 内层循环：一轮扫描内，两两比较，进行交换
		for j := 0; j <= sortIndex; j++ { // - i 的原因是后面的元素已经被排序过
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				sorted = false
				lastIndex = j
			}
			n++
		}
		if sorted {
			break
		}
		sortIndex = lastIndex
	}
	fmt.Println("总的循环次数", n)
}
