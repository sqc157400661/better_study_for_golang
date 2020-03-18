/**
	学习扩展
 */

package main

import (
	"fmt"
	"runtime"
)

type TwoInts struct {
	a int
	b int
}

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%dKb\n", m.Alloc / 1024)

	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Println(two1)


	for i:=1;i<10000;i++ {
		var two1 = TwoInts{12,13}
		get(two1)
	}
	runtime.GC()//显式的触发 GC
	runtime.ReadMemStats(&m)
	fmt.Printf("%dKb\n", m.Alloc / 1024)
}

func (tn *TwoInts) String() string {
	//return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
	return "121"
}

func get(two1 TwoInts){
	two1.a = 100
}

/*
	一个类型如果被定义了String()方法，那么在调用fmt.Printf()，fmt.Print() 和 fmt.Println()时，会自动使用string方法
	可使用 String() 方法来定制类型的字符串形式的输出

	垃圾收集器（GC）,程序自动用gc用来回收垃圾
	runtime 包访问 GC 进程，如调用 runtime.GC() 函数可以显式的触发 GC
 */


