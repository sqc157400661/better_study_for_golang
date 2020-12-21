package main

import "fmt"

func main() {
	arr1 := []int{1,2,3,4}
	arr2 := []int{1,5,6,7,8}
	re :=MergeSort(arr1,arr2)
	fmt.Println(re)
}

//  采用归并排序的思路
func MergeSort(arr1,arr2 []int)[]int{
	m := len(arr1)
	n := len(arr2)
	result := make([]int,0,m+n)
	i,j:=0,0
	for(i<m && j<n){
		if arr1[i] < arr2[j]{
			result = append(result,arr1[i])
			i++
		}else{
			result = append(result,arr2[j])
			j++
		}
	}
	if i==m && j<n{
		result = append(result,arr2[j:]...)
	}else if(i<m && j==n){
		result = append(result,arr1[i:]...)
	}
	return result
}