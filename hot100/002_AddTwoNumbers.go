package main

import (
	"fmt"
)

/*
题目：两数相加 (Add Two Numbers)
描述：给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
输入示例：
3
2 4 3
3
5 6 4
输出示例：
7 0 8
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return dummy.Next
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
	var n1, n2 int
	
	// Read List 1
	if _, err := fmt.Scan(&n1); err != nil {
		return
	}
	nums1 := make([]int, n1)
	for i := 0; i < n1; i++ {
		fmt.Scan(&nums1[i])
	}
	l1 := buildList(nums1)

	// Read List 2
	if _, err := fmt.Scan(&n2); err != nil {
		return
	}
	nums2 := make([]int, n2)
	for i := 0; i < n2; i++ {
		fmt.Scan(&nums2[i])
	}
	l2 := buildList(nums2)

	res := addTwoNumbers(l1, l2)
	printList(res)
}
