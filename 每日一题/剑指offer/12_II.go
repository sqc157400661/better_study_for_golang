/*
	7牛真题
 */
package main

import "fmt"

var sum int
func main(){
	board := [][]string{
		{"*",".",".","*"},
		{"*","*",".","."},
		{"*",".","*","*"},
		{".","*","*","."},
	}
	fmt.Println(board)
	rows := len(board)
	cols := len(board[0])
	max := 0
	for i:=0; i < rows; i++  {
		for j := 0; j < cols; j++ {
			sum = 0
			tmp := dfsSum(&board,i,j)
			if tmp > max {
				max = sum
			}
		}
	}
	fmt.Println(max)
	fmt.Println(board)
}

func dfsSum(board *[][]string, i, j int) int {
	rows := len(*board)
	cols := len((*board)[0])
	if i<0 || j<0|| i>= rows || j >= cols || (*board)[i][j] != "*"{
		return 0
	}
	sum++
	(*board)[i][j] = "-"
	fmt.Println(*board)
	// 右边
	dfsSum(board,i,j+1)
	// 左边
	dfsSum(board,i,j-1)
	// 下面
	dfsSum(board,i+1,j)
	(*board)[i][j] = "*"
	return sum
}