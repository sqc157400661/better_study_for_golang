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
	l1.Next.Next.Next.Val = 3
	l1.Next.Next.Next.Next = &ListNode{}
	l1.Next.Next.Next.Next.Val = 3
	l1.Next.Next.Next.Next.Next = &ListNode{}
	l1.Next.Next.Next.Next.Next.Val = 4
	l1.Next.Next.Next.Next.Next.Next = &ListNode{}
	l1.Next.Next.Next.Next.Next.Next.Val = 4
	l1.Next.Next.Next.Next.Next.Next.Next = &ListNode{}
	l1.Next.Next.Next.Next.Next.Next.Next.Val = 5

	l2.Next = &ListNode{}
	l2.Next.Val = 1
	l2.Next.Next = &ListNode{}
	l2.Next.Next.Val = 1
	l2.Next.Next.Next = &ListNode{}
	l2.Next.Next.Next.Val = 1
	l2.Next.Next.Next.Next = &ListNode{}
	l2.Next.Next.Next.Next.Val = 2
	l2.Next.Next.Next.Next.Next = &ListNode{}
	l2.Next.Next.Next.Next.Next.Val = 3
}
func main() {
	var l1 = &ListNode{}
	var l2 = &ListNode{}
	CreateTestNode(l1, l2)
	PrintNode("l1: ", l1)
	PrintNode("l2: ", l2)
	PrintNode("去除重复l1: ", deleteDuplicates(l1))
	//PrintNode("去除重复l2: ",  deleteDuplicates(l2))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	node := &ListNode{Next:head}
	node1 := node
	var eq int
	for node.Next != nil{
		node2 := node.Next
		for node2.Next != nil{
			if(node.Next.Val == node2.Next.Val || eq ==node2.Next.Val){
				fmt.Println(node.Next.Val)
				eq = node2.Next.Val
				node2 = node2.Next.Next
			}else{
				node2 = node2.Next
			}
		}
		node = node.Next
	}
	return node1.Next
}
//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}
