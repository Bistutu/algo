package main

import "fmt"

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}

func main() {
	// Manually constructing the tree based on the provided image
	root := &TreeNode{4, nil, nil}
	root.Left = &TreeNode{1, nil, nil}
	root.Right = &TreeNode{6, nil, nil}
	root.Left.Left = &TreeNode{0, nil, nil}
	root.Left.Right = &TreeNode{2, nil, nil}
	root.Right.Left = &TreeNode{5, nil, nil}
	root.Right.Right = &TreeNode{7, nil, nil}
	root.Left.Right.Right = &TreeNode{3, nil, nil}
	root.Right.Right.Right = &TreeNode{8, nil, nil}

	// Function to print the tree in-order for verification
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node != nil {
			inOrder(node.Left)
			fmt.Println(node.Val)
			inOrder(node.Right)
		}
	}

	// Print the tree
	inOrder(root)
	convertBST(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
