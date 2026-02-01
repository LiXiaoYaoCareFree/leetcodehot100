package main

import "fmt"

/*
题目：翻转二叉树 (Invert Binary Tree)
描述：给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
输入示例：
4 2 7 1 3 6 9
输出示例：
[4 7 2 9 6 3 1]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func buildTree(nodes []string) *TreeNode {
	if len(nodes) == 0 || nodes[0] == "null" {
		return nil
	}
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

// Level order print for verification
func printLevelOrder(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	q := []*TreeNode{root}
	var res []interface{}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			res = append(res, "null")
			continue
		}
		res = append(res, node.Val)
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
	invertTree(root)
	printLevelOrder(root)
}
