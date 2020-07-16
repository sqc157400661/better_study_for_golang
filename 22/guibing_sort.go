/**
归并排序 1
*/
package main

import "fmt"

func main() {
	arr := []int{5, 4, 9, 8, 7, 6, 0, 1, 3, 2, 2}
	MergeSort(arr, 0, len(arr))
	fmt.Println(arr)
	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort(list2, 0, len(list2))
	fmt.Println(list2)
}

// 归并排序中的合并算法
func Merge(array []int, start int, mid int, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := mid - start         // 左边数组的长度
	rightSize := end - mid          // 右边数组的长度
	newSize := leftSize + rightSize // 辅助数组的长度
	result := make([]int, 0, newSize)

	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := array[start+l] // 左边数组的元素
		rValue := array[mid+r]   // 右边数组的元素
		// 小的元素先放进辅助数组里
		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
	}

	// 将剩下的元素追加到辅助数组后面
	result = append(result, array[start+l:mid]...)
	result = append(result, array[mid+r:end]...)

	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		array[start+i] = result[i]
	}
	return
}

// 归并排序
// 自顶向下归并排序，排序范围在 [begin,end) 的数组
func MergeSort(array []int, start int, end int) {
	// 元素数量大于1时才进入递归
	if end - start >1 {
		// 将数组一分为二，分为 array[begin,mid) 和 array[mid,high)
		mid := start + (end-start+1)/2
		// 对前半部分进行排序
		MergeSort(array, start, mid)
		// 对后半部分进行排序
		MergeSort(array, mid, end)
		// 合并前后两部分
		Merge(array, start, mid, end)
	}
}
