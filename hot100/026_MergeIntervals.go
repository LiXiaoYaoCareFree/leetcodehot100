package main

import (
	"fmt"
	"sort"
)

/*
题目：合并区间 (Merge Intervals)
描述：以数组 intervals 表示若干个区间的集合，其中单个区间 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
输入示例：
4
1 3
2 6
8 10
15 18
输出示例：
[[1 6] [8 10] [15 18]]
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{}
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

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
	res := merge(intervals)
	fmt.Println(res)
}
