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
	dic := map[*Node]*Node{}
	cur := head
	// 复制各节点，并建立 “原节点 -> 新节点” 的 Map 映射
	for cur != nil {
		dic[cur] = &Node{Val:cur.Val}
		cur = cur.Next
	}
	cur = head
	// 构建新链表的 next 和 random 指向
	for cur != nil {
		dic[cur].Next = dic[cur.Next]
		dic[cur].Random = dic[cur.Random]
		cur = cur.Next
	}
	// 返回新链表的头节点
	return dic[head]
}