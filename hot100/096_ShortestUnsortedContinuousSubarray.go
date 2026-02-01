package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 581. 最短无序连续子数组
// 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
// 请你找出符合题意的 最短 子数组，并输出它的长度。

// 示例 1：
// 输入：nums = [2,6,4,8,10,9,15]
// 输出：5
// 解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。

// 示例 2：
// 输入：nums = [1,2,3,4]
// 输出：0

// 示例 3：
// 输入：nums = [1]
// 输出：0

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	maxn, right := math.MinInt64, -1
	minn, left := math.MaxInt64, -1
	for i := 0; i < n; i++ {
		if maxn > nums[i] {
			right = i
		} else {
			maxn = nums[i]
		}
		if minn < nums[n-1-i] {
			left = n - 1 - i
		} else {
			minn = nums[n-1-i]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		nums := make([]int, len(parts))
		for i, p := range parts {
			nums[i], _ = strconv.Atoi(p)
		}
		fmt.Println(findUnsortedSubarray(nums))
	}
}
