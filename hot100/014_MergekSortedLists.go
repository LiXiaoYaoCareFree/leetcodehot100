package main

import (
	"fmt"
)

/*
题目：合并K个升序链表 (Merge k Sorted Lists)
描述：给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
输入示例：
3
3
1 4 5
3
1 3 4
2
2 6
输出示例：
1 1 2 3 4 4 5 6
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	mid := l + (r-l)/2
	l1 := merge(lists, l, mid)
	l2 := merge(lists, mid+1, r)
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
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
	var k int
	if _, err := fmt.Scan(&k); err != nil {
		return
	}
	lists := make([]*ListNode, k)
	for i := 0; i < k; i++ {
		var n int
		fmt.Scan(&n)
		nums := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&nums[j])
		}
		lists[i] = buildList(nums)
	}

	res := mergeKLists(lists)
	printList(res)
}
