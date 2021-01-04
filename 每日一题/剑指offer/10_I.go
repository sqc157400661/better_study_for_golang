package main

import "fmt"
//记忆化递归法
var result = map[int]int{}
func fib(n int) int {
	if n == 0 {
		return 0
	}else if n==1 {
		return 1
	}else{
		//记忆化递归法 重复遇到某数字则直接取用，避免了重复的递归计算
		if _,ok:=result[n];!ok{
			result[n]=  fib(n-1) + fib(n-2)
		}
		return result[n]%1000000007
	}
}

func main(){
	fmt.Println(fib(2))
	fmt.Println(fib(5))
	fmt.Println(fib(45))

	fmt.Println(fibV2(2))
	fmt.Println(fibV2(5))
	fmt.Println(fibV2(45))
}

func fibV2(n int) int {
	dp := make(map[int]int,n)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++{
		dp[i] = dp[i-1] + dp[i-2]
		dp[i] %= 1000000007
	}
	return dp[n]
}

// 优化内存版本
func fibV3(n int) int {
	 dpA := 0
	 dpB := 1
	 var sum int
	for i := 0; i <  n; i++{
		sum = (dpA + dpB)%1000000007
		dpA = dpB
		dpB = sum
	}
	return dpA
}