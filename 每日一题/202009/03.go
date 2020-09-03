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
	CreateNode(head, 4)
	PrintNode("顺序输出", head)
	a := reversePrint(head.Next)
	fmt.Println("反序输出",a)
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	a := []int{}
	node := head
	if node.Next !=nil {
		a = reversePrint(node.Next)
	}
	a = append(a, node.Val)
	return a
}

//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}

