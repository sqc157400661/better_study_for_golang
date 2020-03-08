/**
	学习控制结构-跳转控制
 */
package main

import (
	"fmt"
)

func main() {
	OuterLoop:
		for i := 0; i < 2; i++ {
			for j := 0; j < 5; j++ {
				switch j {
				case 2:
					fmt.Println("continue start",i, j)
					continue
					fmt.Println("continue after",i, j)
				case 3:
					fmt.Println(i, j)
					break OuterLoop
				}
			}
		}
}
/*
	总结：
	1、break默认会跳出最近的for循环
	2、break语句出现在多层嵌套的语句块中时，可以通过标签指明要终止的是哪一层语句块
	3、关键字 continue 只能被用于 for 循环中
	4、continue 语句用于结束本次循环，继续执行下一次循环。
	5、continue 语句出现在多层嵌套的循环语句体中时， 可以通过标签指明要跳过的是哪一层循环, 这 个和前面的 break 标签的使用的规则一样.

*/



