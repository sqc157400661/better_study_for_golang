/**
学习基础数据类型 -字符串
*/

package main

import "fmt"

func main() {
	var st string
	fmt.Println(st)
	var str = "hello  world"
	fmt.Println(str)
	//Go 的字符串是由单个字节连接起来的 字符串实现基于 UTF-8 编码
	fmt.Println(str[0]) //打印h的utf-8编码值
	//字符串一旦赋值了，字符串就不能修改了:在 Go 中字符串是不可变的。
	//str[0] = 'w'

	// 可以使用双引号定义字符串, 双引号会识别转义字符
	var str1 = "自如\n伟大"
	fmt.Println(str1)
	/*
	   	在Go语言中，使用双引号书写字符串的方式是字符串常见表达方式之一，被称为字符串字面量（string literal）
	   	这种双引号字面量不能跨行,不能原生输出
		可以使用反引号定义，如下
	 */
	str2 := `
		func main() {
			fmt.Println(address)
		}
	`
	fmt.Println(str2)

	//4. 字符串拼接方式 :使用字符串拼接符“+”
	//当一行字符串太长时，需要使用到多行字符串，可以如下处理
	str3 := "xiao"
	str4 := str3 + "1" + "22" +
		"333" + "4444" +
		"55555"
	fmt.Println(str3,"\n",str4)
}
/*
 总结
	1、字符串即string类型，可以使用双引号""来定义字符串，字符串中可以使用转义字符来实现换行、缩进等效果
	2、字符串就是一串固定长度的字符连接起来的字符序列。
	3、反引号定义字符串，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	4、使用“+”进行字符串拼接
	5、字符串的零值是空字符串
 */
