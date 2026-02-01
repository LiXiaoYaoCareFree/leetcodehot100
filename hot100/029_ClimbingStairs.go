package main

import "fmt"

/*
题目：爬楼梯 (Climbing Stairs)
描述：假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
输入示例：
2
输出示例：
2
*/

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(climbStairs(n))
}
