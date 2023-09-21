package main

import (
	"fmt"
)

func main() {
	//root1=[1,3,2,5]
	//root2=[2,1,3,null,4,null,7]
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}
	root1.Right = &TreeNode{Val: 2}

	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}
	root2.Right = &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}

	node := mergeTrees(root1, root2)
	middleRange(node)
}

// 官方
func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	mergeTrees(t1.Left, t2.Left)
	mergeTrees(t1.Right, t2.Right)
	return t1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度遍历（中序）
func middleRange(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val)
	middleRange(node.Left)
	middleRange(node.Right)
}
