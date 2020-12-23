package main

import "fmt"

/*
	方法一：递归 [尽量不要用递归,尽量使用最优解:动态规划]
	题目分析，假设f[i]表示在第i个台阶上可能的方法数。逆向思维。如果我从第n个台阶进行下台阶，下一步有2中可能，一种走到第n-1个台阶，一种是走到第n-2个台阶。所以f[n] = f[n-1] + f[n-2].
	那么初始条件了，f[0] = f[1] = 1。
	所以就变成了：f[n] = f[n-1] + f[n-2], 初始值f[0]=1, f[1]=1，目标求f[n]
	看到公式很亲切，代码秒秒钟写完。
	int Fibonacci(int n) {
    	if (n<=1) return 1;
    	return Fibonacci(n-1) + Fibonacci(n-2);
	}
 */
func jumpFloor( number int ) int {
	if number<=1 {
		return 1
	}
	return jumpFloor(number-1) + jumpFloor(number-2)
}

func main(){
	fmt.Println(jumpFloor(5))
}