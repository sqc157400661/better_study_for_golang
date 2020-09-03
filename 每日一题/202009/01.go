package main

import "fmt"

//链表定义
type ListNode struct {
	Val interface{} // 数据域
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
	CreateNode(head, 4)
	PrintNode("顺序输出", head)
	removeNthFromEnd1(head,1)
	PrintNode("顺序输出", head)
}

//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	var node1 = head
	var node2 = head
	var pre *ListNode
	m:=0
	i:=1
	for i<n &&  node1.Next != nil{
		node1 = node1.Next
		i++
	}
	for node1.Next != nil  {
		m++
		pre = node2
		node1 = node1.Next
		node2 = node2.Next
	}
	if m == 0 {
		head = node2.Next
	}else{
		pre.Next = node2.Next
	}
	return head
}

/**
优化版
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	hh := &ListNode{}
	hh.Next = head
	var node1 = hh
	var node2 = hh
	i:=1
	for i<=n+1{
		node1 = node1.Next
		i++
	}
	for node1 != nil  {
		node1 = node1.Next
		node2 = node2.Next
	}

	node2.Next = node2.Next.Next
	return hh.Next
}