/**
	学习函数 - 函数玩法 - 闭包及应用
 */
package main

import (
	"strings"
	"fmt"
)

func main() {
	// 闭包的记忆效应可用于实现类似于设计模式
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后", f2("winter"))
	fmt.Println("文件名处理后", f2("bird.jpg"))
}



//闭包案例
func makeSuffix(suffix string) func (string) string {
	//如果name没有指定后缀，就加上，否则返回原来的名字
	return func (name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}





