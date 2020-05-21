/**
	如何把链表相邻元素翻转
	就地逆序
*/
package main

import "fmt"

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

func main() {
	fmt.Println("相邻元素的翻转")
	head := &LNode{}
	CreateNode(head, 8)
	PrintNode("顺序输出", head)
	Reverse(head)
	PrintNode("逆序输出", head)
}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

// 带头节点的逆序
func Reverse(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	pre := head
	cur := head.Next
	var next *LNode
	for cur != nil && cur.Next !=nil {
		next = cur.Next.Next
		pre.Next = cur.Next
		pre.Next.Next = cur
		cur.Next = next
		pre = cur
		cur = next
	}

}