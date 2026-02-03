package main

import (
	"fmt"
)

/*
题目：73. 矩阵置零 (Set Matrix Zeroes)
描述：给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
输入示例：
3 3
1 1 1
1 0 1
1 1 1
输出示例：
1 0 1
0 0 0
1 0 1
*/

func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])

	row0HasZero := false
	col0HasZero := false

	// Check if first row has zero
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			row0HasZero = true
			break
		}
	}

	// Check if first col has zero
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0HasZero = true
			break
		}
	}

	// Use first row and col as markers
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Set zeroes based on markers
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Handle first row and col
	if row0HasZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if col0HasZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	setZeroes(matrix)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j])
			if j < n-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
