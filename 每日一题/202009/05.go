package main

import "fmt"

//链表定义
type ListNode struct {
	Val int // 数据域
	Next *ListNode      //  对下一个节点的作用
}

//创建链表
func CreateNode(node *ListNode) {
	node.Next = &ListNode{}
	node.Next.Val = 5
	node.Next.Next = &ListNode{}
	node.Next.Next.Val = 3
	node.Next.Next.Next = &ListNode{}
	node.Next.Next.Next.Val = 3
}

func main(){
	head := &ListNode{}
	CreateNode(head)
	PrintNode("顺序输出", head)
	newNode := removeDuplicateNodes1(head)
	PrintNode("移除重复节点输出", newNode)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	tmpMap := map[int]bool{}
	//初始化一个虚拟节点
	newNode := &ListNode{}
	pre := newNode
	node := head
	pre.Next = head
	for node !=nil {
		if _,ok := tmpMap[node.Val];ok{
			pre.Next = node.Next
			node = node.Next
		}else{
			tmpMap[node.Val] = true
			node = node.Next
			pre = pre.Next
		}

	}
	return newNode.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes1(head *ListNode) *ListNode {
	node := head
	for node != nil{
		node2 := node
		for node2.Next != nil{
			if(node.Val == node2.Next.Val){
				node2.Next = node2.Next.Next
			}else{
				node2 = node2.Next
			}

		}
		node = node.Next
	}
	return head
}

//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}
