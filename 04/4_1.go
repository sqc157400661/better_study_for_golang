/**
	学习切片
 */
package main

import "fmt"

func main() {
	var months = [...]string{1:"January","February","March","April","May","June","July","August","September","October","November","December"}
	fmt.Printf("The months  is %v\n",months) // 打印数组
	Q2 :=  months[4:7]
	summer := months[6:9]
	summer[0] = "June1";
	fmt.Printf("The Q2  is %v\n",Q2)
	fmt.Printf("The length of months is %d\n", len(months)) // 数组长度
	fmt.Printf("The length of Q2 is %d\n", len(Q2))
	fmt.Printf("The capacity of Q2 is %d\n", cap(Q2))
	fmt.Printf("The summer  is %v\n",summer)

	summer1 := summer[0:1] // 切片组成切片，注意切的时候只能向后移动 如Spring = summer[-1:] 会导致编译错误，不能在获取summer的前一个元素了
	fmt.Printf("The summer1  is %v\n",summer1)
}




