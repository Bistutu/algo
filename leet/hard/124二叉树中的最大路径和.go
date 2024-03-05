package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxPathSum(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}))
}

func maxPathSum(root *TreeNode) int {
	maxCount := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 求左右最大值（舍弃负值）
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)

		count := node.Val + left + right
		if count > maxCount {
			maxCount = count
		}
		return node.Val + max(left, right)
	}
	dfs(root)
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
