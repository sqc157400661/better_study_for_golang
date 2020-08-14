/*
	如何找出排序二叉树上任意两个结点的最近共同父结点
 */
package main

import (
	"fmt"
	"sync"
)

// 声明二叉树的结构
type BNode struct{
	Data interface{}
	LeftChild *BNode
	RightChild *BNode
}

//slice所写的stack
type SliceStack struct {
	arr       []interface{}
	stackSize int
	sync.RWMutex
}

//Push栈元素
func (p *SliceStack) Push(t interface{}) {
	p.Lock()
	defer p.Unlock()
	p.arr = append(p.arr, t)
	p.stackSize = p.stackSize + 1
}

//弹出栈元素
func (p *SliceStack) Pop() interface{} {
	p.Lock()
	defer p.Unlock()
	if p.stackSize > 0 {
		p.stackSize--
		ret := p.arr[p.stackSize]
		p.arr = p.arr[:p.stackSize]
		return ret
	}
	return nil
}

func main(){
	data :=[]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("数组:",data)
	root := arrayToTree(data,0,len(data)-1)
	node1 := root.LeftChild.LeftChild.LeftChild
	node2 := root.LeftChild.RightChild
	result := FindParentNode(root,node1,node2)
	if result !=nil {
		fmt.Println(node1.Data,"与",node2.Data,"最近的公共父节点为：",result.Data)
	}else{
		fmt.Println("没有公共的父节点")
	}

}

/**
	方法功能：获取二叉树从根节点root到node结点的路径
	输入参数：root：跟节点；node二叉树中某个节点，s：用来存储路径的栈
 */
func GetPathFromRoot(root *BNode,node *BNode,s *SliceStack) bool {
	if root == nil {
		return false
	}
	if root.Data.(int) == node.Data.(int) {
		s.Push(root)
		return true
	}
	// 如果node结点在root结点的左子树或者右子树上，那么root就是node的祖先结点，把它加到栈里
	if GetPathFromRoot(root.LeftChild,node,s) || GetPathFromRoot(root.RightChild,node,s){
		s.Push(root)
		return true
	}
	return false
}

func FindParentNode(root,node1,node2 *BNode) *BNode{
	stack1 := &SliceStack{arr: make([]interface{}, 0)} // 保存从root到node1的路径
	stack2 := &SliceStack{arr: make([]interface{}, 0)} // 保存从root到node2的路径
	// 获取从root到node1的路径
	GetPathFromRoot(root,node1,stack1)
	// 获取从root到node2的路径
	GetPathFromRoot(root,node2,stack2)
	var commonParent *BNode
	for t1,t2 := stack1.Pop().(*BNode),stack2.Pop().(*BNode);t1 !=nil && t2 !=nil && t1.Data.(int) == t2.Data.(int);{
		commonParent = t1
		t1 = stack1.Pop().(*BNode)
		t2 = stack2.Pop().(*BNode)
	}
	return commonParent
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

