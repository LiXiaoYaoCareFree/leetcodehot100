package main

import (
	"fmt"
	"strconv"
)

/*
题目：二叉树的最大深度 (Maximum Depth of Binary Tree)
描述：给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
输入示例：
3 9 20 null null 15 7
输出示例：
3
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func buildTree(nodes []string) *TreeNode {
	if len(nodes) == 0 || nodes[0] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	i := 1
	for i < len(nodes) {
		parent := queue[0]
		queue = queue[1:]

		if i < len(nodes) && nodes[i] != "null" {
			val, _ := strconv.Atoi(nodes[i])
			parent.Left = &TreeNode{Val: val}
			queue = append(queue, parent.Left)
		}
		i++

		if i < len(nodes) && nodes[i] != "null" {
			val, _ := strconv.Atoi(nodes[i])
			parent.Right = &TreeNode{Val: val}
			queue = append(queue, parent.Right)
		}
		i++
	}
	return root
}

func main() {
	var nodes []string
	var val string
	for {
		_, err := fmt.Scan(&val)
		if err != nil {
			break
		}
		nodes = append(nodes, val)
	}

	root := buildTree(nodes)
	fmt.Println(maxDepth(root))
}
