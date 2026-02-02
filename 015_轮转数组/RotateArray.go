package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func myrotate(nums []int, k int) {
	k %= len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Read n
	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(parts[i])
	}

	if !scanner.Scan() {
		return
	}
	k, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	myrotate(nums, k)
	fmt.Println(nums)
}
