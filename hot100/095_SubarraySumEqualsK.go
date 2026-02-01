package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 560. 和为 K 的子数组
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

// 示例 1：
// 输入：nums = [1,1,1], k = 2
// 输出：2

// 示例 2：
// 输入：nums = [1,2,3], k = 3
// 输出：2

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for _, num := range nums {
		pre += num
		if v, ok := m[pre-k]; ok {
			count += v
		}
		m[pre]++
	}
	return count
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
		
		if scanner.Scan() {
			k, _ := strconv.Atoi(scanner.Text())
			fmt.Println(subarraySum(nums, k))
		}
	}
}
