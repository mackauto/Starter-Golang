package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}

func inorderTraversal(root *TreeNode) []int {
	var nums []int
	if root == nil {
		return nums
	}
	nums = append(nums, inorderTraversal(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, inorderTraversal(root.Right)...)
	return nums
}

func main() {
	tree := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 0,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 0}},
		Right: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 1}}}
	fmt.Println(inorderTraversal(tree))
	newTree := pruneTree(tree)
	fmt.Println(inorderTraversal(newTree))
}
