package main

import (
	"fmt"
	"strconv"
)

/*
题目：二叉树展开为链表 (Flatten Binary Tree to Linked List)
描述：给你二叉树的根节点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而 left 子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
输入示例：
1 2 5 3 4 null 6
输出示例：
[1 null 2 null 3 null 4 null 5 null 6]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
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

func printFlattened(root *TreeNode) {
	var res []string
	curr := root
	for curr != nil {
		res = append(res, strconv.Itoa(curr.Val))
		res = append(res, "null") // Since left is always null
		curr = curr.Right
	}
	// Just print values for simplicity or follow output format
	// The example output [1 null 2 null ...] implies checking right pointers
	fmt.Println(res)
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
	flatten(root)
	printFlattened(root)
}
