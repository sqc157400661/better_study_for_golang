package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := []int{1,3}
	re,_ := json.Marshal(str)
	fmt.Println(string(re))

	res :=  []int{}
	json.Unmarshal(re,&res)
	fmt.Println(res)

}