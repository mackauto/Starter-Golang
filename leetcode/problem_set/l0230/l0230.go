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

func kthSmallest1(root *TreeNode, k int) int {
	var nums []int
	inOrder1(root, &nums)
	return nums[k-1]
}

func inOrder1(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrder1(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrder1(root.Right, nums)
}

func kthSmallest2(root *TreeNode, k int) int {
	var target, cnt int
	DFS(root, &cnt, &target, k)
	return target
}

func DFS(root *TreeNode, cnt, target *int, k int) bool {
	if root == nil {
		return false
	}
	found := DFS(root.Left, cnt, target, k)
	if found {
		return true
	}
	*cnt++
	if *cnt == k {
		*target = root.Val
		return true
	}
	found = DFS(root.Right, cnt, target, k)
	return found
}

func kthSmallest3(root *TreeNode, k int) int {
	var stack []*TreeNode
	for node := root; node != nil; node = node.Left {
		stack = append(stack, node)
	}
	for cnt := 1; len(stack) != 0; cnt++ {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cnt == k {
			return node.Val
		}
		for node = node.Right; node != nil; node = node.Left {
			stack = append(stack, node)
		}
	}
	return -1
}

func kthSmallest4(root *TreeNode, k int) int {
	leftSize := getSize(root.Left)
	switch {
	case k <= leftSize:
		return kthSmallest4(root.Left, k)
	case leftSize+1 < k:
		return kthSmallest4(root.Right, k-leftSize-1)
	default:
		return root.Val
	}
}

func getSize(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getSize(root.Left) + getSize(root.Right)
}

func main() {
	root := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 1,
			Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 4}}
	fmt.Println(kthSmallest(root, 1))

	fmt.Println(kthSmallest1(root, 1))

	fmt.Println(kthSmallest2(root, 1))

	fmt.Println(kthSmallest3(root, 1))

	fmt.Println(kthSmallest4(root, 1))
}
