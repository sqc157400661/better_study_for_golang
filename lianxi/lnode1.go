/**
实现链表的逆序
原地反转
*/
package main

import "fmt"

type Lnode struct{
	Data interface{}
	Next *Lnode
}

func createNode(head *Lnode,len int) {
	cur := head
	for i:=0;i<len;i++{
		cur.Next = &Lnode{}
		cur.Next.Data = i
		cur = cur.Next;
	}
}


func PrintNode(node *Lnode){
	for cur := node;cur!=nil;cur = cur.Next{
		fmt.Println(cur)
	}
}

func main(){
	node := &Lnode{}
	createNode(node,5)
	PrintNode(node)
	rev(node)
	PrintNode(node)
}

func rev(node *Lnode)  {
	head := node
	cur := node.Next.Next
	pre := node.Next
	pre.Next = nil
	var next *Lnode

	for cur.Next!=nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next

	}
	cur.Next = pre
	head.Next = cur
}