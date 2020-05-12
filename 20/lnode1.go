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
	for i:=1;i<max;i++{
		node.Next = &LNode{}
		node.Next.Data = i
		node = node.Next
		fmt.Printf("%p \n",node)
	}
	fmt.Println(node.Data)
}

func main(){
 	head := &LNode{}
 	CreateNode(head,8)
	fmt.Printf("%p \n",head)
	fmt.Println(head.Data)
	//PrintNode("逆序前",head)

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
	var cur *LNode // 定义当前节点
	next := node.Next // 后继极点

	for next != nil{
		pre = cur
		cur = node.Next
		next = node.Next.Next
	}
	node.Next = pre
}