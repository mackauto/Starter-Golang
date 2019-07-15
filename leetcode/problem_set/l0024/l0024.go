package main

import "github.com/JPMike/Starter-Golang/leetcode/util"

func swapPairs(head *util.ListNode) *util.ListNode {
	// base condition
	if head == nil || head.Next == nil {
		return head
	}
	// things to do in this unit
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

func main() {
	head := util.GetOrderedListNode(5, false)
	util.PrintListNode(head, false)
	head = swapPairs(head)
	util.PrintListNode(head, false)
}
