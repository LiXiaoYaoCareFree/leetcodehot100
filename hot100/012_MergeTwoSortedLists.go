package main

import "fmt"

/*
题目：合并两个有序链表 (Merge Two Sorted Lists)
描述：将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
输入示例：
3
1 2 4
3
1 3 4
输出示例：
1 1 2 3 4 4
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
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
	
	if _, err := fmt.Scan(&n1); err != nil {
		return
	}
	nums1 := make([]int, n1)
	for i := 0; i < n1; i++ {
		fmt.Scan(&nums1[i])
	}
	l1 := buildList(nums1)

	if _, err := fmt.Scan(&n2); err != nil {
		return
	}
	nums2 := make([]int, n2)
	for i := 0; i < n2; i++ {
		fmt.Scan(&nums2[i])
	}
	l2 := buildList(nums2)

	res := mergeTwoLists(l1, l2)
	printList(res)
}
