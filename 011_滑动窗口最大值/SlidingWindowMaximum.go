package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	var k int
	fmt.Scan(&k)
	fmt.Println(maxSlidingWindow(nums, k))
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	q := []int{}

	for i, x := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		left := i - k + 1
		if q[0] < left {
			q = q[1:]
		}
		if left >= 0 {
			ans[left] = nums[q[0]]
		}
	}
	return ans
}

/*

@'
8
1 3 -1 -3 5 3 6 7
3
'@ | go run .\SlidingWindowMaximum.go


*/
