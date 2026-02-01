package main

import (
	"fmt"
)

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

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		s := (right - left) * min(height[left], height[right])
		ans = max(s, ans)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}
