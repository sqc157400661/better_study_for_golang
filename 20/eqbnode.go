/*
	何判断两棵二叉树是否相等
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
	data1 :=[]int{1,2,3,4,5,6,7,8,9,10}
	data2 :=[]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("数组:",data1,data2)
	root1 := arrayToTree(data1,0,len(data1)-1)
	root2 := arrayToTree(data2,0,len(data2)-1)
	isEq := IsEqual(root1,root2)
	if isEq {
		fmt.Print("这两棵树相等")
	}else{
		fmt.Print("这两棵树不相等")
	}
}
// 有序数组转成树结构
func arrayToTree(arr []int,start int,end int) *BNode{
	var root *BNode
	if end >= start{
		root = &BNode{}
		mid := (start + end +1)/2
		// 树的根节点为数组的中间元素
		root.Data = arr[mid]
		// 递归的用左半部分数据构造root的左子树
		root.LeftChild = arrayToTree(arr,start,mid-1)
		// 递归的用右半部分数据构造root的右子树
		root.RightChild = arrayToTree(arr,mid+1,end)
	}
	return root
}

/*
	判断两棵二叉树是否相等
	参数：root1 与 root2 分别为两棵二叉树根节点
	返回值：true：如果两棵树相等则返回true，否则返回false
 */
func IsEqual(root1 *BNode,root2 *BNode) bool {
	if root1 == nil && root2 == nil{
		return true
	}
	if root1 == nil && root2 != nil{
		return false
	}
	if root1 != nil && root2 == nil{
		return false
	}
	if root1.Data == root2.Data {
		return IsEqual(root1.LeftChild,root2.LeftChild) && IsEqual(root1.RightChild,root2.RightChild)
	}
	return false
}