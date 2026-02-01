package main

import "fmt"

/*
题目：最大矩形 (Maximal Rectangle)
描述：给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
输入示例：
4 5
1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0
输出示例：
6
*/

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i := 0; i < m; i++ {
		left[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if j == 0 {
					left[i][j] = 1
				} else {
					left[i][j] = left[i][j-1] + 1
				}
			}
		}
	}
	ret := 0
	for j := 0; j < n; j++ {
		up := make([]int, m)
		down := make([]int, m)
		stk := []int{}
		for i := 0; i < m; i++ {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= left[i][j] {
				stk = stk[:len(stk)-1]
			}
			if len(stk) == 0 {
				up[i] = -1
			} else {
				up[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		stk = []int{}
		for i := m - 1; i >= 0; i-- {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= left[i][j] {
				stk = stk[:len(stk)-1]
			}
			if len(stk) == 0 {
				down[i] = m
			} else {
				down[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		for i := 0; i < m; i++ {
			height := down[i] - up[i] - 1
			area := height * left[i][j]
			if area > ret {
				ret = area
			}
		}
	}
	return ret
}

func main() {
	var m, n int
	if _, err := fmt.Scan(&m, &n); err != nil {
		return
	}
	matrix := make([][]byte, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			var s string
			fmt.Scan(&s)
			if len(s) > 0 {
				matrix[i][j] = s[0]
			}
		}
	}
	fmt.Println(maximalRectangle(matrix))
}
