package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := treeDepth(root.Left)
	rightDepth := treeDepth(root.Right)
	return maxThree(leftDepth+rightDepth, diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(treeDepth(root.Left), treeDepth(root.Right))

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxThree(i, j, k int) int {
	m := max(i, j)
	return max(m, k)
}

func diameterOfBinaryTreeMethod1(root *TreeNode) int {
	ans := 0
	depth(root, &ans)
	return ans
}

func depth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	leftDepth := depth(root.Left, ans)
	rightDepth := depth(root.Right, ans)
	*ans = max(*ans, leftDepth+rightDepth)
	return 1 + max(leftDepth, rightDepth)
}

func main() {
	t := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3}}
	fmt.Println(diameterOfBinaryTree(t))
	fmt.Println(diameterOfBinaryTreeMethod1(t))
}
