package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
题目：缺失的第一个正数 (First Missing Positive)
描述：给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
输入示例：
3
1 2 0
输出示例：
3
*/

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 只要 nums[i] 在 [1, n] 范围内，且 nums[i] 不在正确的位置上 (nums[i] != nums[nums[i]-1])
		// 就把它交换到正确的位置上
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Read n
	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	// Read nums
	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(parts[i])
	}

	fmt.Println(firstMissingPositive(nums))
}
