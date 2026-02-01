package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 437. 路径总和 III
// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

// 示例 1：
// 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// 输出：3
// 解释：和等于 8 的路径有 3 条，如图所示。

// 示例 2：
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	preSum := map[int64]int{0: 1}
	var ans int
	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)
		ans += preSum[curr-int64(targetSum)]
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
	}
	dfs(root, 0)
	return ans
}

// Helper to build tree from level order array string (like LeetCode input)
// e.g. "10 5 -3 3 2 null 11"
func buildTree(input string) *TreeNode {
	if input == "" {
		return nil
	}
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}
	if parts[0] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(parts[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	i := 1
	for i < len(parts) {
		node := queue[0]
		queue = queue[1:]

		if i < len(parts) && parts[i] != "null" {
			val, _ := strconv.Atoi(parts[i])
			node.Left = &TreeNode{Val: val}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(parts) && parts[i] != "null" {
			val, _ := strconv.Atoi(parts[i])
			node.Right = &TreeNode{Val: val}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		root := buildTree(line)
		
		if scanner.Scan() {
			targetSum, _ := strconv.Atoi(scanner.Text())
			fmt.Println(pathSum(root, targetSum))
		}
	}
}
