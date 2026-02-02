package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) (ans int) {
	f := make([]int, len(nums))
	f[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(f[i-1], 0) + nums[i]
	}
	return slices.Max(f)
}
