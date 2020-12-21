package main

import (
	"fmt"
	"sync"
)

// 队列
type SliceQueue struct {
	Arr []int
}

// 新元素加入
func (s *SliceQueue) EnQueue(key int) {
	newQueue := []int{key}
	s.Arr = append(newQueue, s.Arr...)
}

// 从队尾删除元素
func (s *SliceQueue) PopQueue() int {
	item := s.Arr[len(s.Arr)-1]
	s.Arr = s.Arr[:len(s.Arr)-1]
	return item
}

// 移动节点到头部
func (s *SliceQueue) MoveToHead(key int) {
	var index int
	for k, v := range s.Arr {
		if v == key {
			index = k
			break
		}
	}
	// 注意这里需要判断边际
	if index > 0 {
		tmp := s.Arr[index] // 注意这里一定要用个中间变量,因为s.Arr append 操作中 s.Arr[:index]的值已经不是想要的值
		if index < (len(s.Arr) - 1) {
			s.Arr = append([]int{tmp}, append(s.Arr[:index], s.Arr[index+1:]...)...)
		} else if index == (len(s.Arr) - 1) {
			s.Arr = append([]int{tmp}, s.Arr[:index]...)
		}
	}

}

// hash
type Set struct {
	mu sync.RWMutex
	m  map[int]int
}

// 获取元素
func (s *Set) get(key int) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if v, ok := s.m[key]; ok {
		return v
	}
	return -1
}

// 添加元素 注意这里要避免设置重复的值
func (s *Set) Add(key int, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
	return
}

// 删除元素
func (s *Set) Remove(key int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
}

type LRUCache struct {
	size    int
	queue   *SliceQueue
	hashSet *Set
}

// 判断缓存队列是否已满
func (l *LRUCache) IsFull() bool {
	return l.size == len(l.queue.Arr)
}

func (l *LRUCache) Set(key, val int) {
	// 判断是否已经在队列中
	if l.hashSet.get(key) != -1 {
		l.queue.MoveToHead(key)
		l.hashSet.Add(key, val)
		return
	}
	// 如果队列满了,要删除队尾的元素
	if l.IsFull() {
		tpm := l.queue.PopQueue()
		l.hashSet.Remove(tpm)
	}
	// 新节点添加到队列中
	l.queue.EnQueue(key)
	// 把新缓存的节点同时添加到hash节点中
	l.hashSet.Add(key, val)
}

func (l *LRUCache) Get(key int) int {
	// 判断是否在队列中
	if l.hashSet.get(key) != -1 {
		l.queue.MoveToHead(key)
		return l.hashSet.get(key)
	}
	return -1
}

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
func LRU(operators [][]int, k int) []int {
	lru := LRUCache{size: k, queue: &SliceQueue{Arr: make([]int, 0, k)}, hashSet: &Set{m: map[int]int{}}}
	var re []int
	for _, v := range operators {
		if v[0] == 1 {
			lru.Set(v[1], v[2])
		} else {
			re = append(re, lru.Get(v[1]))
		}

		if len(lru.queue.Arr) != len(lru.hashSet.m) {
			fmt.Println(lru.queue.Arr, lru.hashSet.m)
		}
	}
	return re
}

func main() {
	// 测试用例
	//var param = [][]int{[]int{1, 1, 1}, []int{1, 2, 2}, []int{1, 3, 2}, []int{2, 1}, []int{2, 3}, []int{1, 4, 4}, []int{2, 2}}
	var param = [][]int{{1, -324690837, -72487934}, {1, -723504364, -369145172}, {2, -324690837}, {1, 724101438, -30727452}, {1, 366967562, 290286156}, {2, 366967562}, {1, -21417066, -450706222}, {1, -484359960, -121414361}, {1, -629538923, -759874959}, {1, -461538894, 749719150}, {1, -338664886, -3080586}, {2, 522415046}, {1, 134352387, -391032350}, {1, 283492390, 210901529}, {2, -328994470}, {2, -254674447}, {2, 85161833}, {1, 36899859, -561167545}, {1, -232060336, 889831435}, {1, -848557701, -189598178}, {2, -81051921}, {1, 57162090, -193776405}, {1, 157908494, -941008658}, {2, -21417066}, {2, -485579587}, {2, 449908952}, {1, 12002448, -273547528}, {1, -267371849, 634867878}, {1, 128963067, 102756157}, {2, 36899859}, {2, 145140817}, {2, 335181121}, {1, 89294557, -213039610}, {2, 204879021}, {2, 36899859}, {2, -371773103}, {1, -354568989, -788098655}, {2, 283492390}, {2, 148848181}, {1, 641416456, -646693824}, {1, -970438681, 341383327}, {1, 452656493, -68334619}, {1, 209021588, 229415889}, {1, 580338695, -286744639}, {1, 263639066, -99669610}, {1, -586507345, 255143085}, {2, 89294557}, {2, 209021588}, {1, 888670782, -378802985}, {1, -751520980, 8161338}, {2, 12002448}, {1, -967501590, -351688503}, {2, 514432065}, {2, 80361112}, {1, 220830017, -93263524}, {2, 32131049}, {1, -264084048, -155050456}, {1, -482193643, -232757642}, {2, 134352387}, {1, -303458351, -974058471}, {2, 452656493}, {1, 505909889, -508809796}, {1, 28202319, 577206570}, {2, -991605984}, {2, -266360965}, {2, 553306411}, {1, -91084058, 95483640}, {1, 141869753, -629556641}, {1, -48119420, 783684131}, {1, 610948343, 6884450}, {2, 641416456}, {1, -412360441, 31222570}, {1, 431900251, 231767188}, {2, -588592923}, {1, -220001478, -183147041}, {2, -69417139}, {2, -34241947}, {1, 396626603, -435781001}, {2, 161660462}, {2, -788025808}, {2, -622427224}, {1, -18852505, 647730929}, {2, -978490}, {1, -203099668, -173289604}, {1, -9216750, -116999009}, {2, 241113210}, {2, -586507345}, {1, 439742836, 3892905}, {1, 362692206, 64119965}, {1, -343557119, 25898321}, {1, 351163245, 558921656}, {2, 220830017}, {2, -139705440}, {1, -63307654, 582782467}, {1, 174271875, 647087188}, {1, -366938003, 141242360}, {2, 225133310}, {2, -49488568}, {2, -303458351}, {2, 505909889}, {2, -808960104}, {2, 396626603}, {1, -26838658, -388178020}, {2, -153211277}, {2, -144774404}, {1, -536064025, -36365386}, {2, 734467312}, {2, 220830017}, {1, 365062309, -318717846}, {2, -324832567}, {2, 439742836}, {1, -180868327, -159966095}, {2, -751520980}, {2, 580338695}, {2, -354568989}, {2, -970438681}, {2, 26297015}, {2, 53336527}, {2, 362692206}, {1, 764975865, 98827347}, {2, 580338695}, {2, 359090605}}
	fmt.Println(LRU(param, 5))
}
