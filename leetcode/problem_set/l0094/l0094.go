package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func inorderTraversalIterationMethod(root *TreeNode) []int {
	var nums []int
	var stack []*TreeNode
	if root == nil {
		return nums
	}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, cur.Val)
		cur = cur.Right

	}
	return nums
}

func main() {
	t := &TreeNode{Val: 1,
		Right: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 3}}}
	fmt.Println(inorderTraversal(t))
	fmt.Println(inorderTraversalIterationMethod(t))

	t1 := &TreeNode{Val: 2,
		Left: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 1}}}
	fmt.Println(inorderTraversal(t1))
	fmt.Println(inorderTraversalIterationMethod(t1))
}
