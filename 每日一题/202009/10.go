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
	var mid *ListNode
	mid = &ListNode{Val:1}
	node.Next = mid
	node.Next.Next = &ListNode{Val:2}
	//node.Next.Next.Next = mid

	//var mid *ListNode
	//node.Next = &ListNode{}
	//node.Next.Val = 5
	//mid = &ListNode{Val:88}
	//node.Next.Next = mid
	//node.Next.Next.Next = &ListNode{}
	//node.Next.Next.Next.Val = 7
	//node.Next.Next.Next.Next = mid

}

func main() {
	head := &ListNode{}
	CreateNode(head)
	newNode := detectCycle(head.Next)
	fmt.Println(newNode)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head==nil || head.Next==nil{
		return nil
	}
	// 找出相遇的节点
	fast :=head
	slow :=head
	var meeta *ListNode // 相遇的节点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			meeta = fast
			break
		}
	}
	if meeta == nil {
		return meeta
	}
	for meeta != nil {
		if meeta == head {
			break
		}
		meeta = meeta.Next
		head = head.Next
	}
	return meeta
}

//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}
