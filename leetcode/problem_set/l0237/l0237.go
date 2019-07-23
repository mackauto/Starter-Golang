package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	next := node.Next
	// current node become next node
	node.Val = next.Val
	node.Next = next.Next
	// release next node
	next.Next = nil
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
		Next: &ListNode{Val: 5,
			Next: &ListNode{Val: 1,
				Next: &ListNode{Val: 9}}}}
	l.Print()      // 4 5 1 9
	node := l.Next // node 5
	deleteNode(node)
	l.Print() // 4 1 9
}
