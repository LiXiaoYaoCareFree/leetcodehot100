package main

import "fmt"

/*
题目：在排序数组中查找元素的第一个和最后一个位置 (Find First and Last Position of Element in Sorted Array)
描述：给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
输入示例：
6
5 7 7 8 8 10
8
输出示例：
[3 4]
*/

func searchRange(nums []int, target int) []int {
	leftmost := sortSearch(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sortSearch(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func sortSearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
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
	res := searchRange(nums, target)
	fmt.Println(res)
}
