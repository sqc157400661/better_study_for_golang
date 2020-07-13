/**
	如何从数组中找出满足a+b=c+d的两个数对
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("如何从数组中找出满足a+b=c+d的两个数对")
	arr := []int{3,4,7,10,20,9,8}
	FindPairs(arr)
}

func FindPairs(arr []int)  {
	var reMap map[int][][2]int
	reMap = map[int][][2]int{}
	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			sum := arr[i] + arr[j]
			if reMap[sum] == nil || len(reMap[sum]) <1{
				reMap[sum] = [][2]int{{arr[i],arr[j]}}
			}else{
				reMap[sum] = append(reMap[sum], [2]int{arr[i],arr[j]})
			}
		}
	}
	for k,_ := range reMap{
		if(len(reMap[k]) >1){
			fmt.Println(reMap[k],k)
		}
	}
}
