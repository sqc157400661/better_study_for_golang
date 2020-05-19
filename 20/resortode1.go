/**
如何对链表进行重新排序
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
	找出链表Head的中间节点，把链表从中间断成2个子链表
	快慢指针法思路：慢指针每一次走一步，快指针每次走两步，快指针走到尾时，慢指针正好走到中间

	快指针到链表尾部时，当链表长度为奇数时，慢指针指向的即是链表中间指针，当链表长度为偶数时，
	慢指针指向的结点和慢指针指向的结点的下一个结点都是链表的中间结点
 */

func findLastK(head *LNode) *LNode{
	if head==nil || head.Next==nil {
		return head
	}
	slow := head
	fast := head
	for fast !=nil && fast.Next != nil { //防止对空指针的访问操作
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

/**
	对不带头结点的单链表翻转
	输入参数：链表的头结点
 */
func reverse(head *LNode) *LNode{
	var pre *LNode
	var cur *LNode
	var tmp *LNode
	cur = head.Next
	head.Next = nil
	for cur != nil{
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return  pre
}

/**
	方法功能： 将两部分链表连接起来
	输入参数： 链表的假头结点
 */
func reorder(head *LNode){
	if head==nil || head.Next==nil {
		return
	}
	cur := head.Next
	middleNode := findMiddleNode(head)
	cur1 := reverse(middleNode)
	var tmp  *LNode
	//合并链表
	for cur.Next !=nil{
		tmp = cur.Next
		cur.Next = cur1
		cur = tmp
		tmp = cur1.Next
		cur1.Next = cur
		cur1 = tmp
	}
	cur.Next =cur1
}


func main() {
	head1 := &LNode{}
	CreateNode(head1, 8)
	PrintNode("重新排序前", head1)
	reorder(head1)
	PrintNode("重新排序后", head1)

}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

