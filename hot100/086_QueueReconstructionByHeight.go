package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 406. 根据身高重建队列
// 假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
// 请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。

// 示例 1：
// 输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
// 输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]

// 示例 2：
// 输入：people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
// 输出：[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})

	ans := make([][]int, 0)
	for _, person := range people {
		idx := person[1]
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)
	}
	return ans
}

func main() {
	// Input format:
	// h1 k1 h2 k2 ...
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		people := make([][]int, len(parts)/2)
		for i := 0; i < len(parts); i += 2 {
			h, _ := strconv.Atoi(parts[i])
			k, _ := strconv.Atoi(parts[i+1])
			people[i/2] = []int{h, k}
		}
		fmt.Println(reconstructQueue(people))
	}
}
