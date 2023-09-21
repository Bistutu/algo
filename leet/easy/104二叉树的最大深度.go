package main

import "fmt"

func main() {
	fmt.Println(maxDepth(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}

// 递归计算左子树 (root.Left) 和右子树 (root.Right) 的最大深度，然后返回它们中的较大值
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
