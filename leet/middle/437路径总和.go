package main

import "fmt"

func main() {
	// 构建示例树：[10,5,-3,3,2,null,11,3,-2,null,1]
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: -3}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Right = &TreeNode{Val: -2}
	root.Left.Right.Right = &TreeNode{Val: 1}
	fmt.Println("Number of paths:", pathSum(root, 8))
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode, targetNum int) (count int)
	dfs = func(node *TreeNode, targetNum int) (count int) {
		if node == nil {
			return 0
		}
		// total 为从当前 node 节点出发，向下的所有路径中和为 targetNum 的路径数
		total := 0
		if node.Val == targetNum {
			total++
		}
		total += dfs(node.Left, targetNum-node.Val) // 注这里为 targetNum - node.Val
		total += dfs(node.Right, targetNum-node.Val)
		return total
	}
	t1 := dfs(root, targetSum)
	t2 := pathSum(root.Left, targetSum)
	t3 := pathSum(root.Right, targetSum)
	return t1 + t2 + t3
}
