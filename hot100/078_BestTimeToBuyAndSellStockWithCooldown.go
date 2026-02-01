package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 309. 最佳买卖股票时机含冷冻期
// 给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

// 示例 1:
// 输入: prices = [1,2,3,0,2]
// 输出: 3 
// 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

// 示例 2:
// 输入: prices = [1]
// 输出: 0

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// dp[i][0]: 手上持有股票的最大收益
	// dp[i][1]: 手上不持有股票，且处于冷冻期中的最大收益
	// dp[i][2]: 手上不持有股票，且不处于冷冻期中的最大收益
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][2]-prices[i])))
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = int(math.Max(float64(dp[i-1][1]), float64(dp[i-1][2])))
	}
	return int(math.Max(float64(dp[n-1][1]), float64(dp[n-1][2])))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		prices := make([]int, len(parts))
		for i, p := range parts {
			prices[i], _ = strconv.Atoi(p)
		}

		fmt.Println(maxProfit(prices))
	}
}
