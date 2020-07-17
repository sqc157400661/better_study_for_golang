/*
	如何把二叉树转换为双向链表
 */
package main

import "fmt"

// 声明二叉树的结构
type BNode struct{
	Data interface{}
	LeftChild *BNode
	RightChild *BNode
}

var pHead *BNode
var pEnd *BNode

func main(){
	data :=[]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("数组:",data)
	root := arrayToTree(data,0,len(data)-1)
	InOrderBSTree(root)
	fmt.Print("转换后双向链表正向遍历：")
	for cur:=pHead;cur !=nil;cur=cur.RightChild{
		fmt.Print(cur.Data," ")
	}
	fmt.Println()
	fmt.Print("转换后双向链表逆向遍历：")
	for cur:=pEnd;cur !=nil;cur=cur.LeftChild{
		fmt.Print(cur.Data," ")
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
	方法功能：把二叉树转换成双向链表
	输入参数：root：二叉树跟节点
 */
func InOrderBSTree(root *BNode){
	if root == nil {
		return
	}
	// 转换root的左子树
	InOrderBSTree(root.LeftChild)
	// 使用当前节点的左孩子指向双向链表中农的最后一个结点
	root.LeftChild = pEnd
	// 双向列表为空，当前遍历的结点为双向链表的头结点
	if pEnd == nil {
		pHead = root
	}else{
		// 使双向链表中最后一个结点的右孩子指向当前结点
		pEnd.RightChild = root
	}
	// 将当前结点设为双向链表中最后一个结点
	pEnd = root
	// 转换root的右字数
	InOrderBSTree(root.RightChild)
}

