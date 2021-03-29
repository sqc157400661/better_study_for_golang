package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	dic := map[byte]int{}
	var res, tmp int
	for j := 0; j < len(s); j++ {
		i, ok := dic[s[j]] // 获取索引 v
		if !ok {
			i = -1
		}
		dic[s[j]] = j // 更新哈希表
		if tmp < j-i {
			tmp++
		} else {
			tmp = j - i
		}
		res = max(res, tmp) // max(dp[j - 1], dp[j])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(lengthOfLongestSubstring("au"))
}
