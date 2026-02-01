package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 538. 把二叉搜索树转换为累加树
// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

// 示例 1：
// 输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
// 输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]

// 示例 2：
// 输入：root = [0,null,1]
// 输出：[1,null,1]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}

// Helper to build tree from level order array string
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

// Helper to print tree in level order to verify output
func printLevelOrder(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	var res []string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, "null")
		} else {
			res = append(res, strconv.Itoa(node.Val))
			if node.Left != nil || node.Right != nil || len(queue) > 0 { // Optimization: don't add nulls for leaves if not needed? 
				// Actually standard LC format keeps nulls if they are needed to determine structure
				// For simple verification, we'll just push children
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
	}
	// Trim trailing nulls
	for len(res) > 0 && res[len(res)-1] == "null" {
		res = res[:len(res)-1]
	}
	fmt.Println(res)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		root := buildTree(line)
		newRoot := convertBST(root)
		printLevelOrder(newRoot)
	}
}
