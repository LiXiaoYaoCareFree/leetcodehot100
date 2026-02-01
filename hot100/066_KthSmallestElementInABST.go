package main

import (
	"fmt"
)

/*
题目：二叉搜索树中第K小的元素 (Kth Smallest Element in a BST)
描述：给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
输入示例：
3 1 4 null 2
1
输出示例：
1
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
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

func main() {
	var nodes []string
	var val string
	// Read until a number is encountered for k, assuming k is the last input.
	// But k is an integer, nodes are strings (some "null").
	// A bit tricky to distinguish without count or line separation.
	// Standard input: tree on one line, k on next.
	// We'll read all into a buffer, then parse.
	// Last element is k.
	
	// Better approach for interactive/pipe: read everything.
	var allInputs []string
	for {
		_, err := fmt.Scan(&val)
		if err != nil {
			break
		}
		allInputs = append(allInputs, val)
	}
	
	if len(allInputs) == 0 {
		return
	}
	
	kStr := allInputs[len(allInputs)-1]
	k := 0
	fmt.Sscanf(kStr, "%d", &k)
	
	nodes = allInputs[:len(allInputs)-1]
	root := buildTree(nodes)
	fmt.Println(kthSmallest(root, k))
}
