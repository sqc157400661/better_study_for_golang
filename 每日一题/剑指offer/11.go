package main

func minArray(numbers []int) int {
	var i=0
	var j = len(numbers) - 1
	for i<j {
		var m = (i+j)/2
		if numbers[m] >numbers[j] {
			// 旋转点 xx 一定在 [m + 1, j][m+1,j] 闭区间内
			i = m+1
		}else if numbers[m] < numbers[j] {
			//旋转点 xx 一定在[i, m][i,m]
			j =m
		}else{
			j--
		}
	}
	return numbers[i]
}
