package main

import "fmt"

/**
 * 摩尔投票法
 * @param numbers int整型一维数组
 * @return int整型
 */
func MoreThanHalfNum_Solution( numbers []int ) int {
	var x,votes int
	for i:=0;i< len(numbers);i++{
		if votes ==0 {
			x = numbers[i]
			votes = 1
		}else{
			if x == numbers[i]{
				votes +=1
			}else{
				votes -=1
			}
		}
	}
	return x
}

func main(){
	var arr = []int{1,2,3,2,2,2,5,4,2}
	var arr1 = []int{3,3,4,8,8,8,8}
	fmt.Println(MoreThanHalfNum_Solution(arr))
	fmt.Println(MoreThanHalfNum_Solution(arr1))
}