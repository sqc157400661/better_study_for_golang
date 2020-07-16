/*
	如何从顶部开始逐层打印二叉树结点数据
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
func NewBNode() *BNode {
	return &BNode{}
}

// 队列定义
type SliceQueue struct {
	Arr []interface{}
	sync.RWMutex
}

//返回队列的大小
func (p *SliceQueue) Size() int {
	return len(p.Arr)
}

//删除队列头元素
func (p *SliceQueue) DeQueue() interface{} {
	p.Lock()
	defer p.Unlock()
	if(len(p.Arr) > 0){
		first := p.Arr[0]
		p.Arr = p.Arr[1:]
		return first
	}else{
		return nil
	}

}

//把新元素加入队列尾
func (p *SliceQueue) EnQueue(item interface{}) {
	p.Lock()
	defer p.Unlock()
	p.Arr = append(p.Arr,item)
}


func main(){
	data :=[]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("数组:",data)
	root := arrayToTree(data,0,len(data)-1)
	fmt.Println("树的层级遍历结果为：")
	PrintTreeLayer(root)
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

// 方法功能：用层级遍历的方式打印出二叉树的节点内容
func PrintTreeLayer(root *BNode)  {
	var p *BNode
	queue := &SliceQueue{Arr: make([]interface{}, 0)} // 初始化队列
	// 树根节点进队列
	queue.EnQueue(root)
	for queue.Size()>0 {
		p = queue.DeQueue().(*BNode)
		// 访问当前节点
		fmt.Print(p.Data," ")
		// 如果节点的左孩子不为空则进入队列
		if p.LeftChild !=nil {
			queue.EnQueue(p.LeftChild)
		}
		// 如果节点的右孩子不为空则进入队列
		if p.RightChild !=nil {
			queue.EnQueue(p.RightChild)
		}
	}
}