/**
 	如何计算两个单链表所代表的数之和
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
func CreateNodeT() (node1 *LNode, node2 *LNode) {
	node1 = &LNode{}
	node2 = &LNode{}
	cur := node1
	for i := 1; i < 7; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i + 2
		cur = cur.Next
	}
	cur = node2
	for i := 9; i >4; i-- {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
	return
}

/*
	双链表相加
 */

func Add(h1 *LNode,h2 *LNode) *LNode{
	if h1==nil || h1.Next==nil {
		return h2
	}
	if h2==nil || h2.Next==nil {
		return h1
	}
	c := 0 // 记录进位
	sum := 0 // 记录两个节点相加的值
	resultHead := &LNode{}// 相加后的链表表头节点
	p1 := h1.Next // 遍历h1
	p2 := h2.Next // 遍历h2
	p := resultHead // 指向链表resultHead最后一个节点
	for p1 !=nil && p2 !=nil {
		p.Next = &LNode{} // 指向新创建的存储相加和的节点
		sum = p1.Data.(int) + p2.Data.(int) + c
		p.Next.Data = sum %10  //两个节点相加和
		c = sum/10 // 进度
		p = p.Next
		p1 = p1.Next
		p2 = p2.Next
	}

	// 链表h2 比 h1长，接下来只需要考虑h2剩余结点的值
	if p1 == nil {
		for p2 != nil{
			p.Next = &LNode{} // 指向新创建的存储相加和的结点
			sum = p2.Data.(int) + c
			p.Next.Data = sum %10  //两个节点相加和
			c = sum/10 // 进度
			p = p.Next
			p2 = p2.Next
		}
	}

	// 链表h1 比 h2长，接下来只需要考虑h1剩余结点的值
	if p2 == nil {
		for p1 != nil{
			p.Next = &LNode{} // 指向新创建的存储相加和的结点
			sum = p1.Data.(int) + c
			p.Next.Data = sum %10  //两个节点相加和
			c = sum/10 // 进度
			p = p.Next
			p1 = p1.Next
		}
	}

	if c==1 {
		p.Next = &LNode{}
		p.Next.Data = 1
	}
	return resultHead
}

func main() {
	fmt.Println("链表相加")
	node1,node2 := CreateNodeT()
	PrintNode("原链表1=>", node1)
	PrintNode("原链表2=>", node2)
	addResult := Add(node1,node2)
	PrintNode("相加后 =>", addResult)
}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

