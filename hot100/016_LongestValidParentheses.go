package main

import "fmt"

/*
题目：最长有效括号 (Longest Valid Parentheses)
描述：给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
输入示例：
(()
输出示例：
2
*/

func longestValidParentheses(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			if dp[i] > maxAns {
				maxAns = dp[i]
			}
		}
	}
	return maxAns
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(longestValidParentheses(s))
}
