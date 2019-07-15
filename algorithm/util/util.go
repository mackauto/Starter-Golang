package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetOrderedListNode(count int) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	cur := head
	for i := 1; i <= count; i++ {
		node := &ListNode{Val: i, Next: nil}
		cur.Next = node
		cur = node
	}
	return head.Next
}

func PrintListNode(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print(" ")
		head = head.Next
	}
	fmt.Print("\n")
}
