package main

import "fmt"

/*
题目：最小路径和 (Minimum Path Sum)
描述：给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
输入示例：
3 3
1 3 1
1 5 1
4 2 1
输出示例：
7
*/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < columns; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][columns-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var m, n int
	if _, err := fmt.Scan(&m, &n); err != nil {
		return
	}
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&grid[i][j])
		}
	}
	fmt.Println(minPathSum(grid))
}
