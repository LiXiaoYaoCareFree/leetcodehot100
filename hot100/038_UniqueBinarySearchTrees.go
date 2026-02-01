package main

import "fmt"

/*
题目：不同的二叉搜索树 (Unique Binary Search Trees)
描述：给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
输入示例：
3
输出示例：
5
*/

func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(numTrees(n))
}
