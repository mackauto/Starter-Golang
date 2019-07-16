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

func kthLastIterationMethod(head *util.ListNode, k int) *util.ListNode {
	slow := head
	fast := head
	for k > 0 {
		if fast == nil {
			return nil
		}
		fast = fast.Next
		k -= 1
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func main() {
	listNodeLen := 7
	k := 6
	head := util.GetOrderedListNode(listNodeLen, false)
	util.PrintListNode(head, false)

	head1 := head
	kth1 := kthLastRecursiveMethod(head1, k)
	if kth1 != nil {
		fmt.Println(kth1.Val)
	}

	head2 := head
	kth2 := kthLastIterationMethod(head2, k)
	if kth2 != nil {
		fmt.Println(kth2.Val)
	}
}
