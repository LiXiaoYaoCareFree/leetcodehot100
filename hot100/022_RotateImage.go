package main

import "fmt"

/*
题目：旋转图像 (Rotate Image)
描述：给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
输入示例：
3
3
1 2 3
4 5 6
7 8 9
输出示例：
[[7 4 1] [8 5 2] [9 6 3]]
*/

func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	// The input format requires reading n x n matrix.
	// But first input is usually size in competitive programming context.
	// Wait, the input example says:
	// 3 (size)
	// 3 (cols? usually just n x n so we just need one dimension)
	// Let's assume standard competitive programming input:
	// First line: n (rows), m (cols) - here n=m
	// Or just n
	// The example input I wrote in comment:
	// 3
	// 3 (this might be redundant or rows/cols)
	// Let's stick to reading n, then maybe m if provided, or just n x n.
	// Re-reading example:
	// 3
	// 3
	// ...
	// Let's assume first is n.
	
	// Actually, let's just read n.
	// And if there is another number, read it (m).
	// For this specific problem matrix is n x n.
	
	var m int // consume potential second dimension input if user provides it
	// But let's stick to simplest: Read n, then read n*n elements.
	
	// Correction based on my previous files:
	// I usually read n, then array.
	// Here it is a matrix.
	
	// Let's try to read n.
	// The user might input rows cols.
	
	// Let's just read n.
	
	var cols int
	fmt.Scan(&cols) // Assuming first input is n (rows)
	// If the user inputs "3 3", the second 3 will be read as first element of matrix if I am not careful.
	// But standard leetcode input usually is just the array.
	// My previous examples used explicit count.
	
	// Let's assume input is:
	// n
	// row1...
	// row2...
	
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n) // It is n x n
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	
	rotate(matrix)
	fmt.Println(matrix)
}
