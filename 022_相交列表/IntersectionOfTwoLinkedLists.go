package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func main() {
	var iv int
	fmt.Scan(&iv)
	var nA int
	fmt.Scan(&nA)
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

func getIntersectionNode(headA *Node, headB *Node) *Node {
	p, q := headA, headB
	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}
		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}
	return p
}
