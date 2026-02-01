package main

import "fmt"

/*
题目：括号生成 (Generate Parentheses)
描述：数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
输入示例：
3
输出示例：
[((())) (()()) (())() ()(()) ()()()]
*/

func generateParenthesis(n int) []string {
	res := new([]string)
	backtrack(res, "", 0, 0, n)
	return *res
}

func backtrack(res *[]string, current string, open int, close int, max int) {
	if len(current) == max*2 {
		*res = append(*res, current)
		return
	}

	if open < max {
		backtrack(res, current+"(", open+1, close, max)
	}
	if close < open {
		backtrack(res, current+")", open, close+1, max)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	res := generateParenthesis(n)
	fmt.Println(res)
}
