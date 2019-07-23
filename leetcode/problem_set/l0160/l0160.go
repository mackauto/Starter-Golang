package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	countA := 0
	countB := 0
	curA := headA
	curB := headB
	for curA != nil {
		countA++
		curA = curA.Next
	}
	for curB != nil {
		countB++
		curB = curB.Next
	}
	curA = headA
	curB = headB
	if countA >= countB {
		diff := countA - countB
		for diff > 0 {
			curA = curA.Next
			diff--
		}
	} else {
		diff := countB - countA
		for diff > 0 {
			curB = curB.Next
			diff--
		}
	}
	for curA != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}

func main() {
	nodeA1 := &ListNode{Val: 4}
	nodeA2 := &ListNode{Val: 1}
	nodeB1 := &ListNode{Val: 1}
	nodeC1 := &ListNode{Val: 8}
	nodeC2 := &ListNode{Val: 4}
	nodeA1.Next = nodeA2
	nodeA2.Next = nodeC1
	nodeC1.Next = nodeC2
	nodeB1.Next = nodeC1
	intersection := getIntersectionNode(nodeA1, nodeB1)
	if intersection != nil {
		fmt.Println(true, intersection.Val)
	} else {
		fmt.Println(false)
	}
}
