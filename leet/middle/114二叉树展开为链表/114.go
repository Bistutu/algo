package main

import "fmt"

func flatten(root *TreeNode) {
	res := make([]*TreeNode, 0)

	temp := &TreeNode{Right: root}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			res = append(res, node)
		} else {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	for _, v := range res {
		v.Left = nil
		temp.Right = v
		temp = v
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 6}

	// 展开为链表
	flatten(root)

	// 打印展开后的链表
	for node := root; node != nil; node = node.Right {
		fmt.Printf("%d ", node.Val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
