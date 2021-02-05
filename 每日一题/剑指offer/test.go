package main

import "fmt"

type A struct {
	B
	C
}
type B struct {
	BB string
}
type C struct {
	BB string
}
func main(){
	a := get()
	fmt.Println(a["aa"])
}

func get() map[string]string {
	return nil
}
