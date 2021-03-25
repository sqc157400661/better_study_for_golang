package main

import "fmt"

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, board, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(i, j int, board [][]byte, word string, k int) bool {
	fmt.Println(i,j,board)
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
		return false
	}
	if k == (len(word) - 1) {
		return true
	}
	board[i][j] = ' ' // 标记成空元素防止重复访问
	// 上方
	top := dfs(i-1, j, board, word, k+1)
	// 下方
	down := dfs(i+1, j, board, word, k+1)
	// 左方
	left := dfs(i, j-1, board, word, k+1)
	// 右方
	right := dfs(i, j+1, board, word, k+1)

	res := top || down || left || right
	board[i][j] = word[k] // 元素还原至初始值，即 word[k]，因为走到自己这里board[i][j]初始值必定等于word[k]
	return res
}

func main() {
	test := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(test, "ABCCED"))
}
