package main

import "fmt"

func main (){
	arr := []int{7,3,19,40,4,7,1}
	fmt.Println("数组：",arr)
	max,min := GetMaxAndMinRE(arr,0,len(arr)-1)
	fmt.Println("变形的分治法")
	fmt.Println("max=",max)
	fmt.Println("min=",min)
}

func GetMaxAndMinRE(arr []int,start,end int) (max,min int){
	if arr == nil {
		return 0,0
	}
	// 求中点
	m :=(start + end)/2

	// start与end之间只有一个元素
	if start == end {
		max,min = arr[start],arr[end]
		return
	}
	// start与end之间只有二个元素
	if start + 1 == end {
		if arr[start] >= arr[end]{
			max,min = arr[start],arr[end]
		}else{
			max,min = arr[end],arr[start]
		}
		return
	}
	// 递归计算左半部分
	lmax,lmin := GetMaxAndMinRE(arr,start,m)
	// 递归计算右半部分
	rmax,rmin := GetMaxAndMinRE(arr,m+1,end)

	// 总的最大值
	if lmax > rmax {
		max = lmax
	}else{
		max = rmax
	}
	// 总的最小值
	if lmin < rmin {
		min = lmin
	}else{
		min = rmin
	}
	return
}