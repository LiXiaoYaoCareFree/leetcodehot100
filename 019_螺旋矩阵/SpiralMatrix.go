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
	result := spiralOrder(matrix)
	fmt.Println(result)
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	u, d := 0, len(matrix)-1
	l, r := 0, len(matrix[0])-1

	ans := []int{}

	for {
		// 从左到右
		for i := l; i <= r; i++ {
			ans = append(ans, matrix[u][i])
		}
		u++
		if u > d {
			break
		}

		// 从上到下
		for i := u; i <= d; i++ {
			ans = append(ans, matrix[i][r])
		}
		r--
		if r < l {
			break
		}

		// 从右到左
		for i := r; i >= l; i-- {
			ans = append(ans, matrix[d][i])
		}
		d--
		if d < u {
			break
		}

		// 从下到上
		for i := d; i >= u; i-- {
			ans = append(ans, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}

	return ans
}
