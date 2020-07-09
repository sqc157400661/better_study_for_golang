package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 232, 34, 565, 656, 7}
	fmt.Println(len(arr))
	QuickSort(arr,0,8)
	fmt.Println(arr)
}

// 快速排序
func QuickSort(arr []int,left int,right int){
	if(left > right){
		return;
	}
	key := arr[left]
	l := left
	r := right
	for l!=r {
		for arr[r] >= key && l<r{
			r--
		}
		for arr[l] >= key && l<r{
			l++
		}
		if(l < r){
			t := arr[l];
			arr[l] = arr[r];
			arr[r] = t;
		}
	}
	arr[left] = arr[l];
	arr[l] =key;
	QuickSort(arr,left,l-1);
	QuickSort(arr,l+1,right);
}

