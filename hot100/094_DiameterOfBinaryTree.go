package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 543. 二叉树的直径
// 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

// 示例 :
// 给定二叉树
//           1
//          / \
//         2   3
//        / \     
//       4   5    
// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		if l+r > ans {
			ans = l + r
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}
	dfs(root)
	return ans
}

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		root := buildTree(line)
		fmt.Println(diameterOfBinaryTree(root))
	}
}
