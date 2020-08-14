package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 2, 5, 4}
	fmt.Println("数据映射法")
	fmt.Println(FindDupByMap(arr))
}
func FindDupByMap(arr []int) int {
	if arr == nil {
		return -1
	}
	len := len(arr)
	index := 0
	for i := 0; i < len; i++ {
		// 数组中元素的值，只能小于len，否则会溢出
		if arr[index] >= len {
			return -1
		}
		if arr[index] < 0 {
			break
		}
		// 访问过，通过变相番薯的方法进行标记
		arr[index] *= -1
		fmt.Println(arr[index])
		// index的后继为array[index]
		index = arr[index] * -1
		if index >= len {
			fmt.Println("数组中有非法数字")
			return -1
		}
	}
	return index
}
