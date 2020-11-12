package main

//void SayHello1(_GoString_ s);
import "C"
import "fmt"

func main(){
	C.SayHello1("Hello,World\n")
}

//export SayHello1
func SayHello1(s string){
	fmt.Print(s)
}