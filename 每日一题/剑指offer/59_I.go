package main

import (
	"fmt"
	"reflect"
	"sync"
)

// 队列的实现
type queue struct {
	arr []int
	sync.RWMutex
}

// 从队列尾部加入元素
func (q *queue) Append(nums int) {
	defer q.Unlock()
	q.Lock()
	q.arr = append(q.arr, nums)
}

// 从队列首部加入元素
func (q *queue) DeQue(nums int) {
	defer q.Unlock()
	q.Lock()
	if len(q.arr) == 0 {
		q.arr[0] = nums
	} else {
		q.arr = append([]int{nums}, q.arr...)
	}
}

// 从尾部删除元素
func (q *queue) Pop() int {
	defer q.Unlock()
	q.Lock()
	if len(q.arr) == 0 {
		return -1
	}
	re := q.arr[len(q.arr)-1]
	if len(q.arr) > 1 {
		q.arr = q.arr[:len(q.arr)-1]
	} else {
		q.arr = []int{}
	}
	return re
}

// 从头部删除元素
func (q *queue) PopLeft() int {
	defer q.Unlock()
	q.Lock()
	if len(q.arr) == 0 {
		return -1
	}
	re := q.arr[0]
	if len(q.arr) > 1 {
		q.arr = q.arr[1:]
	} else {
		q.arr = []int{}
	}
	return re
}

func (q *queue) isEmpty() bool {
	return len(q.arr) == 0
}

// 查看尾部的数据
func (q *queue) peekLast() int {
	defer q.RUnlock()
	q.RLock()
	if !q.isEmpty() {
		return q.arr[len(q.arr)-1]
	}
	return 0
}

// 查看头部元素
func (q *queue) peekFirst() int {
	defer q.RUnlock()
	q.RLock()
	if !q.isEmpty() {
		return q.arr[0]
	}
	return 0
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums)==0 || k==0 {
		return []int{}
	}
	var res = []int{}
	deque := &queue{arr:[]int{}}
	// 未形成窗口
	for i := 0; i < k; i++ {
		// 循环删除 dequedeque 内所有<nums[j] 的元素，以保持deque递减
		for !deque.isEmpty() && deque.peekLast() < nums[i] {
			deque.Pop()
		}
		deque.Append(nums[i])
	}
	res = append(res,deque.peekFirst())
	// 形成窗口后
	for i := k; i < len(nums); i++ {
		// 队首元素 deque[0] == 被删除元素nums[i−1] ：则队首元素出队；
		if deque.peekFirst() == nums[i - k] {
			deque.PopLeft()
		}
		for !deque.isEmpty() && deque.peekLast() < nums[i] {
			deque.Pop()
		}
		deque.Append(nums[i])
		res = append(res,deque.peekFirst())
	}
	return res
}

func main() {
	demo := []int{1, 3, -1, -3, 5, 3, 6, 7}
	actual := maxSlidingWindow(demo, 3)
	expected := []int{3, 3, 5, 5, 6, 7}
	if !reflect.DeepEqual(actual, []int{3, 3, 5, 5, 6, 7}) {
		fmt.Printf("Expected %v (type %v) - Got %v (type %v)",
			expected,
			reflect.TypeOf(expected),
			actual,
			reflect.TypeOf(actual),
		)
	}
}
