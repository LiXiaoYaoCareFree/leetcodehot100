package main

import (
	"fmt"
)

/*
题目：54. 螺旋矩阵 (Spiral Matrix)
描述：给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
输入示例：
3 3
1 2 3
4 5 6
7 8 9
输出示例：
[1 2 3 6 9 8 7 4 5]
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0, m*n)

	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		// Left to Right
		for j := left; j <= right; j++ {
			ans = append(ans, matrix[top][j])
		}
		top++

		// Top to Bottom
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--

		// Right to Left
		if top <= bottom {
			for j := right; j >= left; j-- {
				ans = append(ans, matrix[bottom][j])
			}
			bottom--
		}

		// Bottom to Top
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}
	return ans
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

	result := spiralOrder(matrix)
	fmt.Println(result)
}
