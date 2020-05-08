/**
	下面代码输出什么？
 */
package main

import (
	"fmt"
)

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

/**
	[0 0 0 0 0 1 2 3]

	考点：make默认值和append
	make初始化是有默认值的哦，此处默认值为0

 */