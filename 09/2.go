/**
	Go的map中删除子map，内存会自动释放吗？
	实验2：map套子map，顶层map是int到子map的映射，子map是int到int的映射，同样先执行delete，再设置为nil，分别看垃圾回收情况。
 */
package main

import (
	"log"
	"runtime"
)

var intMapMap map[int]map[int]int

var cnt = 1024
var lastTotalFreed uint64 // size of last memory has been freed

func main() {
	// 1
	printMemStats()

	// 2
	initMapMap()
	runtime.GC()
	printMemStats()

	// 3
	fillMapMap()
	runtime.GC()
	printMemStats()

	// 4
	log.Println(len(intMapMap))
	for i := 0; i < cnt; i++ {
		delete(intMapMap, i)
	}
	log.Println(len(intMapMap))
	runtime.GC()
	printMemStats()

	// 5
	intMapMap = nil
	runtime.GC()
	printMemStats()
}

func initMapMap() {
	intMapMap = make(map[int]map[int]int, cnt)
	for i := 0; i < cnt; i++ {
		intMapMap[i] = make(map[int]int, cnt)
	}
}

func fillMapMap() {
	for i := 0; i < cnt; i++ {
		for j := 0; j < cnt; j++ {
			intMapMap[i][j] = j
		}
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v  Just Freed = %v Sys = %v NumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, ((m.TotalAlloc-m.Alloc)-lastTotalFreed)/1024, m.Sys/1024, m.NumGC)

	lastTotalFreed = m.TotalAlloc - m.Alloc
}