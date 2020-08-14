package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 232, 34, 565, 656, 7}
	fmt.Println(len(arr))
	QuickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}

// 快速排序
func QuickSort(arr []int,left int,right int){
	if(left > right){
		return;
	}
	key := arr[left] //基准点
	l := left
	r := right
	for l!=r {
		// 移动右边，找到比基准值大的  找到后停止移动
		for arr[r] >= key && l<r{
			r--
		}
		// 移动左边 找个比基准值小的 找到后停止移动
		for arr[l] <= key && l<r{
			l++
		}

		// 左右2边都停止移动后 交换左右标记
		if(l < r){
			t := arr[l]
			arr[l] = arr[r]
			arr[r] = t
		}
	}
	// 左右2边相遇，相遇点于基准点互换
	arr[left] = arr[l];
	arr[l] = key;
	QuickSort(arr,left,l-1)
	QuickSort(arr,l+1,right)
}