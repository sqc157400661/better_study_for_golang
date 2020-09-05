package main

import "fmt"

//链表定义
type ListNode struct {
	Val int // 数据域
	Next *ListNode      //  对下一个节点的作用
}

//创建链表
func CreateNode(node *ListNode, max int) {
	cur := node // 不用这个变量试试下面打印  这里有个go函数调用与栈的关系
	for i := 1; i < max; i++ {
		cur.Next = &ListNode{}
		cur.Next.Val = i
		cur = cur.Next
	}
}

func main(){
	head := &ListNode{}
	CreateNode(head, 6)
	PrintNode("顺序输出", head)
	newNode := reverseList(head)
	PrintNode("反转链表输出", newNode)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 和readme中的不一样 是因为这里demo的链表有个空的头部
	if head == nil {
		return head
	}
	pre := head.Next
	cur := head.Next.Next
	pre.Next = nil
	var next *ListNode
	for cur !=nil {
		next =cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	head.Next = pre
	return  head
}


//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}
