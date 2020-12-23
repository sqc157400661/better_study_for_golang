package main

type TreeNode struct {
	val int
	leftChild *TreeNode
	rightChild *TreeNode
}

// 借助队列
type Queue struct{
	arr []*TreeNode
}
func(s *Queue) EnQueue(node *TreeNode){
	s.arr = append(s.arr,node)
}

func(s *Queue) DeQueue()  *TreeNode{
	if len(s.arr) ==0 {
		return nil
	}
	re := s.arr[0]
	if len(s.arr)>1 {
		s.arr =s.arr[1:]
	}else{
		s.arr = nil
	}
	return re
}
func(s *Queue) Size() int{
	return len(s.arr)
}


/**
	层级遍历法
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	queue:=Queue{}
	var dep int
	queue.EnQueue(root)
	for queue.Size()>0 {
		tmp := queue.DeQueue()
		if tmp.leftChild != nil{
			queue.EnQueue(tmp.leftChild)
		}
		if tmp.rightChild !=nil {
			queue.EnQueue(tmp.rightChild)
		}
		dep++
	}
	return dep
}