package main

import "fmt"

func main() {
	var s string
	if _, err := fmt.Scan(&s); err != nil {
		return
	}
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) (ans int) {
	left := 0
	m := map[rune]int{}
	for k, v := range s {
		m[v]++
		for m[v] > 1 {
			m[rune(s[left])]--
			left++
		}
		ans = max(ans, k-left+1)
	}
	return
}
