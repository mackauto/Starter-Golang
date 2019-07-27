package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if leftDepth == 0 || rightDepth == 0 {
		return 1 + leftDepth + rightDepth
	}
	return 1 + min(leftDepth, rightDepth)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	t := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7}}}
	fmt.Println(minDepth(t))
}
