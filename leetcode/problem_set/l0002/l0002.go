package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	cur := l
	var val1, val2, carry int
	for l1 != nil || l2 != nil || carry != 0 {
		val1 = 0
		val2 = 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		newNode := &ListNode{Val: (val1 + val2 + carry) % 10}
		carry = (val1 + val2 + carry) / 10
		cur.Next = newNode
		cur = newNode
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return l.Next
}

func main() {
	l1 := &ListNode{Val: 2,
		Next: &ListNode{Val: 4,
			Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5,
		Next: &ListNode{Val: 6,
			Next: &ListNode{Val: 4}}}
	l := addTwoNumbers(l1, l2)
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}

	fmt.Println()

	l3 := &ListNode{Val: 1}
	l4 := &ListNode{Val: 9,
		Next: &ListNode{Val: 9}}
	l = addTwoNumbers(l3, l4)
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
}
