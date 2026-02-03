package main

import "fmt"

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
	setZeros(matrix)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j])
			if j < n-1 {
				fmt.Print("")
			}
		}
		fmt.Println()
	}
}

func setZeros(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rowHasZero := make([]bool, m)
	colHasZero := make([]bool, n)

	for i, row := range matrix {
		for j, x := range row {
			if x == 0 {
				rowHasZero[i] = true
				colHasZero[j] = true
			}
		}
	}

	for i, row0 := range rowHasZero {
		for j, col0 := range colHasZero {
			if row0 || col0 {
				matrix[i][j] = 0
			}
		}
	}
}
