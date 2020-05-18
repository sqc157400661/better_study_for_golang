/**
如何从无序链表中移除重复项
*/
package main

import (
	"fmt"
	"math/rand"
)

//链表定义
type LNode struct {
	Data interface{} // 数据域
	Next *LNode      //  对下一个节点的作用
}

//创建链表
func CreateNodeT(node *LNode, max int) {
	cur := node // 不用这个变量试试下面打印  这里有个go函数调用与栈的关系
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = rand.Intn(8)
		cur = cur.Next
	}
}

func main() {
	head := &LNode{}
	CreateNodeT(head, 20)
	PrintNode("删除重复节点前", head)
	RemoveDup(head)
	PrintNode("删除重复节点后", head)
}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

func RemoveDup(node *LNode) {
	var ipre *LNode
	for cur := node.Next; cur != nil; cur = cur.Next {
		ipre = cur
		for innerCur := cur.Next; innerCur != nil; innerCur = innerCur.Next {
			if (cur.Data == innerCur.Data) {
				ipre.Next = innerCur.Next
				if(innerCur.Next !=nil){
					innerCur = innerCur.Next
				}
			}else{
				ipre = innerCur
			}
		}
	}
}
