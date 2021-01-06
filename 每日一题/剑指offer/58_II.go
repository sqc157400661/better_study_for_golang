package main

import (
	"fmt"
	"strings"
)

func reverseLeftWords(s string, n int) string {
	l:= len(s)
	if n<=0 {
		return ""
	}
	if n >=l{
		return s
	}
	return strings.Join([]string{s[n:],s[:n]},"")
}

func main(){
	test := "abcder"
	fmt.Printf("|%s|",reverseLeftWords(test,2))
}