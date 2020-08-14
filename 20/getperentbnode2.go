/*
	如何找出排序二叉树上任意两个结点的最近共同父结点 - 后续遍历法
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
	node1 := root.LeftChild.LeftChild.LeftChild
	node2 := root.LeftChild.RightChild
	result := FindParentNodeReverse(root,node1,node2)
	if result !=nil {
		fmt.Println(node1.Data,"与",node2.Data,"最近的公共父节点为：",result.Data)
	}else{
		fmt.Println("没有公共的父节点")
	}

}



func FindParentNodeReverse(root,node1,node2 *BNode) *BNode{
	if root ==nil || root.Data.(int) == node1.Data.(int) || root.Data.(int) == node2.Data.(int) {
		return root
	}
	lChild := FindParentNodeReverse(root.LeftChild,node1,node2)
	rChild := FindParentNodeReverse(root.RightChild,node1,node2)
	// root的左子树没有结点node1和node2，那么一定在root的右子树上
	if lChild == nil{
		return rChild
	}else if rChild == nil{
		return lChild
	}else{
		// node1 与node2 分表位于左子树与右子树上，root就是他们最近的共同父节点
		return root
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

