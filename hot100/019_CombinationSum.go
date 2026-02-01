package main

import (
	"fmt"
	"sort"
)

/*
题目：组合总和 (Combination Sum)
描述：给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
输入示例：
4
2 3 6 7
7
输出示例：
[[2 2 3] [7]]
*/

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var dfs func(target, idx int, path []int)
	dfs = func(target, idx int, path []int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		dfs(target, idx+1, path)
		if target-candidates[idx] >= 0 {
			dfs(target-candidates[idx], idx, append(path, candidates[idx]))
		}
	}
	dfs(target, 0, []int{})
	return res
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	candidates := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&candidates[i])
	}
	var target int
	fmt.Scan(&target)
	res := combinationSum(candidates, target)
	fmt.Println(res)
}
