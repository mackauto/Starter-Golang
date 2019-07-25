package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	swapTreeLR(root.Right)
	return isSameTree(root.Left, root.Right)
}

func swapTreeLR(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	swapTreeLR(root.Left)
	swapTreeLR(root.Right)
}

func isSameTree(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 != nil {
		return root1.Val == root2.Val && isSameTree(root1.Left, root2.Left) && isSameTree(root1.Right, root2.Right)
	}
	return false
}

func isSymmetricMirrorMethod(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 != nil {
		return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
	}
	return false
}

func main() {
	t1 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3}}}
	fmt.Println(isSymmetricMirrorMethod(t1))
	fmt.Println(isSymmetric(t1))

	t2 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2,
			Right: &TreeNode{Val: 3}}}
	fmt.Println(isSymmetricMirrorMethod(t2))
	fmt.Println(isSymmetric(t2))
}
