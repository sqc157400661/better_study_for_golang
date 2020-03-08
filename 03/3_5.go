/**
	学习控制结构-for-range 结构
 */
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string ="hello world 呵呵"
	for i:=0;i<len(str);i++{
		fmt.Printf("%c \n",str[i])
	}
	for index,val := range str{
		fmt.Printf("index=%d value=%c\n", index, val)
	}
	/*
		 注意：如果我们字符串中含有中文，那么配普通的for循环会出现乱码，
			原因是因为for循环遍历是按照字节来遍历的，中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。
			如何解决？可以将str转成[]rune切片
	*/

	var str1 = "呵呵"
	fmt.Println("len(str):", len(str1)) // 输出6 golang中string底层是通过byte数组实现的 len 实际是在按字节长度计算
	fmt.Println("chinese word length", len([]rune(str1))) // 输出chinese word length 2
	fmt.Println("chinese word length", utf8.RuneCountInString(str1)) // 输出chinese word length 2
	/*
		 解释：
			byte 等同于int8，常用来处理ascii字符
		    rune 等同于int32,常用来处理unicode或utf-8字符
	*/

	// 对上面中文乱码问题的解决
	str2 := []rune(str)
	for i:=0;i<len(str2) ;i++  {
		fmt.Printf("%c \n",str2[i])
	}
}
/*
	总结：
	1、
*/



