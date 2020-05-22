/**
	如何把链表以k个结点为一组进行翻转
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

func main() {
	fmt.Println("把链表以k个结点为一组进行翻转")
	head := &LNode{}
	CreateNode(head, 8)
	PrintNode("顺序输出", head)
	ReverseK(head,3)
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

// 对不带头结点的单链表进行翻转
func ReverseWithNoHead(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}
	var pre *LNode
	var next *LNode
	for node != nil{
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return  pre
}

func ReverseK(head *LNode,k int)  {
	if head == nil || head.Next == nil {
		return
	}
	pre := head
	begin := head.Next
	var end  *LNode
	var pNext  *LNode
	for begin !=nil {
		end  = begin
		// 对应图中第（1）步，找到从begin开始的第k个结点
		for i:=1;i<k;i++{
			if end.Next !=nil {
				end = end.Next
			}else{
				return
			}
		}
		pNext = end.Next // 图中第（2）步
		end.Next = nil // 图中第（3）步
		pre.Next = ReverseWithNoHead(begin) // 图中第（4）（5）步
		begin.Next = pNext
		pre = begin // 图中第（7）步
		begin = pNext // 图中第（8）步
	}

}