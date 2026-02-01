package main

import "fmt"

/*
题目：删除链表的倒数第 N 个结点 (Remove Nth Node From End of List)
描述：给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
输入示例：
5
1 2 3 4 5
2
输出示例：
1 2 3 5
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
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
	var len int
	fmt.Scan(&len)
	nums := make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Scan(&nums[i])
	}
	var n int
	fmt.Scan(&n)

	head := buildList(nums)
	res := removeNthFromEnd(head, n)
	printList(res)
}
