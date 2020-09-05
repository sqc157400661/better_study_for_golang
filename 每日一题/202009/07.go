package main

import "fmt"

//链表定义
type ListNode struct {
	Val  int       // 数据域
	Next *ListNode //  对下一个节点的作用
}

//创建测试链表
func CreateTestNode(l1 *ListNode, l2 *ListNode) {
	l1.Next = &ListNode{}
	l1.Next.Val = 1
	l1.Next.Next = &ListNode{}
	l1.Next.Next.Val = 2
	l1.Next.Next.Next = &ListNode{}
	l1.Next.Next.Next.Val = 4

	l2.Next = &ListNode{}
	l2.Next.Val = 1
	l2.Next.Next = &ListNode{}
	l2.Next.Next.Val = 3
	l2.Next.Next.Next = &ListNode{}
	l2.Next.Next.Next.Val = 4
}
func main() {
	var l1 = &ListNode{}
	var l2 = &ListNode{}
	CreateTestNode(l1, l2)
	PrintNode("l1: ", l1)
	PrintNode("l2: ", l2)
	PrintNode("合并两个有序链表后: ", mergeTwoLists(l1,l2))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var pre *ListNode
	l3 := &ListNode{}
	pre = l3
	for l1 != nil || l2 != nil {
		if l2 ==nil && l1 != nil{
			pre.Next = l1
			l1 = l1.Next
		}

		if l1 ==nil && l2 != nil{
			pre.Next = l2
			l2 = l2.Next
		}
		if(l1 != nil && l2 != nil){
			if l1.Val <= l2.Val{
				pre.Next = l1
				l1 = l1.Next
			}else{
				pre.Next = l2
				l2 = l2.Next
			}
		}
		pre = pre.Next
	}
	return l3.Next
}

//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}
