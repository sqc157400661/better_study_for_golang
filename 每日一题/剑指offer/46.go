package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	str := strconv.Itoa(num)
	a := 1
	b := 1
	c := 0
	for i := 2; i <= len(str); i++ {
		newnum := str[i-2 : i]
		if newnum >= "10" && newnum <= "25" {
			c = a + b
		} else {
			c = a
		}
		b = a
		a = c
	}
	return a
}

func main() {
	fmt.Println(translateNum(12258))
}
