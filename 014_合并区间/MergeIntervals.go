package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	intervals := make([][]int, n)
	for i := 0; i < n; i++ {
		intervals[i] = make([]int, 2)
		fmt.Scan(&intervals[i][0], &intervals[i][1])
	}
	res := mymerge(intervals)
	fmt.Println(res)
}

func mymerge(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] })
	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] {
			ans[m-1][1] = max(p[1], ans[m-1][1])
		} else {
			ans = append(ans, p)
		}
	}
	return
}
