/*
你必须在源文件中非注释的第一行指明这个文件属于哪个包，
如：package main。package main 表示一个可独立执行的程序，
每个 Go 应用程序都包含一个名为 main 的包。
 */

package main

import "fmt"

func init(){
	fmt.Println("hello ziroom")
}

func main() {
	fmt.Println("hello world")
}

