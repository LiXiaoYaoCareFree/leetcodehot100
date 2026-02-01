package main

import "fmt"

/*
题目：回文链表 (Palindrome Linked List)
描述：给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
输入示例：
4
1 2 2 1
输出示例：
true
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	head := buildList(nums)
	fmt.Println(isPalindrome(head))
}
