package main

import "fmt"

func main() {
	// root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
	nums := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	root := insertNode(nums, 0)
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return right
	}
	if left != nil {
		return left
	}
	return right
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertNode(nums []int, index int) *TreeNode {
	if index >= len(nums) || nums[index] == -1 {
		return nil
	}
	node := &TreeNode{Val: nums[index]}
	node.Left = insertNode(nums, 2*index+1)
	node.Right = insertNode(nums, 2*index+2)
	return node
}
