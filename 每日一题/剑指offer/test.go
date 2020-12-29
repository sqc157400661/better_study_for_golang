package main

import (
	"fmt"
	"strings"
)

func main(){
	aa := []string{"1","9"}
	res := strings.Join(aa,",")
	fmt.Println(res)
}
