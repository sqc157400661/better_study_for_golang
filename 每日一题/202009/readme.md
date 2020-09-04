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
### 题目：从尾到头打印链表 [来自力扣|难度简单]
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：
```
输入：head = [1,3,2]
输出：[2,3,1]

ps限制：
0 <= 链表长度 <= 10000
```
解答:
```
// 递归法 [想一想用栈怎么实现]
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	a := []int{}
	node := head
	if node.Next !=nil {
		a = reversePrint(node.Next)
	}
	a = append(a, node.Val)
	return a
}
```

## 2020年9月4号
### 题目：删除链表的节点 [来自力扣|难度简单]
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。
ps说明：
题目保证链表中节点的值互不相同

示例 1：
```
输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
```
示例 2:
```
输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
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
func deleteNode(head *ListNode, val int) *ListNode {
    pre := &ListNode{}
    newNode := pre
	pre.Next = head
	cur := head
	for cur !=nil {
        fmt.Println(cur.Val)
		if(cur.Val==val){
			if(cur.Next !=nil){
				pre.Next = cur.Next
			}else{
				pre.Next = nil
			}
		}
		cur = cur.Next
		pre = pre.Next
	}
	return  newNode.Next
}
```
## 2020年9月5号
### 题目：移除重复节点 [来自力扣|难度简单]
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

进阶:不得使用临时缓冲区
示例 1：
```
输入：[1, 2, 3, 3, 2, 1]
输出：[1, 2, 3]
```
示例 2:
```
输入：[1, 1, 1, 1, 2]
输出：[1, 2]
```
解答1:  [不得使用临时缓冲区 2层遍历实现]
```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
    node := head
	for node != nil{
		node2 := node
		for node2.Next != nil{
			if(node.Val == node2.Next.Val){
				node2.Next = node2.Next.Next
			}else{
				node2 = node2.Next
			}

		}
		node = node.Next
	}
	return head
}
```
解答2:  [使用临时缓冲区 map]
```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	tmpMap := map[int]bool{}
	//初始化一个虚拟节点
	newNode := &ListNode{}
	pre := newNode
	node := head
	pre.Next = head
	for node !=nil {
		if _,ok := tmpMap[node.Val];ok{
			pre.Next = node.Next
			node = node.Next
		}else{
			tmpMap[node.Val] = true
			node = node.Next
			pre = pre.Next
		}

	}
	return newNode.Next
}
```

## 2020年9月6号
### 题目： 删除中间节点 [来自力扣|难度简单]
实现一种算法，删除单向链表中间的某个节点（即不是第一个或最后一个节点），假定你只能访问该节点。

示例 1：
```
输入：单向链表a->b->c->d->e->f中的节点c
结果：不返回任何数据，但该链表变为a->b->d->e->f
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
func deleteNode(node *ListNode) {
	/*
		将当前指针指向的值修改为下一个节点
		注意:node = node.Next 不行  
		因为golang函数的传递都为值传递,传递进来的指针只是一个指针的值拷贝，
		如在函数内部把a指针的拷贝修改为b指针，并不会修改原来的a指针。 所以无效。
	 */
	*node = *node.Next  
}

```
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