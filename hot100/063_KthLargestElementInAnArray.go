package main

import "fmt"

/*
题目：数组中的第K个最大元素 (Kth Largest Element in an Array)
描述：给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
输入示例：
6
3 2 1 5 6 4
2
输出示例：
5
*/

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func randomPartition(a []int, l, r int) int {
	i := l // Simple pivot selection (first element) to avoid rand import complexity in pure alg
	// Or use middle. For standard quick select, simple partition is fine for random data.
	// To be robust, swap mid or random to end.
	// Let's swap mid to right for pivot.
	mid := l + (r-l)/2
	a[mid], a[r] = a[r], a[mid]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
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
	fmt.Println(findKthLargest(nums, k))
}
