package main

import (
	"fmt"
	"strconv"
)

/*
题目：二叉树的中序遍历 (Binary Tree Inorder Traversal)
描述：给定一个二叉树的根节点 root ，返回它的 中序 遍历。
输入示例：
1 null 2 3
输出示例：
[1 3 2]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
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
	// Read all inputs until newline or EOF for tree construction
	var nodes []string
	var val string
	// Simple input reading: assume inputs are separated by space/newline
	// We read until we can. For interactive, this might be tricky without count.
	// But let's assume standard input flow.
	// We'll read line by line or just all tokens.
	
	// Since we don't know the count, we can just read until error.
	// But in ACM mode often there's N or line based.
	// Let's assume input is on a single line for simplicity in this template.
	
	// To make it robust for pipe input:
	for {
		_, err := fmt.Scan(&val)
		if err != nil {
			break
		}
		nodes = append(nodes, val)
	}

	root := buildTree(nodes)
	fmt.Println(inorderTraversal(root))
}
