package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// 221. 最大正方形
// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

// 示例 1：
// 输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
// 输出：4

// 示例 2：
// 输入：matrix = [["0","1"],["1","0"]]
// 输出：1

// 示例 3：
// 输入：matrix = [["0"]]
// 输出：0

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	maxSide := 0
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				}
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Input format:
	// M N
	// row1 (string like "10100")
	// row2
	// ...
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var m, n int
		fmt.Sscanf(line, "%d %d", &m, &n)
		matrix := make([][]byte, m)
		for i := 0; i < m; i++ {
			if scanner.Scan() {
				// Assumes row is given as "1 0 1 0 0" or "10100"
				// Standard ACM usually spaces, but let's handle non-spaced string too
				rowStr := strings.ReplaceAll(scanner.Text(), " ", "")
				// or rowStr might be comma separated if copy-pasted.
				// Let's assume compact string for simplicity or handle spaces
				row := make([]byte, 0, n)
				for _, ch := range rowStr {
					if ch == '0' || ch == '1' {
						row = append(row, byte(ch))
					}
				}
				matrix[i] = row
			}
		}
		fmt.Println(maximalSquare(matrix))
	}
}
