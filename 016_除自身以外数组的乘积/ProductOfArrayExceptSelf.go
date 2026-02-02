package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	pre[0] = 1
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}

	suf := make([]int, n)
	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	ans := make([]int, n)
	for i, p := range pre {
		ans[i] = p * suf[i]
	}
	return ans
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
	fmt.Println(productExceptSelf(nums))
}
