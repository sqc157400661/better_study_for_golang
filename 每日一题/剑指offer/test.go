package main

import "fmt"

func main() {
	t1 := []int{1,3}
	testSlice(t1)

	t2 := map[int]int{1:2}
	testMap(t2)

	fmt.Println(t1,t2)
}
func testSlice(s []int){
	s = append(s,5,6)
}
func testMap(s map[int]int){
	s[55] = 55
}