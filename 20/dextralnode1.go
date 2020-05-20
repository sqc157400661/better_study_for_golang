/**
 如何将单链表向右旋转k个位置
*/
package main

import (
	"fmt"
)

//链表定义
type LNode struct {
	Data interface{} // 数据域
	Next *LNode      //  对下一个节点的作用
}

//创建链表
func CreateNode(node *LNode, max int) {
	cur := node // 不用这个变量试试下面打印  这里有个go函数调用与栈的关系
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

/*
	把链表右旋K个位置
*/
func RotateK(head *LNode, k int) *LNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head

	for i := 1; i <= k && fast != nil; i++ {
		fast = fast.Next
	}
	for fast.Next != nil { //防止对空指针的访问操作
		fast = fast.Next
		slow = slow.Next
	}
	tmp := head.Next
	head.Next = slow.Next
	slow.Next = nil
	fast.Next = tmp
	return head
}



func main() {
	fmt.Println("右旋k个位置")
	head1 := &LNode{}
	CreateNode(head1, 8)
	PrintNode("原链表 ", head1)
	PrintNode("右旋3个位置", RotateK(head1, 3))
}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}
