package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	// base condition
	if root == nil {
		return 0
	}
	// things to do in the unit
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	sum += rangeSumBST(root.Left, L, R)
	sum += rangeSumBST(root.Right, L, R)
	return sum
}

func main() {
	t := &TreeNode{Val: 10,
		Left: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7}},
		Right: &TreeNode{Val: 15,
			Right: &TreeNode{Val: 18}},
	}
	fmt.Println(rangeSumBST(t, 7, 15)) // 32
}
