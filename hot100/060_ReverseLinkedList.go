package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 206. 反转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

// 示例 1：
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]

// 示例 2：
// 输入：head = [1,2]
// 输出：[2,1]

// 示例 3：
// 输入：head = []
// 输出：[]

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	var res []string
	for head != nil {
		res = append(res, strconv.Itoa(head.Val))
		head = head.Next
	}
	fmt.Println(strings.Join(res, " "))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var nums []int
		for _, p := range parts {
			val, _ := strconv.Atoi(p)
			nums = append(nums, val)
		}
		head := buildList(nums)
		newHead := reverseList(head)
		printList(newHead)
	}
}
