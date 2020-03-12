/**
	学习复合数据类型-map
 */
package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice":31,
		"char" : 32,
	}
	ages["sug"] = ages["sug"] + 1  //key不在map中也没有关系；如果一个查找失败将返回value类型对应的零值
	//fmt.Println(&ages["sug"])  // map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作
	fmt.Println(ages)

	//var ages1 map[string]int
	ages1 := make(map[string]int,10)
	//ages1 := map[string]int{} // 注意这个和var ages1 map[string]int 有啥不一样
	//ages1["carol"] = 21  // 在向map存数据前必须先创建map var声明没有赋初始值直接这样操作会panic: assignment to entry in nil map
	fmt.Println(ages1)
	fmt.Println(ages1 == nil) //map类型的零值是nil
	fmt.Println(len(ages1) == 0)
}
/*
	总结：
	1、和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。
    2、我们也可以指定map的初始容量：make(map[keytype]valuetype, cap)
	3、虽然map会自动加。但是出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。
 */



