package main

import (
	"fmt"
	"net/url"
)

func main(){
	aa := "%D0%9C%C2%B0"
	fmt.Println(url.QueryUnescape(aa))
}
