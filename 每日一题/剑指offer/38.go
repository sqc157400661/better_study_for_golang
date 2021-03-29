package main

import (
	"fmt"
)

func permutation(s string) []string {
	res =  []string{}
	r := []byte(s)
	dfs(r,0)
	return res
}

var res []string

func dfs(c []byte,x int){
	if x == len(c) -1 {
		res = append(res, string(c))
	}
	dic := map[byte]bool{}
	for i := x; i < len(c); i++{
		if _,ok := dic[c[i]];ok {
			continue // 重复，因此剪枝
		}
		dic[c[i]] = true
		c[i], c[x] = c[x], c[i]  // 交换，将 s[i] 固定在第 x 位
		dfs(c,x + 1)  // 开启固定第 x + 1 位字符
		c[i], c[x] = c[x], c[i]  // 恢复交换
	}
}

func main()  {
	fmt.Println(permutation("aab"))
}