# 2020年9月份刷题记录 - 链表
主要是突击链表相关操作

## 2020年9月1号
### 题目：删除链表的倒数第N个节点  [来自力扣|难度中等]

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
```
示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
```
解答:
```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    hh := &ListNode{}
    hh.Next = head
	var node1 = hh
	var node2 = hh
	i:=1
	for i<=n+1{
		node1 = node1.Next
		i++
	}
	for node1 != nil  {
		node1 = node1.Next
		node2 = node2.Next
	}
	
	node2.Next = node2.Next.Next
	return hh.Next
}
```

## 2020年9月2号
### 题目：两数相加 [来自力扣|难度中等]
给出两个**非空**的链表用来表示两个非负的整数。其中，它们各自的位数是按照 **逆序** 的方式存储的，并且它们的每个节点只能存储 **一位** 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
```
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
```
解答:
```
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmp int = 0 // 记录进位的变量
	l3 := &ListNode{} // 加和链表
	sumNode := l3 // 用于滑动计算结果的链表指针
	p1 := l1.Next // 记录链表1的首位
	p2 := l2.Next // 记录链表2的首位
	for p1 != nil || p2 !=nil || tmp !=0 {  // 有进位或者2个链表有一个没有遍历完 执行循环
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
```
## 2020年9月3号
## 2020年9月4号
## 2020年9月5号
## 2020年9月6号
## 2020年9月7号
## 2020年9月8号
## 2020年9月9号
## 2020年9月10号
## 2020年9月11号
## 2020年9月12号
## 2020年9月13号
## 2020年9月14号
## 2020年9月15号
## 2020年9月16号
## 2020年9月17号
## 2020年9月18号
## 2020年9月19号
## 2020年9月20号
## 2020年9月21号
## 2020年9月22号
## 2020年9月23号
## 2020年9月24号
## 2020年9月25号
## 2020年9月26号
## 2020年9月27号
## 2020年9月28号
## 2020年9月29号
## 2020年9月30号