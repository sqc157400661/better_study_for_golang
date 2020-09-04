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
	newNode := deleteNode(head,3)
	PrintNode("删除链表的指定值的节点输出", newNode)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	//初始化一个虚拟节点
	pre := &ListNode{}
	newNode := pre
	//让虚拟节点指向头结点
	pre.Next = head
	cur := head
	for cur !=nil {
		//如果找到要删除的结点，直接把他删除
		if(cur.Val==val){
			if(cur.Next !=nil){
				pre.Next = cur.Next
			}else{
				pre.Next = nil
			}
		}
		cur = cur.Next
		pre = pre.Next
	}
	return  newNode.Next
}

//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}
