package main


func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix)==0{
		return res
	}
	l :=0
	r := len(matrix[0]) -1
	t :=0
	b := len(matrix) -1
	for  {
		// left to right
		for i := l; i <= r; i++ {
			res  = append(res,matrix[t][i])
		}
		t++
		if t >b {
			break
		}
		// top to bottom.
		for i := t; i <= b; i++{
			res  = append(res,matrix[i][r])
		}
		r--
		if l > r {
			break
		}
		// right to left.
		for i:= r; i >= l; i--{
			res  = append(res,matrix[b][i])
		}
		b--
		if t > b{
			break
		}
		// bottom to top.
		for i := b; i >= t; i--{
			res  = append(res,matrix[i][l])
		}
		l++
		if l > r{
			break
		}
	}
	return res
}

func main(){

}