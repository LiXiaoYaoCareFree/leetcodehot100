package main

import "fmt"

/*
题目：最大子数组和 (Maximum Subarray)
描述：给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。
输入示例：
9
-2 1 -3 4 -1 2 1 -5 4
输出示例：
6
*/

func maxSubArray(nums []int) int {
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	return maxVal
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
	fmt.Println(maxSubArray(nums))
}
