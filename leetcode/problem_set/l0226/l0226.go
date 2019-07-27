package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderPrint(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderPrint(root.Left)
	fmt.Print(root.Val, " ")
	inOrderPrint(root.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := invertTree(root.Right)
	left := invertTree(root.Left)
	root.Left = right
	root.Right = left
	return root
}

func main() {
	t := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9}}}
	inOrderPrint(t)
	fmt.Println()
	invertTree(t)
	inOrderPrint(t)
}
