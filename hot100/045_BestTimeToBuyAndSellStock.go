package main

import "fmt"

/*
题目：买卖股票的最佳时机 (Best Time to Buy and Sell Stock)
描述：给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
输入示例：
6
7 1 5 3 6 4
输出示例：
5
*/

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&prices[i])
	}
	fmt.Println(maxProfit(prices))
}
