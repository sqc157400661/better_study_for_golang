/**
	Go的map中删除子map，内存会自动释放吗？
	实验2：map套子map，顶层map是int到子map的映射，子map是int到int的映射，同样先执行delete，再设置为nil，分别看垃圾回收情况。
 */
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// 开启pprof，监听请求
	ip := "0.0.0.0:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}