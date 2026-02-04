package main

import (
	"fmt"
)

/*
题目：48. 旋转图像 (Rotate Image)
描述：给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
输入示例：
3 3
1 2 3
4 5 6
7 8 9
输出示例：
[[7 4 1] [8 5 2] [9 6 3]]
*/

func rotate(matrix [][]int) {
	n := len(matrix)
	// 1. 水平翻转 (上下颠倒)
	// 1 2 3      7 8 9
	// 4 5 6  =>  4 5 6
	// 7 8 9      1 2 3
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	// 2. 主对角线翻转 (转置)
	// 7 8 9      7 4 1
	// 4 5 6  =>  8 5 2
	// 1 2 3      9 6 3
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main() {
	var n int
	// 为了兼容 M N 格式的输入（尽管这里是 N x N），我们尝试读取两个数。
	// 如果输入是 "3 3"，我们读取两次。
	// 如果输入只有 "3"，这会有问题。
	// 但考虑到之前的文件都是 M N，我们这里假设输入包含列数（虽然对于 N x N 来说是冗余的）。
	var m int
	fmt.Scan(&n, &m)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	rotate(matrix)
	fmt.Println(matrix)
}
