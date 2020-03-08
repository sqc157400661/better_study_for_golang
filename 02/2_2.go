/**
学习基础数据类型 -字符
*/

package main

import "fmt"

func main() {
	var ch byte = 65 // var ch byte = '\x41'    // （\x 总是紧跟着长度为 2 的 16 进制数）
	fmt.Println(ch)
	//说明1. 如果我们需要按照字符输出，可以使用%c格式化输出
	fmt.Printf("%c\n",ch)

	//说明2. 如果我们保存的字符在ASCII表中，我们可以直接使用byte保存
	var c1 byte = 'a'
	var c2 byte =97
	fmt.Printf("c1=%c c2=%c\n", c1, c2)
	//说明3. 如果我们保存的字符值大于255 这时我们可以使用int保存
	var c3 int = '美'
	fmt.Printf("c3=%c c3对应的码值=%d\n", c3, c3)

	//说明4. 字符类型是可以进行运算的，相当于一个整数，因为它都对应有 Unicode 码.
	var n1 = 10 + 'a'
	fmt.Println("n1=", n1)


}

/*
总结：
1. Go语言的字符有以下两种：
一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。

2. 字符是使用单引号扩起来的单个字符
3. Go中允许使用转义字符'\'来将其后的字符转变为特殊字符型常量。var c4 char = '\n'
4.  Go 语言的字符使用 UTF-8 编码 ，如果想查询字符对应的 utf8 码值 http://www.mytju.com/classcode/tools/encode_utf8.asp
英文字母-1个字节 汉字-3个字节
5. 在 Go 中，字符的本质是一个整数，直接输出时，是该字符对应的 UTF-8 编码的码值。
6. 可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的 unicode 字符
 */
