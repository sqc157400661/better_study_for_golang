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
	deleteNode(head.Next.Next)
	PrintNode("删除指定的中间节点输出", head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	/*
		将当前指针指向的值修改为下一个节点
		注意:node = node.Next 不行
		因为golang函数的传递都为值传递,传递进来的指针只是一个指针的值拷贝，
		如在函数内部把a指针的拷贝修改为b指针，并不会修改原来的a指针。 所以无效。
	 */
	*node = *node.Next
}

//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}
