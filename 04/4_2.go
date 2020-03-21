/**
	学习切片-多维，for-range
 */
package main
import "fmt"

func main() {

	slice1 := make([]int, 0, 10)
	// 切片可以反复扩展直到占据整个相关数组
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0:i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}


	// For-range 结构
	for i,v := range slice1{
		fmt.Printf("Slice at %d is %d\n", i, v)
	}

	// 多维切片
	slice2 := [][]int{{10}, {100, 200}}
	// 为第一个切片追加值为 20 的元素
	slice2[0] = append(slice2[0], 20)
	fmt.Println(slice2)
}




