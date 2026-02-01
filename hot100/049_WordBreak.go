package main

import "fmt"

/*
题目：单词拆分 (Word Break)
描述：给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
输入示例：
leetcode
2
leet code
输出示例：
true
*/

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	var s string
	fmt.Scan(&s)
	var n int
	fmt.Scan(&n)
	wordDict := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&wordDict[i])
	}
	fmt.Println(wordBreak(s, wordDict))
}
