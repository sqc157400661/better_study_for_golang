/*
	如何把一个有序整数数组放到二叉树中
 */
package main

import "fmt"

// 声明二叉树的结构
type BNode struct{
	Data interface{}
	LeftChild *BNode
	RightChild *BNode
}
func NewBNode() *BNode {
	return &BNode{}
}
func main(){
	data :=[]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("数组:",data)
	root := arrayToTree(data,0,len(data)-1)
	fmt.Println("转换成树的先序遍历为：")
	PrintTreePreOrder(root)
	fmt.Println("\n转换成树的中序遍历为：")
	PrintTreeMidOrder(root)
	fmt.Println("\n转换成树的后序遍历为：")
	PrintTreeAfterOrder(root)
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

// 中序遍历
func PrintTreeMidOrder(root *BNode)  {
	if root == nil {
		return
	}
	//遍历root结点的左子树
	if root.LeftChild !=nil {
		PrintTreeMidOrder(root.LeftChild)
	}
	//遍历root结点
	fmt.Print(root.Data, " ")
	//遍历root结点的右子树
	if root.RightChild != nil {
		PrintTreeMidOrder(root.RightChild)
	}
}

// 后序遍历
func PrintTreeAfterOrder(root *BNode)  {
	if root == nil {
		return
	}
	//遍历root结点的左子树
	if root.LeftChild !=nil {
		PrintTreeAfterOrder(root.LeftChild)
	}
	//遍历root结点的右子树
	if root.RightChild != nil {
		PrintTreeAfterOrder(root.RightChild)
	}
	//遍历root结点
	fmt.Print(root.Data, " ")
}


/*
	说明：
	先序遍历：先访问根节点，再访问左子树，最后访问右子树。
	后序遍历：先访问左子树，再访问右子树，最后访问根节点。
	中序遍历：先访问左子树，再访问根节点，最后访问右子树。
	层次遍历：每一层从左到右访问每一个节点。
	参考：https://www.jianshu.com/p/45d75aeb3b01
 */