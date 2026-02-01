package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 207. 课程表
// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？

// 示例 1：
// 输入：numCourses = 2, prerequisites = [[1,0]]
// 输出：true
// 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

// 示例 2：
// 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
// 输出：false
// 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		indeg[p[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	visited := 0
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		visited++
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return visited == numCourses
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		numCourses, _ := strconv.Atoi(line)

		var prerequisites [][]int
		if scanner.Scan() {
			line2 := scanner.Text()
			// Assume input format for prerequisites: a b c d ... (pairs)
			// or standard LeetCode string representation.
			// Let's use simplified ACM input: pairs separated by space
			parts := strings.Fields(line2)
			for i := 0; i < len(parts); i += 2 {
				if i+1 < len(parts) {
					a, _ := strconv.Atoi(parts[i])
					b, _ := strconv.Atoi(parts[i+1])
					prerequisites = append(prerequisites, []int{a, b})
				}
			}
		}
		fmt.Println(canFinish(numCourses, prerequisites))
	}
}
