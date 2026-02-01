package main

import "fmt"

/*
题目：搜索二维矩阵 II (Search a 2D Matrix II)
描述：编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
输入示例：
5 5
1 4 7 11 15
2 5 8 12 19
3 6 9 16 22
10 13 14 17 24
18 21 23 26 30
5
输出示例：
true
*/

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

func main() {
	var m, n int
	if _, err := fmt.Scan(&m, &n); err != nil {
		return
	}
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	var target int
	fmt.Scan(&target)
	fmt.Println(searchMatrix(matrix, target))
}
