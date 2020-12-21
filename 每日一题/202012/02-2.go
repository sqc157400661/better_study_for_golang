package main

import "fmt"

/*
	类快排思路 本次是从小到大排
*/
func FindSmallKV2(arr []int, left, right, k int) int {
	//选择 基准点
	base := arr[left]
	l := left
	r := right
	for l != r {
		// 基数在左边先移动右游标。在右边先移动左游标。
		// 先移动右边边游标
		for arr[r] >= base && l < r {
			r--
		}

		// 后移动左边游标
		for arr[l] <= base && l < r {
			l++
		}
		if l < r {
			t := arr[l]
			arr[l] = arr[r]
			arr[r] = t
		}
	}
	// 游标相遇 相遇点与基准点互换
	arr[left] = arr[l]
	arr[l] = base
	subOffset := l - left

	if subOffset == k-1 {
		//递归出口
		//return arr[0:l+1]  这个可以返回前k个小的数
		return arr[l]
	} else if subOffset > k-1 {
		return FindSmallKV2(arr, left, subOffset-1, k)
	} else {
		return FindSmallKV2(arr, subOffset+1, right, k-subOffset-1)
	}
}

func main() {
	arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	re := FindSmallKV2(arr, 0, len(arr)-1, 3)
	fmt.Println(re)
}
