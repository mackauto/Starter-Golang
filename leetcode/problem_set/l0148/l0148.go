package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// recursive method
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	sortedL := sortList(head.Next)
	head.Next = sortedL
	prev := head
	cur := sortedL
	for cur != nil && cur.Val < prev.Val {
		cur.Val, prev.Val = prev.Val, cur.Val
		prev = cur
		cur = cur.Next
	}
	return head
}

func (l *ListNode) Print() {
	cur := l
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func main() {
	l := &ListNode{Val: 4,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 1,
				Next: &ListNode{Val: 3}}}}
	l.Print()
	l = sortList(l)
	l.Print()

	l1 := &ListNode{Val: -1,
		Next: &ListNode{Val: 5,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 0}}}}}
	l1.Print()
	l1 = sortList(l1)
	l1.Print()
}
