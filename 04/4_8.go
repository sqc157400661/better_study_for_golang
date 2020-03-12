/**
	学习复合数据类型-map
 */
package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice":31,
		"char" : 32,
		"char1" : 32,
		"char2" : 32,
		"char3" : 32,
	}

	/*
		说明：
		isPresent 返回一个 bool 值
		如果 key1 存在于 map1，val1 就是 key1 对应的 value 值，并且 isPresent 为 true；
		如果 key1 不存在，val1 就是一个空值，并且 isPresent 会返回 false。
	 */
	//var1,isPresent := ages["sug"]
	var1,isPresent := ages["char"]
	fmt.Println(var1,isPresent)
	// 常见判断key是否存在格式： 不能直接判断 ages["alice"]== nil  key不存在，那么将得到value对应类型的零值  区分已经存在的0，和不存在而返回零值的0
	if var2,isPresent2 := ages["alice"]; isPresent2 {
		fmt.Println(var2)
	}
	fmt.Println(ages)

	//从map中删除元素
	delete(ages, "alice") // remove element ages["alice"]
	delete(ages, "alice111") // 刪除不存在的key 不会报错
	var ages1 map[string]int
	delete(ages1, "alice111") // 当map是nil的时候 刪除不存在的key 不会报错
	//ages1["ww"]  =1 // 当map是nil的时候 存入元素将导致一个panic异常 向map存数据前必须先创建map
	fmt.Println(ages)

	//for-range遍历
	for key, value := range ages {
		fmt.Printf("key is: %v - value is: %v\n", key, value)
	}
}
/*
	总结：
	1、map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。但是向一个nil值的map存入元素将导致一个panic异常。
    2、从 map1 中删除 key1：delete(map1, key1) 如果 key1 不存在，该操作不会产生错误
	3、val1, isPresent := map1[key1]格式来判断map中的key是否存在。
	4、和数组切片遍历一样，可以for-range遍历，但是需要注意的是遍历的顺序是随机的
	5、和slice一样，map之间也不能进行相等比较；唯一的例外是和nil进行比较，要判断两个map是否包含相同的key和value，我们必须通过一个循环实现
 */



