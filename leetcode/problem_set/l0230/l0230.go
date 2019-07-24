package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	nums := inOrder(root)
	return nums[k-1]
}

func inOrder(root *TreeNode) []int {
	var nums []int
	if root == nil {
		return nums
	}
	nums = append(nums, inOrder(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, inOrder(root.Right)...)
	return nums
}

func main() {
	root := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 1,
			Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 4}}
	fmt.Println(kthSmallest(root, 1))
}
