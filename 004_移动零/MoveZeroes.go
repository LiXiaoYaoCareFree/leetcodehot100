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
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		nums := make([]int, len(parts))
		for i, p := range parts {
			nums[i], _ = strconv.Atoi(p)
		}
		moveZeroes(nums)
		fmt.Println(nums)
	}
}

func moveZeroes(nums []int) {
	size := 0
	for _, v := range nums {
		if v != 0 {
			nums[size] = v
			size++
		}
	}
	clear(nums[size:])
}
