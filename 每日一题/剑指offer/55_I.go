package main

import "sync"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 队列的实现
type queue struct {
	arr []*TreeNode
	sync.RWMutex
}

// 从队列尾部加入元素
func (q *queue) EnQueue(nums *TreeNode) {
	defer q.Unlock()
	q.Lock()
	q.arr = append(q.arr, nums)
}

// 弹出队列头部元素
func (q *queue) DeQueue() *TreeNode {
	defer q.Unlock()
	q.Lock()
	if len(q.arr) == 0 {
		return nil
	}
	re := q.arr[0]
	if len(q.arr) > 1 {
		q.arr = q.arr[1:]
	} else {
		q.arr = []*TreeNode{}
	}
	return re
}

func (q *queue) isEmpty() bool {
	return len(q.arr) == 0
}
func (q *queue) Size() int {
	return len(q.arr)
}

func maxDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := queue{arr: []*TreeNode{root}}
	for !queue.isEmpty() {

		for i := queue.Size(); i > 0; i-- {
			node := queue.DeQueue()
			if node.Left != nil {
				queue.EnQueue(node.Left)
			}
			if node.Right != nil {
				queue.EnQueue(node.Right)
			}
		}
		res++
	}
	return res
}
