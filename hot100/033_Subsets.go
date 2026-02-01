package main

import (
	"fmt"
	"math"
)

/*
题目：子集 (Subsets)
描述：给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
输入示例：
3
1 2 3
输出示例：
[[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
*/

func subsets(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, set)
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	res := subsets(nums)
	fmt.Println(res)
}
