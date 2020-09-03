package main

import "fmt"

//链表定义
type ListNode struct {
	Val int // 数据域
	Next *ListNode      //  对下一个节点的作用
}
//创建测试链表
func CreateTestNode(l1 *ListNode, l2 *ListNode) {
	l1.Next = &ListNode{}
	l1.Next.Val = 5
	//l1.Next.Next = &ListNode{}
	//l1.Next.Next.Val = 4
	//l1.Next.Next.Next = &ListNode{}
	//l1.Next.Next.Next.Val = 3

	l2.Next = &ListNode{}
	l2.Next.Val = 5
	//l2.Next.Next = &ListNode{}
	//l2.Next.Next.Val = 6
	//l2.Next.Next.Next = &ListNode{}
	//l2.Next.Next.Next.Val = 4
}

func main(){
	var l1 = &ListNode{}
	var l2 = &ListNode{}
	CreateTestNode(l1,l2)
	PrintNode("l1: ",l1)
	PrintNode("l2: ",l2)
	l3 :=addTwoNumbers(l1,l2)
	PrintNode("l3: ",l3)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmp int = 0 // 记录进位的变量
	l3 := &ListNode{} // 加和链表
	sumNode := l3 // 用于滑动计算结果的链表指针
	p1 := l1.Next // 记录链表1的首位
	p2 := l2.Next // 记录链表2的首位
	for p1 != nil || p2 !=nil || tmp !=0 {  // 有进位 或者 2个链表有一个没有遍历完 执行循环
		var sum int
		if(p1 != nil){
			sum += p1.Val
			p1 = p1.Next
		}
		if(p2 != nil){
			sum += p2.Val
			p2 = p2.Next
		}
		sum += tmp
		tmp = 0 // 进位使用后归零
		// 判断是否有新的进位
		if sum >= 10 {
			tmp = 1
			sum = sum % 10
		}
		sumNode.Next = &ListNode{}
		sumNode.Next.Val = sum
		sumNode = sumNode.Next
	}
	return l3
}

//打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}