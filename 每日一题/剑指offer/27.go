package main

import "sync"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var tmp = root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(tmp)
	return root
}

// 辅助栈（或队列）
type stack struct{
	arr []*TreeNode
}
func(s *stack) Push(node *TreeNode){
	s.arr = append(s.arr,node)
}

func(s *stack) Pop() *TreeNode{
	if len(s.arr) ==0 {
		return nil
	}
	re := s.arr[len(s.arr)-1]
	if len(s.arr)>1 {
		s.arr =s.arr[0:len(s.arr)-1]
	}else{
		s.arr = nil
	}
	return re
}

func(s *stack) Size() int{
	return len(s.arr)
}

func mirrorTreeV2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := &stack{}
	stack.Push(root)
	for stack.Size() >0{
		node := stack.Pop()
		if node.Left !=nil {
			stack.Push(node.Left)
		}
		if node.Right !=nil {
			stack.Push(node.Right)
		}
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp
	}
	return root
}
