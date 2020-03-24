/**
	学习基于共享变量的并发- sync.Mutex
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享变量
var (
	m  sync.Mutex
	v1 int
)

// 修改共享变量
// 在Lock()和Unlock()之间的代码部分是临界区
func change(i int) {
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
	}
	m.Unlock()
}

// 访问共享变量
// 在Lock()和Unlock()之间的代码部分是是临界区
func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func main() {
	var numGR = 21
	var wg sync.WaitGroup

	fmt.Printf("%d", read())

	// 循环创建numGR个goroutine
	// 每个goroutine都执行change()、read()
	// 每个change()和read()都会持有锁
	for i := 0; i < numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Printf(" -> %d", read())
		}(i)
	}

	wg.Wait()
}

/*
	解释：
	1、在for循环中，会不断激活新的goroutine(共21个)执行匿名函数，在每个匿名函数中都会执行change()和read()，意味着每个goroutine都会申请两次锁、释放两次锁
	2、change()和read()中都申请锁，对于这21个goroutine将要分别执行的42个critical section，Lock()保证了在某一时间点只有其中一个goroutine能访问其中一个critical section。
当释放了一个critical section，其它的Lock()将争夺互斥锁，也就是所谓的竞争现象(race condition)。因为竞争的存在，这42个critical section被访问的顺序是随机的，完全无法保证哪个critical section先被访问
	3、每个goroutine中的read()也都会参与锁竞争，所以并不能保证每次change(i)之后会随之执行到read()，
可能goroutine 1的change()执行完后，会跳转到goroutine 3的change()上，这样一来，goroutine 1的read()就无法读取到goroutine 1所修改的v1值，而是访问到其它goroutine中修改后的值

	总结：
	1、Mutex保证了每个critical section安全，某一时间点只有一个goroutine访问到这部分，但也因此而出现了随机性。
	2、适用于读写不确定，并且只有一个读或者写的场景
	3、如果Lock()后忘记了Unlock()，将会永久阻塞而出现死锁
 */