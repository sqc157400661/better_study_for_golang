package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	re := ""
	i := len(s)-1
	j := len(s)-1
	for i>=0{
		// 去掉单词 间隔的空格
		for j>=0 && s[j] ==' '{
			j--
			i--
		}
		// 左移动指针 找到空格点
		for i>=0 && s[i] !=' '{
			i--
		}

		if i<j{
			re += s[i+1:j+1] +" "
			j=i
		}
	}
	if len(re) >0 {
		return re[0:len(re)-1]
	}
	return ""
}

// 优化版
func reverseWordsV2(s string) string {
	re := []string{}
	i := len(s)-1
	j := len(s)-1
	for i>=0{
		// 去掉单词 间隔的空格
		for j>=0 && s[j] ==' '{
			j--
			i--
		}
		// 左移动指针 找到空格点
		for i>=0 && s[i] !=' '{
			i--
		}

		if i<j{
			re = append(re,s[i+1:j+1])
			j=i
		}
	}
	if len(re) >0 {
		return strings.Join(re," ")
	}
	return ""
}


func main(){
	test := "   "
	fmt.Printf("|%s|",reverseWords(test))
	fmt.Printf("|%s|",reverseWordsV2(test))
}