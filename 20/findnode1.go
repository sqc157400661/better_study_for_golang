/**
 如何找出单链表中的倒数第k个元素
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

//创建测试链表
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
	快慢指针查找
 */

func findLastK(head *LNode,k int) *LNode{
	if head==nil || head.Next==nil {
		return head
	}
	slow := head
	fast := head

	for i:=1;i<=k && fast !=nil; i++{
		fast = fast.Next
	}
	for fast !=nil { //防止对空指针的访问操作
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	fmt.Println("寻找倒数K")
	head1 := &LNode{}
	CreateNode(head1, 8)
	PrintNode("原链表 ", head1)
	fmt.Println("寻找倒数K",findLastK(head1,3))

}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

