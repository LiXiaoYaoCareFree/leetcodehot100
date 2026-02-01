package main

import (
	"fmt"
	"math"
)

/*
题目：二叉树的右视图 (Binary Tree Right Side View)
描述：给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
输入示例：
1 2 3 null 5 null 4
输出示例：
[1 3 4]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == n-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

func buildTree(nodes []string) *TreeNode {
	if len(nodes) == 0 || nodes[0] == "null" {
		return nil
	}
	// Basic parsing from previous templates
	// Assume simple integer parsing, ignoring errors for brevity in this snippet
	// Real implementation should be robust
	val := 0
	fmt.Sscanf(nodes[0], "%d", &val)
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	i := 1
	for i < len(nodes) {
		parent := queue[0]
		queue = queue[1:]

		if i < len(nodes) && nodes[i] != "null" {
			v := 0
			fmt.Sscanf(nodes[i], "%d", &v)
			parent.Left = &TreeNode{Val: v}
			queue = append(queue, parent.Left)
		}
		i++

		if i < len(nodes) && nodes[i] != "null" {
			v := 0
			fmt.Sscanf(nodes[i], "%d", &v)
			parent.Right = &TreeNode{Val: v}
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
	fmt.Println(rightSideView(root))
}
