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
	// 当 x = len(c) - 1 时，代表所有位已固定（最后一位只有 11 种情况），则将当前组合 c 转化为字符串并加入 res ，并返回
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
		dfs(c,x + 1)  // 进入下一层递归 开启固定第 x + 1 位字符
		//返回时交换回来，这样保证到达第1层的时候，一直都是abc。这里捋顺一下，开始一直都是abc，那么第一位置总共就3个交换
		//分别是a与a交换，这个就相当于 x = 0, i = 0;
		//     a与b交换            x = 0, i = 1;
		//     a与c交换            x = 0, i = 2;
		//就相当于上图中开始的三条路径
		//第一个元素固定后，每个引出两条路径,
		//     b与b交换            x = 1, i = 1;
		//     b与c交换            x = 1, i = 2;
		//所以，结合上图，在每条路径上标注上i的值，就会非常容易好理解了
		c[i], c[x] = c[x], c[i]  // 恢复交换
	}
}

func main()  {
	fmt.Println(permutation("aab"))
}