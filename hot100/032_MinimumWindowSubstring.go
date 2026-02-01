package main

import "fmt"

/*
题目：最小覆盖子串 (Minimum Window Substring)
描述：给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
输入示例：
ADOBECODEBANC
ABC
输出示例：
BANC
*/

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	lenVal := sLen + 1
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < lenVal {
				lenVal = r - l + 1
				ansL, ansR = l, l+lenVal
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)
	fmt.Println(minWindow(s, t))
}
