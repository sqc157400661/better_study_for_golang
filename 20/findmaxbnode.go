/*
	如何求一棵二叉树的最大子树和
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

var maxSum int
func main(){
	data :=[]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("数组:",data)
	root := arrayToTree(data,0,len(data)-1)
	maxRoot := &BNode{}
	FindMaxSubTree(root,maxRoot)
	fmt.Println("最大子树和为：",maxSum)

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

// 求最大子树
func FindMaxSubTree(root *BNode,maxRoot *BNode) int {
	var sum int
	if root == nil {
		return 0
	}
	// 求root左子树所有节点的和
	lmax := FindMaxSubTree(root.LeftChild,maxRoot)
	// 求root右子树所有节点的和
	rmax := FindMaxSubTree(root.LeftChild,maxRoot)
	sum = lmax + rmax + root.Data.(int)
	// 以root为根的子树和大于前面求出的最大值
	if sum > maxSum {
		maxSum = sum
		maxRoot.Data = root.Data
	}
	return sum
}