package main

type Node struct {
	Val  int
	Next *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil{
		return nil
	}
	cur := head
	// 1. 复制各节点，并构建拼接链表
	for cur != nil {
		tmp := &Node{Val:cur.Val}
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = tmp.Next
	}
	// 2. 构建各新节点的 random 指向
	cur = head
	for cur != nil {
		if cur.Random !=nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 3. 拆分两链表
	pre := head
	cur = head.Next
	res := head.Next

	for cur.Next != nil{
		pre.Next = pre.Next.Next
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = nil // 单独处理原链表尾节点
	return res      // 返回新链表头节点
}