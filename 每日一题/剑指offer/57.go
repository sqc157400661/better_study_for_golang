package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	var result = [][]int{}
	//起始 l=1,r=2
	l:=1
	r:=2
	for(l<r){
		sum := (l+r)*(r-l+1)/2
		if sum < target {
			r++
		}else if sum > target{
			l++
		}else{
			tmp := []int{}
			for i:=l;i<=r;i++{
				tmp = append(tmp,i)
			}
			result = append(result,tmp)
			l++ // 包含l的只能有一个连续数
		}
	}
	return result
}

func main(){
	fmt.Println(findContinuousSequence(9))
}