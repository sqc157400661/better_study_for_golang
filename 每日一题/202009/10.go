package main

import (
	"fmt"
	"regexp"
)

func main(){
	reg := regexp.MustCompile(`(?i)^m([txq])_(\d{8,}$)`)
	buf := "mt_46495465"
	//根据规则提取关键信息
	result1 := reg.FindStringSubmatch(buf)
	fmt.Printf("%t",result1[2])
}
