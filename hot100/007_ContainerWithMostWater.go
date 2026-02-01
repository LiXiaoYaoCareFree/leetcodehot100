package main

import "fmt"

/*
题目：盛最多水的容器 (Container With Most Water)
描述：给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
输入示例：
9
1 8 6 2 5 4 8 3 7
输出示例：
49
*/

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		if area > ans {
			ans = area
		}
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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
	fmt.Println(maxArea(height))
}
