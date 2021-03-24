package main

import "fmt"

func movingCount(m int, n int, k int) int {
	visited = map[int]bool{}
	return dfs(0, 0, 0, 0, m, n, k)
}

var visited map[int]bool

func dfs(i, j, si, sj, m, n, k int) int {
	key := i*100 + j
	_, ok := visited[key]
	if i >= m || j >= n || k < si+sj || ok {
		return 0
	}
	visited[key] = true
	// 下方
	down := dfs(i+1, j, sums(i, si), sj, m, n, k)
	// 右方
	right := dfs(i, j+1, si, sums(j, sj), m, n, k)

	return 1 + down + right
}

func main() {
	fmt.Println(movingCount(6, 6, 9))
}

// 数位和增量
func sums(x, si int) int {
	if (x+1)%10 != 0 {
		return si + 1
	} else {
		return si - 8
	}
}

func fullSums(x int) int {
	s := 0
	for x != 0 {
		s += x % 10
		x = x / 10
	}
	return s
}
