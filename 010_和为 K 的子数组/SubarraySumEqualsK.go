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

		if scanner.Scan() {
			k, _ := strconv.Atoi(scanner.Text())
			fmt.Println(subarraySum(nums, k))
		}
	}
}

func subarraySum(nums []int, k int) (ans int) {
	S := make([]int, len(nums)+1)
	for idx, num := range nums {
		S[idx+1] = S[idx] + num
	}

	cnt := make(map[int]int, len(nums)+1)
	for _, s := range S {
		if val, ok := cnt[s-k]; ok {
			ans += val
		}
		cnt[s]++
	}
	return ans
}

/*

@'
1 1 1
2
'@ | go run .\SubarraySumEqualsK.go

*/
