package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 617. 合并二叉树
// 给你两棵二叉树： root1 和 root2 。
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
// 返回合并后的二叉树。
// 注意: 合并过程必须从两个树的根节点开始。

// 示例 1：
// 输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// 输出：[3,4,5,5,4,null,7]

// 示例 2：
// 输入：root1 = [1], root2 = [1,2]
// 输出：[2,2]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
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

func printLevelOrder(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	var res []string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, "null")
		} else {
			res = append(res, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	// Trim trailing nulls
	for len(res) > 0 && res[len(res)-1] == "null" {
		res = res[:len(res)-1]
	}
	fmt.Println(res)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line1 := scanner.Text()
		if scanner.Scan() {
			line2 := scanner.Text()
			t1 := buildTree(line1)
			t2 := buildTree(line2)
			merged := mergeTrees(t1, t2)
			printLevelOrder(merged)
		}
	}
}
