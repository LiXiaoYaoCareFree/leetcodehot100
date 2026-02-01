package main

import "fmt"

/*
题目：完全平方数 (Perfect Squares)
描述：给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
输入示例：
12
输出示例：
3
*/

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := int(^uint(0) >> 1)
		for j := 1; j*j <= i; j++ {
			if f[i-j*j] < minn {
				minn = f[i-j*j]
			}
		}
		f[i] = minn + 1
	}
	return f[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(numSquares(n))
}
