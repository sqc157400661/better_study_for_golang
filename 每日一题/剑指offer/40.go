package main

import "fmt"

func getLeastNumbers(arr []int, k int) []int {
	return FindSmallK(arr,0,len(arr)-1,k)
}

func FindSmallK(arr []int,left int,right int,k int) []int{
	if(left > right){
		return arr;
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
	if (k == l) {
		// 正好找到最小的 k(m) 个数
		return arr[:k];
	} else if (k < l) {
		// 最小的 k 个数一定在前 m 个数中，递归划分
		return FindSmallK(arr, left, l-1, k);
	} else {
		// 在右侧数组中寻找最小的 k-m 个数
		return FindSmallK(arr, l+1, right, k);
	}
}



func main(){
	arr := []int{3,2,1}
	fmt.Println(getLeastNumbers(arr,2))
	arr = []int{0,1,2,4,2,2,3,1,4}
	fmt.Println(getLeastNumbers(arr,8))
	//fmt.Println(getLeastNumbers(arr,2))
}