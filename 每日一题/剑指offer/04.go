package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i := len(matrix) - 1
	j := 0
	for i >= 0 && j < len(matrix[i]) {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else if matrix[i][j] == target {
			return true
		}
	}
	return false
}
