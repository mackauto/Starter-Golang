package main

import (
	"fmt"
	"github.com/JPMike/Starter-Golang/leetcode/util"
)

func mergeTrees(t1 *util.TreeNode, t2 *util.TreeNode) *util.TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func main() {
	t1 := util.GetExampleTree()
	t1.PrintInOrder()
	fmt.Println()
	t2 := util.GetExampleTree()
	t := mergeTrees(t1, t2)
	t.PrintInOrder()
}
