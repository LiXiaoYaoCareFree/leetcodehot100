package main

import "fmt"

/*
题目：相交链表 (Intersection of Two Linked Lists)
描述：给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
输入示例：
8
4 1 8 4 5
5 6 1 8 4 5
2
3
输出示例：
Intersected at '8'
*/

type Node struct {
	Val  int
	Next *Node
}

func getIntersectionNode(headA, headB *Node) *Node {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

// Helper to build intersecting lists is complex in linear scan.
// Simulating via arrays and explicit intersection index.
// Skip complex input parsing for this specific runner, focus on logic.
// But to run it, we need something.
// Simplified: ListA, ListB, and then manually link.
// Input format: listA size, listA elements, listB size, listB elements, skipA, skipB
// intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
func main() {
	// Custom input parser for this problem to link nodes
	var iv int    // intersectVal
	fmt.Scan(&iv) // This is just for info in LC, but here we might need to construct.
	// Actually, let's just read list A and list B full content as if separate,
	// but we need to know where they intersect.
	// Standard input: intersectVal, listA, listB, skipA, skipB
	// But since we can't easily pointer link by value in generic way (values might duplicate),
	// we will construct two lists and manually link them based on skipA and skipB.

	// Read List A part before intersection
	var nA int
	fmt.Scan(&nA) // actually this is size of A
	valsA := make([]int, nA)
	for i := 0; i < nA; i++ {
		fmt.Scan(&valsA[i])
	}

	var nB int
	fmt.Scan(&nB)
	valsB := make([]int, nB)
	for i := 0; i < nB; i++ {
		fmt.Scan(&valsB[i])
	}

	var skipA, skipB int
	fmt.Scan(&skipA, &skipB)

	// Construct
	dummyA := &Node{}
	curA := dummyA
	for i := 0; i < skipA; i++ {
		curA.Next = &Node{Val: valsA[i]}
		curA = curA.Next
	}

	dummyB := &Node{}
	curB := dummyB
	for i := 0; i < skipB; i++ {
		curB.Next = &Node{Val: valsB[i]}
		curB = curB.Next
	}

	// Common part
	var commonHead *Node
	if skipA < nA { // If there is an intersection
		commonHead = &Node{Val: valsA[skipA]}
		cur := commonHead
		for i := skipA + 1; i < nA; i++ {
			cur.Next = &Node{Val: valsA[i]}
			cur = cur.Next
		}
	}

	curA.Next = commonHead
	curB.Next = commonHead

	res := getIntersectionNode(dummyA.Next, dummyB.Next)
	if res != nil {
		fmt.Printf("Intersected at '%d'\n", res.Val)
	} else {
		fmt.Println("No intersection")
	}
}
