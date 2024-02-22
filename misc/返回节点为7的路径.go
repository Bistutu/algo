package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 前序 1、7、4、5、3、6、7
	// 中序 4、7、5、1、6、3、7
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	res := preOrderII(root)
	for _, v := range res {
		for _, vv := range v {
			print(vv.Val)
		}
		println()
	}
}

func preOrderII(root *TreeNode) [][]*TreeNode {
	res := make([][]*TreeNode, 0) // 结果
	path := make([]*TreeNode, 0)  // 路径

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node)
		if node.Val == 7 {
			temp := make([]*TreeNode, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		dfs(node.Left)
		dfs(node.Right)
		// 回退
		path = path[:len(path)-1]
	}
	dfs(root)

	return res
}
