package main

import "fmt"

/*
 * 方法一 贪心法 O(n)
 *
 * 当叠加的和小于0时，就从下一个数重新开始，
 * 同时更新最大和的值(最大值可能为其中某个值)，
 * 当叠加和大于0时，将下一个数值加入和中，
 * 同时更新最大和的值，依此继续。
 *
 * 举例： nums = [-2,1,-3,4,-1,2,1,-5,4]
 * sum = INT_MIN <= 0-> sum = -2 <= 0 -> sum = 1 > 0 ->
 * -> sum = -2 <= 0 -> sum = 4 > 0 -> sum = 3 > 0 ->
 * -> sum = 5 > 0 -> sum = 6 > 0 -> sum = 1 > 0 ->
 * -> sum = 5 > 0
 * res = [-2, 1, 1, 4, 4, 5, 6, 6, 6]
 * 最终返回 res = 6
 * */


/**
 * max sum of the subarray
 * @param arr int整型一维数组 the array
 * @return int整型
 */
func maxsumofSubarray( arr []int ) int {
	var resSum,curSum int
	for i:=0;i<len(arr);i++{
		// 当sum小于0时，就从下一个数重新开始
		// 同时更新每次叠加的最大值
		if curSum <= 0 {
			curSum = arr[i]
		} else {
			// 和大于0时
			curSum += arr[i]
		}

		// 不断更新子串的最大值
		if (curSum > resSum) {
			resSum = curSum
		}
	}
	return resSum
}

func main() {
	arr := []int{1, -2, 3, 5, -2, 6, -1}
	fmt.Println(maxsumofSubarray(arr))
}