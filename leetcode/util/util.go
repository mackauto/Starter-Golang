package util

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetExampleTree() *TreeNode {
	var t *TreeNode
	nums := GetExampleNums()
	for _, v := range nums {
		t = treeInsert(t, v)
	}
	return t
}

func treeInsert(t *TreeNode, val int) *TreeNode {
	if t == nil {
		return &TreeNode{Val: val, Left: nil, Right: nil}
	}
	if val < t.Val {
		t.Left = treeInsert(t.Left, val)
	} else {
		t.Right = treeInsert(t.Right, val)
	}
	return t
}

func (t *TreeNode) PrintMidOrder() {
	if t == nil {
		return
	}
	t.Left.PrintMidOrder()
	fmt.Print(t.Val, " ")
	t.Right.PrintMidOrder()
}

func (t *TreeNode) PrintPreOrder() {
	if t == nil {
		return
	}
	fmt.Print(t.Val, " ")
	t.Left.PrintPreOrder()
	t.Right.PrintPreOrder()
}

func (t *TreeNode) PrintPostOrder() {
	if t == nil {
		return
	}
	t.Left.PrintPostOrder()
	t.Right.PrintPostOrder()
	fmt.Print(t.Val, " ")
}

func GetOrderedListNode(count int, dummyHead bool) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	cur := head
	for i := 1; i <= count; i++ {
		node := &ListNode{Val: i, Next: nil}
		cur.Next = node
		cur = node
	}
	if dummyHead {
		return head
	}
	return head.Next
}

func PrintListNode(head *ListNode, dummyHead bool) {
	if dummyHead {
		head = head.Next
	}
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print(" ")
		head = head.Next
	}
	fmt.Print("\n")
}

func GetExampleNums() []int {
	return []int{4, 1, 5, 8, 9, 3, 2, 6, 7}
}

func PrintNums(s []int) {
	for _, v := range s {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}

func Swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
