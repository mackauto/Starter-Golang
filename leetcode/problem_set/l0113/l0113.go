package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findPathSum(root *TreeNode, sum int, path *[]int, paths *[][]int) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	// leaf node
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			c := make([]int, len(*path))
			copy(c, *path)
			*paths = append(*paths, c)
		}
		*path = (*path)[:len(*path)-1]
		return
	}
	findPathSum(root.Left, sum-root.Val, path, paths)
	findPathSum(root.Right, sum-root.Val, path, paths)
	*path = (*path)[:len(*path)-1]
}

func pathSum(root *TreeNode, sum int) [][]int {
	var paths [][]int
	var path []int
	findPathSum(root, sum, &path, &paths)
	return paths
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
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1}}}}
	fmt.Println(pathSum(t, 22))
}
