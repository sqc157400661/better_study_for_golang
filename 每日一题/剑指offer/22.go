package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	node1 := head
	node2 := head
	// node2 先向前移动k步
	for i := 0; i < k; i++ {
		if node2 == nil {
			return nil
		}
		node2 = node2.Next
	}

	for node2 != nil {
		node1 = node1.Next
		node2 = node2.Next
	}
	return node1
}
