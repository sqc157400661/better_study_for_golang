/**
	在utf8字符串判断是否包含指定字符串，并返回下标。 “北京天安门最美丽” , “天安门” 结果：2
 */
package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(Utf8Index("北京天安门最美丽", "天安门"))
	//fmt.Println(strings.Index("北京天安门最美丽", "男"))
	//fmt.Println(strings.Index("", "男"))
	//fmt.Println(Utf8Index("12ws北京天安门最美丽", "天安门"))
}

func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	fmt.Println(asciiPos)
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}


/**
	考点：
	 func (b *Reader) ReadRune() (r rune, size int, err error)  从字符串中读出一个 UTF8 编码的字符并返回 size是字节大小

	func Index(s, sep string) int   返回子串sep在字符串s中第一次出现的索引值，不在的话返回-1.


 */