package main

import (
	"github.com/JPMike/Starter-Golang/algorithm/util"
)

func reverseKGroup(head *util.ListNode, k int) *util.ListNode {
	return &util.ListNode{}
}

func reverseListNode(start *util.ListNode, end *util.ListNode) *util.ListNode {
	prev := end
	for start != end {
		next := start.Next
		start.Next = prev
		prev = start
		start = next
	}
	return prev
}

func main() {
	listNodeLen := 5
	head := util.GetOrderedListNode(listNodeLen)
	util.PrintListNode(head)
	end := head
	for i := 1; i < listNodeLen; i++ {
		end = end.Next
	}
	head = reverseListNode(head, end)
	util.PrintListNode(head)
}
