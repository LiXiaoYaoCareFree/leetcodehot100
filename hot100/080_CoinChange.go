package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 322. 零钱兑换
// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
// 你可以认为每种硬币的数量是无限的。

// 示例 1：
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3 
// 解释：11 = 5 + 5 + 1

// 示例 2：
// 输入：coins = [2], amount = 3
// 输出：-1

// 示例 3：
// 输入：coins = [1], amount = 0
// 输出：0

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Expecting two lines:
	// 1 2 5
	// 11
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) == 2 {
			coinsParts := strings.Fields(lines[0])
			coins := make([]int, len(coinsParts))
			for i, p := range coinsParts {
				coins[i], _ = strconv.Atoi(p)
			}
			amount, _ := strconv.Atoi(lines[1])
			fmt.Println(coinChange(coins, amount))
			lines = []string{}
		}
	}
}
