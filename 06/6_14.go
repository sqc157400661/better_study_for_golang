/**
	学习基于共享变量的并发- sync.RWMutex
 */
package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var Password = secret{password: "myPassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

// 通过rwmutex写  Change()函数申请写锁，并睡眠3秒后修改数据，然后释放写锁。
func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("Change with rwmutex lock")
	time.Sleep(3 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

// 通过rwmutex读  rwMutexShow()函数申请读锁，并睡眠一秒后取得数据，并释放读锁。注意，rwMutexShow()中的print和return是相隔一秒钟的。
func rwMutexShow(c *secret) string {
	c.RWM.RLock()
	fmt.Println("show with rwmutex",time.Now().Second())
	time.Sleep(1 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

// 通过mutex读，和rwMutexShow的唯一区别在于锁的方式不同
// mutexShow()函数申请Mutex锁，和RWMutex互不相干。和rwMutexShow()唯一不同之处在于申请的锁不同
func mutexShow(c *secret) string {
	c.M.Lock()
	fmt.Println("show with mutex:",time.Now().Second())
	time.Sleep(1 * time.Second)
	defer c.M.Unlock()
	return c.password
}

func main() {
	// 定义一个稍后用于覆盖(重写)的函数  函数类型和rwMutexShow()、mutexShow()相同可以相互赋值
	var show = func(c *secret) string { return "" }

	// 通过变量赋值的方式，选择并重写showFunc函数
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex!",time.Now().Second())
		show = rwMutexShow
	} else {
		fmt.Println("Using sync.Mutex!",time.Now().Second())
		show = mutexShow
	}

	var wg sync.WaitGroup

	// 激活5个goroutine，每个goroutine都查看
	// 根据选择的函数不同，showFunc()加锁的方式不同
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Go Pass:", show(&Password),time.Now().Second())
		}()
	}

	// 激活一个申请写锁的goroutine
	go func() {
		wg.Add(1)
		defer wg.Done()
		Change(&Password, "123456")
	}()
	// 阻塞，直到所有wg.Done
	wg.Wait()
}

/*
	解释：
	1、for循环中激活了5个goroutine并发运行，for瞬间激活5个goroutine后，继续执行main()代码会激活另一个用于申请写锁的goroutine。这6个goroutine的执行顺序是随机的
	2、如果show选中的函数是rwMutexShow()，则5个goroutine要申请的RLock()锁和写锁是冲突的，但5个RLock()是兼容的。
所以，只要某个时间点调度到了写锁的goroutine，剩下的读锁goroutine都会从那时开始阻塞3秒。
	3、5个goroutine都申请读锁，因为申请后立即print输出，然后睡眠1秒，
但1秒时间足够所有剩下的goroutine申请完读锁，使得show with rwmutex输出是连在一起，输出的Go Pass: 123456又是连在一起的
	4、Mutex和RWMutex都不关联goroutine，但RWMutex显然更适用于读多写少的场景。仅针对读的性能来说，RWMutex要高于Mutex，因为rwmutex的多个读可以并存。
 */