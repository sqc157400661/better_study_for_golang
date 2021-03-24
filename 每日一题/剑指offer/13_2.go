package main

import (
	"fmt"
	"sync"
)

// 队列的实现
type queue struct {
	arr [][]int
	sync.RWMutex
}

// 从队列首部加入元素
func (q *queue) DeQue(i, j, si, sj int) {
	defer q.Unlock()
	q.Lock()
	if len(q.arr) == 0 {
		q.arr = [][]int{{i, j, si, sj}}
	} else {
		q.arr = append(q.arr, []int{i, j, si, sj})
	}
}

// 从尾部删除元素
func (q *queue) Pop() []int {
	defer q.Unlock()
	q.Lock()
	if len(q.arr) == 0 {
		return nil
	}
	re := q.arr[0]
	if len(q.arr) >= 1 {
		q.arr = q.arr[1:len(q.arr)]
	}
	return re
}

func (q *queue) len() int {
	return len(q.arr)
}

var visited map[int]bool

func movingCount(m int, n int, k int) int {
	visited = map[int]bool{}
	queue := queue{arr: [][]int{}}
	queue.DeQue(0, 0, 0, 0)
	for queue.len() > 0 {
		item := queue.Pop()
		i, j, si, sj := item[0], item[1], item[2], item[3]
		key := i*100 + j
		_, ok := visited[key]
		if i >= m || j >= n || k < si+sj || ok {
			continue
		}
		visited[key] = true
		// 下面
		queue.DeQue(i+1, j, sums(i, si), sj)
		// 右边
		queue.DeQue(i, j+1, si, sums(j, sj))
	}
	return len(visited)
}

// 数位和增量
func sums(x, si int) int {
	if (x+1)%10 != 0 {
		return si + 1
	} else {
		return si - 8
	}
}

func main() {
	fmt.Println(movingCount(6, 6, 9))
}
