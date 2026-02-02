package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	res := threeSum(nums)
	fmt.Println((res))

}

func threeSum(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] {
			continue
		}
		if (x + nums[i+1] + nums[i+2]) > 0 {
			break
		}
		j, k := i+1, n-1
		for j < k {
			if x+nums[j]+nums[k] < 0 {
				j++
			} else if x+nums[j]+nums[k] > 0 {
				k--
			} else {
				ans = append(ans, []int{x, nums[k], nums[j]})
				for j++; j < k && nums[j] == nums[j+1]; j++ {
				}
				for k--; j < k && nums[k] == nums[k-1]; k-- {
				}
			}
		}
	}
	return
}
