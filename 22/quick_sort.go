/**
归并排序
算法思想是分而治之(divide - conquer),每个递归过程分为三个步骤:
1) 分解:把待排序的n个元素的序列分解成两个子序列,每个子序列包括n/2个元素
2) 治理:对每个子序列分别调用归并排序(MergeSort),进行归并操作
3) 合并:合并两个排好序的子序列,生成排序结果
 */
package main

import "fmt"

func main() {
	arr := []int{5,4,9,8,7,6,0,1,3,2,2}
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 归并排序中的合并算法
func Merge(array []int, start int, mid int, end int) {
	n1 := mid - start +1
	n2 := end - mid
	L := make([]int,n1)
	R := make([]int,n2)
	// 拷贝前半部分数组
	for i,k:=0,start;i<n1;i,k=i+1,k+1{
		L[i] = array[k]
	}

	// 拷贝后半部分数组
	for i,k:=0,mid+1;i<n2;i,k=i+1,k+1{
		R[i] = array[k]
	}
	fmt.Println(L,R)

	var i,k,j int
	// 逐个扫描两部分数组然后放到相应的位置去
	for i,k,j=0,start,0;i<n1 && j<n2;k++{
		if L[i] > R[j]{
			array[k] = L[i]
			i++
		}else{
			array[k] = R[j]
			j++
		}
		if i<n1 {
			for j=i;j<n1;j,k=j+1,k+1{
				array[k] = L[j]
			}
		}
		if j<n2 {
			for i=j;j<n2;i,k=i+1,k+1{
				array[k] = R[i]
			}
		}
	}
}

// 归并排序
func MergeSort(array []int, start int, end int) {
	if start < end {
		mid := (start + end) / 2
		// 对前半部分进行排序
		MergeSort(array, start, mid)
		// 对后半部分进行排序
		MergeSort(array, mid+1, end)
		// 合并前后两部分
		Merge(array, start, mid, end)
	}
}
