/*
	如何找出排序二叉树上任意两个结点的最近共同父结点
 */
package main

import (
	"fmt"
)

// 声明二叉树的结构
type BNode struct{
	Data interface{}
	LeftChild *BNode
	RightChild *BNode
}

func main(){
	data :=[]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("数组:",data)
	root := arrayToTree(data,0,len(data)-1)
	newTree := copyTree(root)
	fmt.Println("原树的先序遍历为：")
	PrintTreePreOrder(root)
	fmt.Println()
	fmt.Println("复制树的先序遍历为：")
	PrintTreePreOrder(newTree)
}

func copyTree(root *BNode) *BNode {
	if root == nil {
		return nil
	}
	// 二叉树的根节点
	dupTree := &BNode{}
	dupTree.Data = root.Data
	// 复制左子树
	dupTree.LeftChild = copyTree(root.LeftChild)
	// 复制右子树
	dupTree.RightChild = copyTree(root.RightChild)
	return dupTree
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
// 先序排序
func PrintTreePreOrder(root *BNode)  {
	if root == nil {
		return
	}
	//遍历root结点
	fmt.Print(root.Data, " ")
	//遍历root结点的左子树
	if root.LeftChild !=nil {
		PrintTreePreOrder(root.LeftChild)
	}
	//遍历root结点的右子树
	if root.RightChild != nil {
		PrintTreePreOrder(root.RightChild)
	}
}

