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
	// 定义一个稍后用于覆盖(重写)的函数
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