/**
 	何检测一个较大的单链表是否有环
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
func CreateNode(node *LNode, max int) {
	cur := node // 不用这个变量试试下面打印  这里有个go函数调用与栈的关系
	var tmp *LNode
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
		if(max - i ==4){
			tmp = cur
		}
		if(max - i ==1){
			cur.Next = tmp
		}
	}
}

/*
	判断单链表是否有环
 */

func IsLoop(head *LNode) *LNode{
	if head==nil || head.Next==nil {
		return head
	}
	slow := head.Next
	fast := head.Next

	for fast !=nil && fast.Next != nil{ //防止对空指针的访问操作
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast{
			return slow
		}
	}
	return nil
}

// 找出环的入口点
func FindLoopNode(head *LNode,meetNode *LNode) *LNode{
	first := head.Next
	second := meetNode
	for first != second {
		first = first.Next
		second = second.Next
	}
	return first
}

func main() {
	fmt.Println("单链表是否有环")
	head1 := &LNode{}
	CreateNode(head1, 8)
	meetNode := IsLoop(head1)
	fmt.Println(meetNode)
	if meetNode != nil {
		fmt.Println("有环")
		loopNode := FindLoopNode(head1,meetNode)
		fmt.Println("环的入口点为：",loopNode.Data)
	}else{
		fmt.Println("无环")
	}
}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

