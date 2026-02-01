package main

import "fmt"

/*
题目：全排列 (Permutations)
描述：给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
输入示例：
3
1 2 3
输出示例：
[[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
*/

func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(int)
	backtrack = func(first int) {
		if first == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			res = append(res, temp)
			return
		}
		for i := first; i < len(nums); i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return res
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
	res := permute(nums)
	fmt.Println(res)
}
