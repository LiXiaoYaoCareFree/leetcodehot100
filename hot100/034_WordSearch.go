package main

import "fmt"

/*
题目：单词搜索 (Word Search)
描述：给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
输入示例：
3 4
A B C E
S F C S
A D E E
ABCCED
输出示例：
true
*/

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	var canFind func(r, c, i int) bool
	canFind = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}
		if r < 0 || r >= m || c < 0 || c >= n {
			return false
		}
		if used[r][c] || board[r][c] != word[i] {
			return false
		}
		used[r][c] = true
		canFindRest := canFind(r+1, c, i+1) || canFind(r-1, c, i+1) || canFind(r, c+1, i+1) || canFind(r, c-1, i+1)
		if canFindRest {
			return true
		}
		used[r][c] = false
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && canFind(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	var m, n int
	if _, err := fmt.Scan(&m, &n); err != nil {
		return
	}
	board := make([][]byte, m)
	for i := 0; i < m; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			var s string
			fmt.Scan(&s)
			if len(s) > 0 {
				board[i][j] = s[0]
			}
		}
	}
	var word string
	fmt.Scan(&word)
	fmt.Println(exist(board, word))
}
