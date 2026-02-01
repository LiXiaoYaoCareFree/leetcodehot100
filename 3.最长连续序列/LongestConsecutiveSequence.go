package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) (ans int) {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	for k, _ := range m {
		if m[k-1] {
			continue
		}
		tmp := k + 1
		for m[tmp] {
			tmp++
		}
		ans = max(ans, tmp-k)
	}
	return ans
}

/*

@'
6
100 4 200 1 3 2
'@ | go run .\LongestConsecutiveSequence.go
*/
