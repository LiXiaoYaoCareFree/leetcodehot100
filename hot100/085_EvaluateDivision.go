package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 399. 除法求值
// 给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
// 另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。
// 返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。
// 注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

// 示例 1：
// 输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
// 输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
// 解释：
// 条件：a / b = 2.0, b / c = 3.0
// 问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
// 结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 并查集 + 权值
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, ok := id[a]; !ok {
			id[a] = len(id)
		}
		if _, ok := id[b]; !ok {
			id[b] = len(id)
		}
	}

	fa := make([]int, len(id))
	w := make([]float64, len(id))
	for i := range fa {
		fa[i] = i
		w[i] = 1.0
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			f := find(fa[x])
			w[x] *= w[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}

	merge := func(from, to int, val float64) {
		fFrom, fTo := find(from), find(to)
		if fFrom != fTo {
			fa[fFrom] = fTo
			w[fFrom] = val * w[to] / w[from]
		}
	}

	for i, eq := range equations {
		merge(id[eq[0]], id[eq[1]], values[i])
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, end := q[0], q[1]
		idStart, ok1 := id[start]
		idEnd, ok2 := id[end]
		if ok1 && ok2 && find(idStart) == find(idEnd) {
			ans[i] = w[idStart] / w[idEnd]
		} else {
			ans[i] = -1.0
		}
	}
	return ans
}

func main() {
	// Simplified ACM input:
	// N (number of equations)
	// a b val (N lines)
	// M (number of queries)
	// a b (M lines)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		n, _ := strconv.Atoi(line)
		equations := make([][]string, n)
		values := make([]float64, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			parts := strings.Fields(scanner.Text())
			equations[i] = []string{parts[0], parts[1]}
			values[i], _ = strconv.ParseFloat(parts[2], 64)
		}
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		queries := make([][]string, m)
		for i := 0; i < m; i++ {
			scanner.Scan()
			parts := strings.Fields(scanner.Text())
			queries[i] = []string{parts[0], parts[1]}
		}
		fmt.Println(calcEquation(equations, values, queries))
	}
}
