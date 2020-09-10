/*
	排序链表
*/
package main

import "fmt"

//链表定义
type ListNode struct {
	Val  int       // 数据域
	Next *ListNode //  对下一个节点的作用
}

//创建链表
func CreateNode(node *ListNode) {
	node.Next = &ListNode{}
	node.Next.Val = 5
	node.Next.Next = &ListNode{}
	node.Next.Next.Val = 3
	node.Next.Next.Next = &ListNode{}
	node.Next.Next.Next.Val = 7
}

func main() {
	head := &ListNode{}
	CreateNode(head)
	PrintNode("顺序输出", head)
	newNode := sortList(head)
	PrintNode("排序后输出", newNode)
}

// 采用自底向上的归并排序
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {

}

func merge(head *ListNode,left *ListNode, right *ListNode) *ListNode {
	newNode := head
	pre := newNode
	for left != nil && right != nil {
		if left.Val < right.Val {
			pre.Next = left
			left = left.Next
		} else {
			pre.Next = right
			right = right.Next
		}
		pre = pre.Next
	}
	if left == nil {
		pre.Next = right
	} else {
		pre.Next = left
	}
	return newNode.Next
}

//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}
