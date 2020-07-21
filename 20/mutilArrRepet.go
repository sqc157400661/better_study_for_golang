/**
	对于一个给定的自然数N，有一个N+M个元素的数组，其中存放了小于等于N的所有自然数，求重复出现的自然数序列{X}。
 */
package main

import (
	"fmt"
	"sync"
)

type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func NewSet() *Set {
	return &Set{
		m: map[interface{}]bool{},
	}
}
func (s *Set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func main() {
	arr := []int{1, 3, 4, 2, 5, 4, 5}
	fmt.Println("重复出现的自然数序列")
	set :=FindDup(arr,5)
	for k,_ := range set.m{
		fmt.Print(k," ")
	}
}

func FindDup(arr []int,n int) *Set {
	s :=NewSet()
	if arr == nil {
		return s
	}
	len := len(arr)
	index := arr[0]
	m := len -n
	for true {
		//遇到重复元素，将该重复元素修改为接近N+M的元素
		if arr[index] <0 {
			m--
			arr[index] = len -m
			s.Add(index)
		}
		// 如果剩余重复元素是0了，说明遍历完成，直接返回
		if m == 0{
			return s
		}
		// 判断后续元素是否重复
		arr[index] *=-1
		index = arr[index] * -1
	}
	return s
}
