package main

import "fmt"

/*
题目：寻找两个正序数组的中位数 (Median of Two Sorted Arrays)
描述：给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
输入示例：
2
1 3
1
2
输出示例：
2.00000
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		return float64(getKthElement(nums1, nums2, totalLength/2+1))
	}
	return float64(getKthElement(nums1, nums2, totalLength/2)+getKthElement(nums1, nums2, totalLength/2+1)) / 2.0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var m, n int
	fmt.Scan(&m)
	nums1 := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&nums1[i])
	}

	fmt.Scan(&n)
	nums2 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums2[i])
	}

	fmt.Printf("%.5f\n", findMedianSortedArrays(nums1, nums2))
}
