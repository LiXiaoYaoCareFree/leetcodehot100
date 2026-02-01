package main

import (
	"fmt"
)

/*
题目：从前序与中序遍历序列构造二叉树 (Construct Binary Tree from Preorder and Inorder Traversal)
描述：给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
输入示例：
5
3 9 20 15 7
9 3 15 20 7
输出示例：
[3 9 20 null null 15 7]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// Helper to print level order to verify the tree structure
func printLevelOrder(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	q := []*TreeNode{root}
	var res []string
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			res = append(res, "null")
			continue
		}
		res = append(res, fmt.Sprintf("%d", node.Val))
		q = append(q, node.Left)
		q = append(q, node.Right)
	}
	// Trim trailing nulls
	i := len(res) - 1
	for i >= 0 && res[i] == "null" {
		i--
	}
	fmt.Println(res[:i+1])
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	preorder := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&preorder[i])
	}
	inorder := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&inorder[i])
	}
	root := buildTree(preorder, inorder)
	printLevelOrder(root)
}
