/**
如何从给定的车票中找出旅程路线
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("如何从给定的车票中找出旅程路线")
	input := map[string]string{"西安": "成都", "北京": "上海", "大连": "西安", "上海": "大连"}
	PrintResult(input)
}

func PrintResult(input map[string]string) {
	// 用来存储吧input的键与值调换后的信息
	reverseInput := map[string]string{}
	for k, v := range input {
		reverseInput[v] = k
	}
	// 找到起点
	start := ""
	for k, _ := range input {
		if _, v := reverseInput[k]; !v {
			start = k
			break
		}
	}
	if start == "" {
		fmt.Println("输入不合理")
	} else {
		tmpStart := start
		tmpTo := input[tmpStart]
		for tmpTo != "" {
			fmt.Println(tmpStart + " -> " + tmpTo + " ")
			tmpStart = tmpTo
			if _,tmp := input[tmpStart];!tmp{
				tmpTo =""
				break
			}
			tmpTo = input[tmpStart]
		}
	}

}
