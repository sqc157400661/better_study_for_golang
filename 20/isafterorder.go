/*
	如何判断一个数组是否是二元查找树后序遍历的序列
 */
package main

import "fmt"

// 声明二叉树的结构
type BNode struct{
	Data interface{}
	LeftChild *BNode
	RightChild *BNode
}


func main(){
	data :=[]int{1,3,2,5,7,6,4}
	fmt.Println("数组:",data)
	result := IsAfterOrder(data,0,len(data)-1)
	if result {
		fmt.Println("是二元查找树的后续遍历序列")
	}
}


/*
	方法功能：判断一个数组是否是二元查找树的后续遍历序列
	输入参数：arr：数组
 */
func IsAfterOrder(arr []int,start int,end int) bool{
	if arr == nil {
		return false
	}
	// 数组的最后一个结点必定是跟节点
	root := arr[end]
	var i,j int
	// 找到一个大于root的值，那么前面所有的结点都位于root的左子树上
	for i = start;i<end;i++{
		if arr[i] > root {
			break
		}
	}
	// 如果序列是后续遍历的序列，那么i开始的所有值都应该大于根结点root的值
	for j=i;j<end;j++ {
		if arr[j] < root {
			return false
		}
	}
	leftIsAfterOrder := true
	rightIsAfterOrder := true
	// 判断小于root值序列是否是某一二元查找树的后续遍历
	if i>start{
		leftIsAfterOrder = IsAfterOrder(arr,start,i-1)
	}
	// 判断大于root值的序列是否是某二元查找树的后续遍历
	if j<end{
		rightIsAfterOrder = IsAfterOrder(arr,i,end)
	}
	return leftIsAfterOrder && rightIsAfterOrder
}

