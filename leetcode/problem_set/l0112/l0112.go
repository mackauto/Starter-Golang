package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	// leaf node
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		}
		return false
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func main() {
	t := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{Val: 4,
				Right: &TreeNode{Val: 1}}}}
	fmt.Println(hasPathSum(t, 22))
}
