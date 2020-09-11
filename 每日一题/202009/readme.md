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
### 题目： 合并两个有序链表 [来自力扣|难度简单]
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
```
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var pre *ListNode
	l3 := &ListNode{}
	pre = l3
	for l1 != nil || l2 != nil {
		if l2 ==nil && l1 != nil{
			pre.Next = l1
			l1 = l1.Next
		}

		if l1 ==nil && l2 != nil{
			pre.Next = l2
			l2 = l2.Next
		}
		if(l1 != nil && l2 != nil){
			if l1.Val <= l2.Val{
				pre.Next = l1
				l1 = l1.Next
			}else{
				pre.Next = l2
				l2 = l2.Next
			}
		}
		pre = pre.Next
	}
	return l3.Next
}

```

## 2020年9月8号
### 题目： 反转链表。 [来自力扣|难度简单]
反转一个单链表。

示例 1：
```
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
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
func reverseList(head *ListNode) *ListNode {
    if head == nil {
		return head
	}
    pre := head
	cur := head.Next
	pre.Next = nil
	var next *ListNode
	for cur !=nil {
		next =cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	
	return  pre
}
```


## 2020年9月9号
### 题目： 排序链表。 [来自力扣|难度简单]
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1：
```
输入: 4->2->1->3
输出: 1->2->3->4
```
示例 2：
```
输入: -1->5->3->4->0
输出: -1->0->3->4->5
```

解答1:

采用自顶向下的递归的归并排序 参考:http://blog.xiaot123.com/5-x99nf

```
// 本次使用自顶向下的递归归并算法
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	var mid *ListNode
	if head == nil||head.Next == nil {
		return head
	}
	// 1. 找中点，二分，左右分别进行排序,快慢指针 fast比slow快一个是为了让slow.next可以成为中点，便于截断
	fast := head.Next
	slow := head
	for fast !=nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid = slow.Next
	// 注意要将链表截断
	slow.Next = nil
	// 对右边进行分割
	l := sortList(head)
	// 对左边进行分割
	r :=sortList(mid)
	// 合并 将2个链表合并成一个
	tmpHead := &ListNode{}
	return merge(tmpHead, l,r)
}

func merge(head *ListNode,left *ListNode, right *ListNode) *ListNode {
	newNode := head
	pre := newNode
	for left != nil && right != nil {
		if left.Val < right.Val {
			pre.Next = left
			left = left.Next
		} else {
			pre.Next = right
			right = right.Next
		}
		pre = pre.Next
	}
	if left == nil {
		pre.Next = right
	} else {
		pre.Next = left
	}
	return newNode.Next
}
```


解答2(待整理):

采用自底向上的归并排序 参考:http://blog.xiaot123.com/5-x99nf
```
// 采用自底向上的归并排序
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {

}
```

## 2020年9月10号
### 题目： 环路检测。 [来自力扣|难度中等]
给定一个链表，如果它是有环链表，实现一个算法返回环路的开头节点。
有环链表的定义：在链表中某个节点的next元素指向在它前面出现过的节点，则表明该链表存在环路。

示例 1：
```
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。
```
示例 2：
```
输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。
```
示例 3：
```
输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。
```
进阶：
你是否可以不用额外空间解决此题？


解答:

```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
   if head==nil || head.Next==nil{
		return nil
	}
	// 找出相遇的节点
	fast :=head
	slow :=head
	var meeta *ListNode // 相遇的节点
	for fast !=nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			meeta = fast
			break
		}
	}
	if meeta == nil {
		return meeta
	}
	for meeta != nil {
		if meeta == head {
			break
		}
		meeta = meeta.Next
		head = head.Next
	}
	return meeta
}
```

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