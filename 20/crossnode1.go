/**
	如何判断两个单链表（无环）是否交叉
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
func CreateNode(node *LNode, max int,start int) (reNode *LNode){
	cur := node // 不用这个变量试试下面打印  这里有个go函数调用与栈的关系
	for i := 1; i <= max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
		if(start >0 && i == start){
			reNode = cur
		}
	}
	return reNode
}

func main() {
	fmt.Println("判断相交链表")
	node1 := &LNode{}
	node2 := &LNode{}
	reNode := CreateNode(node1,8,5)
	cur := CreateNode(node2,5,5)
	cur.Next = reNode
	interNode := IsIntersect(node1,node2)
	PrintNode("node1:",node1)
	PrintNode("node2:",node2)
	if interNode == nil {
		fmt.Println("这两个链表不相交")
	}else{
		fmt.Println("这两个链表相交节点是：",interNode.Data)
	}
	fmt.Println(interNode)
}

func IsIntersect(node1 *LNode,node2 *LNode)*LNode{
	tmp1 :=node1.Next
	tmp2 :=node2.Next
	n1,n2 := 0,0
	// 遍历链表1，找到尾结点，同时记录链表1的长度
	for tmp1.Next !=nil {
		tmp1 = tmp1.Next
		n1++
	}
	// 遍历链表2，找到尾结点，同时记录链表2的长度
	for tmp2.Next !=nil {
		tmp2 = tmp2.Next
		n2++
	}
	// 链表1和链表2是相同的尾结点
	if tmp1 == tmp2{
		// 长链表先走|n1 - n2|步
		if n1 > n2{
			for n1 - n2>0 {
				node1 = node1.Next
				n1--
			}
		}else{
			for n2 - n1>0 {
				node2 = node2.Next
				n2--
			}
		}
		for node1 !=node2 {
			node1 = node1.Next
			node2 = node2.Next
		}
		return  node1
	}
	return nil
}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

