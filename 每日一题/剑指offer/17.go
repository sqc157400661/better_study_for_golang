package main

import (
	"fmt"
	"strconv"
	"strings"
)

var nums = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
func printNumbers(n int) []int {
	var res = ""// 保存结果的变量
	var dfsNums = make([]string,n) // 定义长度为 n 的字符列表
	dfs(0, n, dfsNums,&res) // 开启全排列递归
	res = res[1:] // 删除头部多余的逗号
	resStrSlice := strings.Split(res,",")
	result :=[]int{}
	for _,v:= range resStrSlice{
		tmp,e := strconv.Atoi(v)
		if tmp==0 || e !=nil{
			continue
		}
		result = append(result,tmp)
	}
	return result
}

// 全排列
func dfs(x, n int, dfsNums []string,res *string){
	if x == n {
		// 终止条件：已固定完所有位
		*res =*res + ","+ strings.Join(dfsNums,"")
		return
	}
	for i := 0; i < 10; i++ { // 遍历 ‘0‘ - ’9‘
		dfsNums[x] = nums[i]  // 固定第 x 位为 i
		dfs(x+1, n, dfsNums,res) // 开启固定第 x + 1 位
	}
}

func main() {
	fmt.Println(printNumbers(2))
}
