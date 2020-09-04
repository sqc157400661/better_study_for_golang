package main

var a string
var done bool
func init(){
	print(1)
}
func init(){
	print(2)
}
func init(){
	print(3)
}
func setup(){
	a = "helllo,world 递归算法的生生世世生生世世水水水水水水水水水水水水水水水水水水水水水水水水水水水水水水水水水水水"
	done = true
}

func main(){
	go setup()
	for !done {}
	print(a)
}