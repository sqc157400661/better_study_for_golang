/**
	学习复合数据类型-map
 */
package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	// 假设我们想获取一个 map 类型的切片，我们必须使用两次 make() 函数，第一次分配切片，第二次分配切片中每个 map 元素
	items := make([]map[int]int,5)
	for i:= range items {
		items[i] = make(map[int]int, 1) // 没有这个会报错
		items[i][1] = 2
	}
	items2:= []map[int]int{map[int]int{1:2},map[int]int{1:2,}}
	fmt.Println(items2)

	/*
		map 的排序
		map 默认是无序的，如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序（使用 sort 包），
		然后可以使用切片的 for-range 方法打印出所有的 key 和 value。
	 */
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}




