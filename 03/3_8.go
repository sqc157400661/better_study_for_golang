/**
	学习函数 - 可变参数
 */
package main

import "fmt"

func main() {
	/**
		说明：
		1、变量 who 的值为 []string{"shiqc", "gaoh", "liupt"}
		2、如果参数被存储在一个`slice` 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数调用变参函数。
	 */
	Greeting("ziroom","shiqc","gaoh","liupt")



	who := []string{"wangwt","guiz","liy"}
	Greeting("ziroom",who...)


}

func Greeting(prefix string, who...string){
	// who的值是个切片
	fmt.Println("prefix:"+prefix,who," Hi!!!!!")
}
/*
	总结：
	1、参数数量可变的函数称为为可变参数函数  【典型的例子就是fmt.Printf和类似函数】  格式： 参数传递：...type
	2、注意的是不定长参数，必须放在参数列表的最后
	3、

*/



