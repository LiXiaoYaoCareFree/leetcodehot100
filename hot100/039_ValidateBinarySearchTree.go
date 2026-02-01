package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
题目：验证二叉搜索树 (Validate Binary Search Tree)
描述：给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
输入示例：
2 1 3
输出示例：
true
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
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
	fmt.Println(isValidBST(root))
}
