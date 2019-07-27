package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(i, j int) int {
	if i > j {
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
	fmt.Println(maxDepth(t))
}
