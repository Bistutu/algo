package main

import (
	"fmt"
)

func levelOrder(root *TreeNode) [][]int {
	m := make(map[int][]int)
	var dfs func(node *TreeNode, n int)
	dfs = func(node *TreeNode, n int) {
		if node == nil {
			return
		}
		m[n] = append(m[n], node.Val)
		dfs(node.Left, n+1)
		dfs(node.Right, n+1)
	}
	dfs(root, 0)
	rs := make([][]int, len(m))
	for k, v := range m {
		rs[k] = v
	}
	return rs
}

func main() {
	root := &TreeNode{3, nil, nil}
	root.Left = &TreeNode{9, nil, nil}
	root.Right = &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}
	fmt.Println(levelOrder(root)) // Output: [[3],[9,20],[15,7]]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
