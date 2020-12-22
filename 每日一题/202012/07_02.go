package main

import "fmt"

/*
 * 方法三 动态规划—— Kadane算法 O(n)
 *
 * 在整个数组或在固定大小的滑动窗口中找到总和或最大值或最小值的问题，
 * 可通过动态规划(DP)在线性时间内解决
 *
 * 两种标志DP适用于数组：
 * 1. 常数空间，沿数组移动并子啊原数组修改；
 * 2. 线性空间，首先沿left->right方向移动，然后沿right->left方向移动，最后合并结果。
 *
 * 本题可通过修改数组跟踪当前位置的最大和，
 * 在知道当前位置的最大和后更新全局最大和。
 * */
/**
 * max sum of the subarray
 * @param arr int整型一维数组 the array
 * @return int整型
 */
func maxsumofSubarray(arr []int) int {
	var n = len(arr)
	var maxSum = arr[0]
	// 如果当前值小于0，
	// 重新开始(全局最大值更新)
	for i := 1; i < n; i++ {
		// 更新当前的最大值
		if arr[i-1] > 0 {
			arr[i] += arr[i-1]
		}
		// 更新全局的最大值
		if arr[i] > maxSum {
			maxSum = arr[i]
		}
	}

	return maxSum
}

func main() {
	arr := []int{1, -2, 3, 5, -2, 6, -1}
	fmt.Println(maxsumofSubarray(arr))
}
