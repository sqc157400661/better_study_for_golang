package main

func isStraight(nums []int) bool {
	repeat := map[int]bool{}
	ma,mi := 0,14
	for _,num :=range nums {
		// 跳过大小王
		if num == 0 {
			continue
		}
		//最大牌
		ma = max(num,ma)
		//最小牌
		mi = min(num,mi)
		//若有重复，提前返回 false
		if _,ok := repeat[num];ok{
			return false
		}
		// 添加牌至 Set
		repeat[num] = true
	}
	//最大牌 - 最小牌 < 5 则可构成顺子
	return ma - mi < 5
}

func max(x,y int) int{
	if x>y {
		return x
	}
	return y
}
func min(x,y int) int{
	if x>y {
		return y
	}
	return x
}