/**
	学习基础数据类型 -整形
 */
package main
import (
	"unsafe"
	"fmt"
)


func main() {
	var i int = 1
	fmt.Println("i=", i)

	var j int8 = 127
	fmt.Println("i=", j)


	var a int = 233
	fmt.Println("a=", a)
	var b uint = 233233
	fmt.Println("b=", b)
	var c byte = 255
	fmt.Println("c=", c)

	// 使用细节
	// 1. Go各整数类型分有无符号。int uint 的大小和系统有关
	// 2. Go的整型默认声明为int型

	var n1 int32 = 1
	fmt.Printf("n1的类型%T 占用的字节数是%d\n ", n1,unsafe.Sizeof(n1))
	// 查看某个变量占用的字节大小和数据类型
	var n2 int64 = 10
	fmt.Printf("n2的类型是 %T n2占用的字节数是%d", n2, unsafe.Sizeof(n2))
	// Go程序中整型变量在使用时，遵循保小不保大的原则。即再保证程序运行正常的情况下，尽量使用占用空间小的数据类型
}

/*
	整数：
	int8（-128 -> 127）
	int16（-32768 -> 32767）
	int32（-2,147,483,648 -> 2,147,483,647） 与rune等价 可以表示一个unicode编码 这两个名称可以互换使用
	int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
	无符号整数：
	uint8（0 -> 255） 与byte等价 byte类型一般用于强调数值是一个原始的字符数据。
	uint16（0 -> 65,535）
	uint32（0 -> 4,294,967,295）
	uint64（0 -> 18,446,744,073,709,551,615）

	整型的零值为 0
*/

