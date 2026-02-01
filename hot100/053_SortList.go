package main

import (
	"fmt"
	"sort"
)

/*
题目：排序链表 (Sort List)
描述：给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
输入示例：
4
4 2 1 3
输出示例：
1 2 3 4
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(tmp)
	h := &ListNode{}
	res := h
	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}
		h = h.Next
	}
	if left != nil {
		h.Next = left
	} else {
		h.Next = right
	}
	return res.Next
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

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
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
	res := sortList(head)
	printList(res)
}
