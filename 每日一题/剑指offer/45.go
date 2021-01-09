package main

import (
	"fmt"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	res := make([]string,len(nums))
	// 转换成[]string
	for i:=0;i< len(nums);i++{
		res[i] = strconv.Itoa(nums[i])
	}
	FastSort(res,0,len(nums)-1)
	// []string 拼接成string
	return strings.Join(res,"")
}

// 注意这里需要转成[]string
func FastSort(nums []string,left,right int){
	if(left > right){
		return
	}
	base := nums[left] // 基准点
	l := left
	r := right
	for l!=r {
		// 这里的+号是字符串连接 不是数字相加 需要注意一下
		for (nums[r] + base) >= (base + nums[r]) && r>l {
			r--
		}
		for (nums[l] + base) <=  (base + nums[l]) && r>l {
			l++
		}
		if(l < r){
			nums[l],nums[r] = nums[r],nums[l]
			fmt.Println(nums[l],nums[r])
		}
	}
	// 相遇点和基准点互换
	nums[left] = nums[l]
	nums[l] = base
	FastSort(nums,left,l-1)
	FastSort(nums,l+1,right)
}

func main(){
	//arr := []int{3,30,34,5,9}
	arr := []int{3,30,5,34,9}
	fmt.Println(minNumber(arr))
}