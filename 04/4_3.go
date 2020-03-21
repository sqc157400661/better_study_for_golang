/**
	学习切片-切片复制
 */
package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3)//[]int{5, 4, 3}
	copyNum := copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	slice3 := slice1
	fmt.Println(copyNum) // copy 函数将返回成功复制的元素的个数，等于两个 slice 中较小的长度
	//copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置

	// 修改slice1原始数据的第一个元素为 999。
	// 这里会影响slice3 因为切片不会因为等号操作进行元素的复制，其实还是共用一个数组结构，
	// slice2是复制的不会受影响  copy为值复制，更改原切片不会影响新切片，而等号复制相反
	slice1[0] = 999

	fmt.Printf("slice1:%v,slice1地址:%p\n", slice1, slice1)
	fmt.Printf("slice2:%v,slice2地址:%p\n", slice2, slice2)
	fmt.Printf("slice3:%v,slice3地址:%p\n", slice3, slice3)

}
/*
	总结：
	1、copy( destSlice, srcSlice []T) int   srcSlice 为数据来源切片，destSlice 为复制的目标（也就是将 srcSlice 复制到 destSlice）
	2、如果想增加切片的容量[cap]，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来
	3、func copy(dst, src []T) int copy 方法将类型为 T 的切片从源地址 src 拷贝到目标地址 dst，
覆盖 dst 的相关元素，并且返回拷贝的元素个数。源地址和目标地址可能会有重叠。拷贝个数是 src 和 dst 的长度最小值。如果 src 是字符串那么元素类型就是 byte
 */




