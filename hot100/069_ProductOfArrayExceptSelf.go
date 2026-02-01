package main

import "fmt"

/*
题目：除自身以外数组的乘积 (Product of Array Except Self)
描述：给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
输入示例：
4
1 2 3 4
输出示例：
[24 12 8 6]
*/

func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	R := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		R *= nums[i]
	}
	return answer
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
	fmt.Println(productExceptSelf(nums))
}
