package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 297. 二叉树的序列化与反序列化
// 序列化是将一个数据结构或者对象转换为连续的比特位的过程，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

// 示例 1：
// 输入：root = [1,2,3,null,null,4,5]
// 输出：[1,2,3,null,null,4,5]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteString(",")
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	parts := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if len(parts) == 0 {
			return nil
		}
		valStr := parts[0]
		parts = parts[1:]
		if valStr == "null" || valStr == "" {
			return nil
		}
		val, _ := strconv.Atoi(valStr)
		node := &TreeNode{Val: val}
		node.Left = build()
		node.Right = build()
		return node
	}
	return build()
}

// Helper to build tree from level order array string (like LeetCode input)
// e.g. "1 2 3 null null 4 5"
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
	// Example input: 1 2 3 null null 4 5
	for scanner.Scan() {
		line := scanner.Text()
		root := buildTree(line)
		
		ser := Constructor()
		deser := Constructor()
		data := ser.serialize(root)
		ans := deser.deserialize(data)
		
		// Re-serialize to verify
		fmt.Println(ser.serialize(ans))
	}
}
