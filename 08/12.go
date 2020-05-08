/**
	实现一个单例
 */
package main

import "sync"

// 实现一个单例

type singleton struct{}

var ins *singleton
var mu sync.Mutex

//懒汉加锁:虽然解决并发的问题，但每次加锁是要付出代价的
func GetIns() *singleton {
	mu.Lock()
	defer mu.Unlock()

	if ins == nil {
		ins = &singleton{}
	}
	return ins
}

//双重锁:避免了每次加锁，提高代码效率
func GetIns1() *singleton {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			ins = &singleton{}
		}
	}
	return ins
}

//sync.Once实现
var once sync.Once

func GetIns2() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
