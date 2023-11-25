package main

import "fmt"

// TreeNode 是二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// main 函数创建一棵二叉树并调用 diameterOfBinaryTree 函数
func main() {
	// 创建一个简单的二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	// 调用函数并打印结果
	fmt.Println("Diameter of Binary Tree:", diameterOfBinaryTree(root))
}

// diameterOfBinaryTree 计算给定二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	// 初始化最大直径为0
	maxDiameter := 0

	// 定义一个深度优先搜索的递归函数
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		// 如果当前节点为空，返回0
		if node == nil {
			return 0
		}
		// 递归计算左子树的深度
		left := dfs(node.Left)
		// 递归计算右子树的深度
		right := dfs(node.Right)
		// 更新最大直径，当前直径为左右子树深度之和
		maxDiameter = max(maxDiameter, left+right)
		// 返回当前节点的最大深度
		return max(left, right) + 1
	}

	// 从根节点开始进行深度优先搜索
	dfs(root)
	// 返回计算得到的最大直径
	return maxDiameter
}

// max 返回两个整数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
