package main

import "fmt"

/*
题目：环形链表 (Linked List Cycle)
描述：给你一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
如果链表中存在环 ，返回 true 。 否则，返回 false 。
输入示例：
4
3 2 0 -4
1
输出示例：
true
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func buildList(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	cur := dummy
	var cycleNode *ListNode
	for i, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
		if i == pos {
			cycleNode = cur
		}
	}
	if pos != -1 {
		cur.Next = cycleNode
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
	var pos int
	fmt.Scan(&pos)
	head := buildList(nums, pos)
	fmt.Println(hasCycle(head))
}
