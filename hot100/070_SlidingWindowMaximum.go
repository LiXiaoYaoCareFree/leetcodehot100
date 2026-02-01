package main

import "fmt"

/*
题目：滑动窗口最大值 (Sliding Window Maximum)
描述：给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。
输入示例：
8
1 3 -1 -3 5 3 6 7
3
输出示例：
[3 3 5 5 6 7]
*/

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
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
	var k int
	fmt.Scan(&k)
	fmt.Println(maxSlidingWindow(nums, k))
}
