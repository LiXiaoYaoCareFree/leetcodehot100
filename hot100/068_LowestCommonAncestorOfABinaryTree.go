package main

import "fmt"

/*
题目：二叉树的最近公共祖先 (Lowest Common Ancestor of a Binary Tree)
描述：给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
输入示例：
3 5 1 6 2 0 8 null null 7 4
5
1
输出示例：
3
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
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

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	l := findNode(root.Left, val)
	if l != nil {
		return l
	}
	return findNode(root.Right, val)
}

func main() {
	var allInputs []string
	var val string
	for {
		_, err := fmt.Scan(&val)
		if err != nil {
			break
		}
		allInputs = append(allInputs, val)
	}
	
	if len(allInputs) < 3 {
		return
	}
	
	// Last two are p and q values
	pValStr := allInputs[len(allInputs)-2]
	qValStr := allInputs[len(allInputs)-1]
	pVal, qVal := 0, 0
	fmt.Sscanf(pValStr, "%d", &pVal)
	fmt.Sscanf(qValStr, "%d", &qVal)
	
	nodes := allInputs[:len(allInputs)-2]
	root := buildTree(nodes)
	
	p := findNode(root, pVal)
	q := findNode(root, qVal)
	
	res := lowestCommonAncestor(root, p, q)
	fmt.Println(res.Val)
}
