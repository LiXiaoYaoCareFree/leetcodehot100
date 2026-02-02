package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
题目：轮转数组 (Rotate Array)
描述：给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
输入示例：
7
1 2 3 4 5 6 7
3
输出示例：
[5 6 7 1 2 3 4]
*/

func myrotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func myreverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
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

	// Read k
	if !scanner.Scan() {
		return
	}
	k, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	myrotate(nums, k)
	fmt.Println(nums)
}
