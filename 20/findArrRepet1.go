package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 2, 5, 4}
	fmt.Println("环形相遇法")
	fmt.Println(FindDupByLoop(arr))
}
func FindDupByLoop(arr []int) int {
	if arr == nil {
		return -1
	}
	slow := 0
	fast := 0
	for ok := true; ok; ok = slow != fast {
		fast = arr[arr[fast]]
		slow = arr[slow]
	}
	head := 0
	// 找到入口点
	for ok := true; ok; ok = slow != head {
		head = arr[head]
		slow = arr[slow]
	}
	return slow
}
