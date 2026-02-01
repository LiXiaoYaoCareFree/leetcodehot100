package main

import "fmt"

/*
题目：接雨水 (Trapping Rain Water)
描述：给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
输入示例：
12
0 1 0 2 1 0 1 3 2 1 2 1
输出示例：
6
*/

func trap(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	height := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&height[i])
	}
	fmt.Println(trap(height))
}
