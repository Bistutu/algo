package main

import (
	"fmt"
	"math"
)

func main() {
	// 2,1,3
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	// 深度遍历搜索
	var dfs func(node *TreeNode, low, high int) bool
	dfs = func(node *TreeNode, low, high int) bool {
		if node == nil {
			return true
		}
		if node.Val <= low || node.Val >= high {
			return false
		}
		return dfs(node.Left, low, node.Val) && dfs(node.Right, node.Val, high)
	}

	return dfs(root, math.MinInt, math.MaxInt)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
