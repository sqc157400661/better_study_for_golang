/**
数据类型转换
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//一、整形不同类型间的转换
	var int1 int8 = 22
	var int2 int16 = 23
	//int2 = int1 + int2 // 报错 int8类型不能被隐式转换成int16
	int2 = int16(int1) + int2
	fmt.Println("int1=", "int2=",int1,int2)
	//转换的结果是按 溢出处理，和我们希望的结果不一样
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2)

	//二、基本数字类型和 string 的转换
	// 1、整形转字符串
	var i int32 = 100
	str := fmt.Sprintf("%d", i)  //  fmt.Sprintf("%参数", 表达式)
	fmt.Printf("Sprintf:str type %T str %q\n", str, str)
	str = strconv.FormatInt(int64(i), 10)  // 使用strconv包的函数进行转换 注意参数是int64
	fmt.Printf("FormatInt:str type %T str %q\n", str, str)

	// 2、浮点数转字符串
	var f float32 = float32(i)
	str2 := fmt.Sprintf("%f", f)
	fmt.Printf("Sprintf:str2 type %T str %q\n", str2, str2)
	str2 = strconv.FormatFloat(float64(f), 'f', 10, 64) //f格式 10 表示小数保留10位，64表示这个小数时float64
	fmt.Printf("FormatInt:str2 type %T str %q\n", str2, str2)

	// 3、布尔转字符串
	isTrue := false
	str3 := fmt.Sprintf("%t", isTrue)
	fmt.Printf("Sprintf:str3 type %T str %q\n", str3, str3)
	str3 = strconv.FormatBool(isTrue)
	fmt.Printf("FormatInt:str3 type %T str %q\n", str3, str3)

	// 4、字符转字符串
	myChar := 'b'
	str4 := fmt.Sprintf("%c", myChar)
	fmt.Printf("Sprintf:str4 type %T str %q\n", str4, str4)
	str4 = strconv.Itoa(int(myChar))
	fmt.Printf("FormatInt:str4 type %T str %q\n", str4, str4)

	fmt.Printf("------------------------------------------------------------------------\n")

	//三、string 类型转基本类型
	var str5 string = "true"
	var b bool
	b, _ = strconv.ParseBool(str5) //ParseBool会返回两个值(value bool, err error) 空白标识符 _ 也被用于抛弃值
	fmt.Printf("b type %T b=%v\n", b, b)

	var str6 string = "123456789"
	var n1 int64
	n1, _ = strconv.ParseInt(str6, 10, 64)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)


	var str7 string = "123456.2345"
	var f1 float64
	f1, _ = strconv.ParseFloat(str7, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)

	//string 转基本数字类型的注意事项
	//在将 String 类型转成 基本数据类型时，要确保 String 类型能够转成有效的数据，
	//比如 我们可以 把 "123" , 转成一个整数，但是不能把 "hello" 转成一个整数，如果这样做，Golang 直接将其转成 0 ，
	//其它类型也是一样的道理. float => 0 bool => false

}


/*
 总结
	1. Go 在不同类型的变量之间赋值时需要显式转换。也就是说Golang中数据类型不能自动转换。
	2. 类型转换常见表达式： T(v) 将值 v 转换为类型 T。
	3、Go中，数据类型的转换可以是从表示范围小-->表示范围大，也可以范围大--->范围小
	4、在转换中，比如将 int64 转成 int8 【-128---127】 ，编译时不会报错，只是转换的结果是按 溢出处理，和我们希望的结果不一样。 因此在转换时，需要考虑范围.
	5、被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化!
*/
