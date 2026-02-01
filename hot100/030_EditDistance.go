package main

import "fmt"

/*
题目：编辑距离 (Edit Distance)
描述：给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符
输入示例：
horse
ros
输出示例：
3
*/

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	if n*m == 0 {
		return n + m
	}
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < m+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			left_down := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				left_down += 1
			}
			dp[i][j] = min(left, min(down, left_down))
		}
	}
	return dp[n][m]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var word1, word2 string
	fmt.Scan(&word1)
	fmt.Scan(&word2)
	fmt.Println(minDistance(word1, word2))
}
