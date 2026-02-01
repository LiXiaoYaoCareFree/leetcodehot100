package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	height := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&height[i])
	}
	fmt.Println(trap(height))
}

func trap(height []int) (ans int) {
	n := len(height)
	sum := 0
	lmax := make([]int, n)
	rmax := make([]int, n)
	lnum, rnum := 0, 0
	for i := 0; i < n-1; i++ {
		if rnum < height[i] {
			rnum = height[i]
		}
		rmax[i] = rnum
	}
	for i := n - 1; i >= 0; i-- {
		if lnum < height[i] {
			lnum = height[i]
		}
		lmax[i] = lnum
	}
	for i := 0; i < n-1; i++ {
		sum = min(rmax[i], lmax[i]) - height[i]
		ans += sum
	}
	return ans
}
