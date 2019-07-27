package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		var current []int
		for levelSize > 0 {
			currentNode := queue[0]
			current = append(current, currentNode.Val)
			queue = queue[1:]
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
			levelSize--
		}
		result = append(result, current)
	}
	return result
}

func levelOrderDFSMethod(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	DFS(root, 0, &result)
	return result
}

func DFS(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}
	if len(*result) < level+1 {
		*result = append(*result, []int{})
	}
	(*result)[level] = append((*result)[level], root.Val)
	DFS(root.Left, level+1, result)
	DFS(root.Right, level+1, result)
}

func main() {
	t := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7}}}
	fmt.Println(levelOrder(t))

	fmt.Println(levelOrderDFSMethod(t))
}
