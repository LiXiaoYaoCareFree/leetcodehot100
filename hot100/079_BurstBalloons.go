package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 312. 戳气球
// 有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
// 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 
// 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
// 求所能获得硬币的最大数量。

// 示例 1：
// 输入：nums = [3,1,5,8]
// 输出：167
// 解释：
// nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
// coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

// 示例 2：
// 输入：nums = [1,5]
// 输出：10

func maxCoins(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 0; i < n; i++ {
		val[i+1] = nums[i]
	}

	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += dp[i][k] + dp[k][j]
				if sum > dp[i][j] {
					dp[i][j] = sum
				}
			}
		}
	}
	return dp[0][n+1]
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

		fmt.Println(maxCoins(nums))
	}
}
