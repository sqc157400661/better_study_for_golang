package main

import "sync"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 队列的实现
type queue struct {
	arr []*TreeNode
	sync.RWMutex
}

// 弹出队列头部元素
func (q *queue) EnQueue(nums *TreeNode) {
	defer q.Unlock()
	q.Lock()
	q.arr = append(q.arr, nums)
}

// 从尾部删除元素
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


func levelOrder(root *TreeNode) []int {
	res := []int{}
	if root ==nil {
		return res
	}
	queue :=queue{arr:[]*TreeNode{root}}
	for !queue.isEmpty(){
		node := queue.DeQueue() // 访问当前节点
		res = append(res,node.Val)
		// 如果左孩子不为空则进入队列
		if node.Left !=nil{
			queue.EnQueue(node.Left)
		}
		// 如果右孩子不为空则进入队列
		if node.Right !=nil{
			queue.EnQueue(node.Right)
		}
	}
	return res
}