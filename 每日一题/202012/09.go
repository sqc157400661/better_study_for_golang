package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
我们使用两个指针 node1，node2 分别指向两个链表 headA，headB 的头结点，然后同时分别逐结点遍历，
当 node1 到达链表 headA 的末尾时，重新定位到链表 headB 的头结点；
当 node2 到达链表 headB 的末尾时，重新定位到链表 headA 的头结点。

这样，当它们相遇时，所指向的结点就是第一个公共结点。
 */

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	node1 := pHead1
	node2 := pHead2
	for node1 !=node2 {
		if node1 ==nil{
			node1 = pHead2
		}else{
			node1 = node1.Next
		}

		if node2 ==nil{
			node2 = pHead1
		}else{
			node2 = node2.Next
		}
	}
	return node1
}
