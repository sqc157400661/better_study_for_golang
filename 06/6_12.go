/**
	学习通道- 等待组（sync.WaitGroup）
 */
package main
func main() {
	var x []int
	go func() { x = make([]int, 10) }()
	go func() { x = make([]int, 1000000) }()
	x[999999] = 1 // NOTE: undefined behavior; memory corruption possible!
}

