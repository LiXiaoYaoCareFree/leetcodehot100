package main

import "fmt"

/*
题目：跳跃游戏 (Jump Game)
描述：给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。
输入示例：
5
2 3 1 1 4
输出示例：
true
*/

func canJump(nums []int) bool {
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		if i <= rightmost {
			if i+nums[i] > rightmost {
				rightmost = i + nums[i]
			}
			if rightmost >= n-1 {
				return true
			}
		}
	}
	return false
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
	fmt.Println(canJump(nums))
}
