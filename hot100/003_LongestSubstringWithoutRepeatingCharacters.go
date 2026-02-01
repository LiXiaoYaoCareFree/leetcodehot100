package main

import "fmt"

/*
题目：无重复字符的最长子串 (Longest Substring Without Repeating Characters)
描述：给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
输入示例：
abcabcbb
输出示例：
3
*/

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		if rk-i+1 > ans {
			ans = rk - i + 1
		}
	}
	return ans
}

func main() {
	var s string
	if _, err := fmt.Scan(&s); err != nil {
		return
	}
	fmt.Println(lengthOfLongestSubstring(s))
}
