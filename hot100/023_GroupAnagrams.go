package main

import (
	"fmt"
	"sort"
)

/*
题目：字母异位词分组 (Group Anagrams)
描述：给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
输入示例：
6
eat tea tan ate nat bat
输出示例：
[[bat] [nat tan] [ate eat tea]]
*/

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&strs[i])
	}
	res := groupAnagrams(strs)
	// Output format might vary, just printing the slice
	// Note: The order in map iteration is random, so output order might differ from example
	fmt.Println(res)
}
