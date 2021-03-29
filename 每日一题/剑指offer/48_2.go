package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	dic := map[byte]int{}
	var res int
	i := -1
	var ok bool
	for j := 0; j < len(s); j++ {
		_, ok = dic[s[j]] // 获取索引 v
		if ok {
			i = max(dic[s[j]], i)
		}
		dic[s[j]] = j // 更新哈希表
		res = max(res, j - i) // 更新结果
		ok = false
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main(){
	fmt.Println(lengthOfLongestSubstring("au"))
}
