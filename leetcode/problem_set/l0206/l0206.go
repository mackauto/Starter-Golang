package main

import "github.com/JPMike/Starter-Golang/leetcode/util"

func reverseList(head *util.ListNode) *util.ListNode {
	var prev *util.ListNode
	cur := head
	for cur != nil {
		// remember the next node
		next := cur.Next
		// change cur node next
		cur.Next = prev
		// remember the prev node
		prev = cur
		// update cur node
		cur = next
	}
	return prev
}

func reverseListRecursiveMethod(head *util.ListNode) *util.ListNode {
	// base condition
	if head == nil || head.Next == nil {
		return head
	}
	// things to do in the unit
	newHead := reverseListRecursiveMethod(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	head := util.GetOrderedListNode(5, false)
	util.PrintListNode(head, false)
	head = reverseList(head)
	util.PrintListNode(head, false)

	head = util.GetOrderedListNode(5, false)
	util.PrintListNode(head, false)
	head = reverseListRecursiveMethod(head)
	util.PrintListNode(head, false)
}
