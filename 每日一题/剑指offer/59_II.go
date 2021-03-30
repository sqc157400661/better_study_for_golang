package main

type queue struct{
	arr []int
}

func (q *queue) isEmpty() bool {
	return len(q.arr) == 0
}

type MaxQueue struct {
	Queue *queue
	Deque *queue
}

// 返回队首元素
func (q *queue) Front() int {
	if len(q.arr) == 0 {
		return -1
	}
	return q.arr[0]
}
// 返回队尾元素
func (q *queue) Back() int {
	if len(q.arr) == 0 {
		return -1
	}
	return q.arr[len(q.arr)-1]
}
// 队首删除元素
func (q *queue) DeQue() int {
	if len(q.arr) == 0 {
		return -1
	}
	re := q.arr[0]
	if len(q.arr) > 1 {
		q.arr = q.arr[1:len(q.arr)]
	} else {
		q.arr = []int{}
	}
	return re
}

// 队尾删除元素
func (q *queue) Pop() int {
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

// 队列加入元素
func (q *queue) EnQue(nums int) {
	q.arr = append(q.arr, nums)
}


func Constructor() MaxQueue {
	return MaxQueue{
		Queue:&queue{[]int{}},
		Deque:&queue{[]int{}},
	}
}


func (this *MaxQueue) Max_value() int {
	if this.Deque.isEmpty() {
		return -1
	}
	return this.Deque.Front()
}


func (this *MaxQueue) Push_back(value int)  {
	this.Queue.EnQue(value)
	for !this.Deque.isEmpty() && this.Deque.Back() < value{ // 注意这里是小于 不是小于等于
		this.Deque.Pop()
	}
	this.Deque.EnQue(value)
}


func (this *MaxQueue) Pop_front() int {
	if this.Queue.isEmpty() {
		return -1
	}
	val := this.Queue.Front()
	if val == this.Deque.Front() {
		this.Deque.DeQue()
	}
	this.Queue.DeQue()
	return val
}




/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
