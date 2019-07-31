package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	memory := make(map[int]bool)
	for _, v := range to_delete {
		memory[v] = true
	}
	var forest []*TreeNode
	del(root, true, &forest, memory)
	return forest
}

func del(root *TreeNode, isRoot bool, forest *[]*TreeNode, memory map[int]bool) *TreeNode {
	if root == nil {
		return nil
	}
	deleted := memory[root.Val]
	if isRoot && !deleted {
		*forest = append(*forest, root)
	}
	root.Left = del(root.Left, deleted, forest, memory)
	root.Right = del(root.Right, deleted, forest, memory)
	if deleted {
		return nil
	} else {
		return root
	}
}

func main() {
	t := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7}}}
	forest := delNodes(t, []int{3, 5})
	for _, n := range forest {
		fmt.Println(n.Val)
	}
}
