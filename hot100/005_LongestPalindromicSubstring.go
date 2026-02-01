package main

import "fmt"

/*
题目：最长回文子串 (Longest Palindromic Substring)
描述：给你一个字符串 s，找到 s 中最长的回文子串。
输入示例：
babad
输出示例：
bab
*/

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		l := max(len1, len2)
		if l > end-start {
			start = i - (l-1)/2
			end = i + l/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(longestPalindrome(s))
}
