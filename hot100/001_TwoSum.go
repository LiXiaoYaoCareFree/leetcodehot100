package main

import (
	"fmt"
)

/*
题目：1.两数之和 (Two Sum)
描述：给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
输入示例：
4
2 7 11 15
9
输出示例：
[0 1]
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return nil
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
	var target int
	fmt.Scan(&target)

	res := twoSum(nums, target)
	fmt.Println(res)
}
