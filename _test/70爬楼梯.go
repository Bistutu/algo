package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findPaths(root *TreeNode, target int) [][]*TreeNode {
	var currentPath []*TreeNode
	var res [][]*TreeNode
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil || node.Val == 3 {
			return
		}
		currentPath = append(currentPath, node)
		if node.Val == target {
			temp := make([]*TreeNode, len(currentPath))
			copy(temp, currentPath)
			res = append(res, temp)
		}
		dfs(node.Left)
		dfs(node.Right)
		currentPath = currentPath[:len(currentPath)-1]
	}
	dfs(root)
	return res
}

func main() {
	// 构建一个示例二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}

	paths := findPaths(root, 7)
	// 遍历输出val
	for _, path := range paths {
		for _, node := range path {
			fmt.Print(node.Val, " ")
		}
		fmt.Println()
	}
}
