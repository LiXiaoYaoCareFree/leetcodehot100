package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 337. 打家劫舍 III
// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
// 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。

// 示例 1:
// 输入: root = [3,2,3,null,3,null,1]
// 输出: 7 
// 解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7

// 示例 2:
// 输入: root = [3,4,5,1,3,null,1]
// 输出: 9
// 解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	val := dfs(root)
	return int(math.Max(float64(val[0]), float64(val[1])))
}

// returns [not_rob, rob]
func dfs(node *TreeNode) [2]int {
	if node == nil {
		return [2]int{0, 0}
	}
	left := dfs(node.Left)
	right := dfs(node.Right)
	
	// Not rob this node
	notRob := int(math.Max(float64(left[0]), float64(left[1]))) + int(math.Max(float64(right[0]), float64(right[1])))
	
	// Rob this node
	rob := node.Val + left[0] + right[0]
	
	return [2]int{notRob, rob}
}

// Helper to build tree from level order array string (like LeetCode input)
// e.g. "3 2 3 null 3 null 1"
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
		root := buildTree(line)
		fmt.Println(rob(root))
	}
}
