package main

import "fmt"

func main() {
	fmt.Println("aaa")
}

// 是将这棵树重新转变
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
func flatten(root *TreeNode) {
	nodes := make([]*TreeNode, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	temp := &TreeNode{}
	for _, v := range nodes {
		v.Left = nil
		temp.Right = v
		temp = temp.Right
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
