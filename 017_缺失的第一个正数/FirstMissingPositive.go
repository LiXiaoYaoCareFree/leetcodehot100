package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Read n
	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	// Read nums
	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(parts[i])
	}

	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range n {
		for 1 <= nums[i] && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			j := nums[i] - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i := range n {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
