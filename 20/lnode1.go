/**
实现链表的逆序
方法一：就地逆序
*/
package main

import "fmt"

//链表定义
type LNode struct{
	Data interface{} // 数据域
	Next *LNode //  对下一个节点的作用
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


func main(){
	head := &LNode{}
	CreateNode(head,8)
	PrintNode("逆序前",head)
	Reverse(head)
	PrintNode("逆序后",head)
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
func Reverse(node *LNode){
	if node==nil || node.Next == nil{
		return
	}
	var pre *LNode // 定义前驱节点
	var cur *LNode // 定义当前节点,临时变量
	next := node.Next // 后继结点存起来，防止丢失 从第二个节点开始循环
	for next != nil { //next为空说明遍历到最后一个了，反之说明没有到最后
		cur = next.Next
		next.Next = pre
		pre = next  //后移前驱结点
		next = cur  //后移后驱结点
	}
	node.Next = pre // 最后一个节点变成首节点
}