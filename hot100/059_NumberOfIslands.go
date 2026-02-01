package main

import "fmt"

/*
题目：岛屿数量 (Number of Islands)
描述：给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
输入示例：
4 5
1 1 1 1 0
1 1 0 1 0
1 1 0 0 0
0 0 0 0 0
输出示例：
1
*/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, r, c int) {
	rows, cols := len(grid), len(grid[0])
	if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

func main() {
	var m, n int
	if _, err := fmt.Scan(&m, &n); err != nil {
		return
	}
	grid := make([][]byte, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			var s string
			fmt.Scan(&s)
			if len(s) > 0 {
				grid[i][j] = s[0]
			}
		}
	}
	fmt.Println(numIslands(grid))
}
