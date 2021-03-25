package main

func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += max(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
