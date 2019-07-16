package main

import (
	"fmt"
	"github.com/JPMike/Starter-Golang/leetcode/util"
)

func kthLastRecursiveMethod(head *util.ListNode, k int) *util.ListNode {
	var counter int
	return kthLastCounter(head, k, &counter)
}

func kthLastCounter(head *util.ListNode, k int, counter *int) *util.ListNode {
	// base condition
	if head == nil {
		return head
	}
	// things to do in the unit
	node := kthLastCounter(head.Next, k, counter)
	*counter = *counter + 1
	if *counter == k {
		return head
	}
	return node
}

func main() {
	head := util.GetOrderedListNode(7, false)
	util.PrintListNode(head, false)
	kth := kthLastRecursiveMethod(head, 2)
	fmt.Print(kth.Val)
}
