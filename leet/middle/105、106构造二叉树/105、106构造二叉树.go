package main

import (
	"fmt"
)

func main() {
	// 1、从前序遍历和中序遍历构造二叉树，见：https://www.hello-algo.com/chapter_divide_and_conquer/build_binary_tree_problem/#4
	root1 := buildTree1([]int{3, 9, 2, 1, 7}, []int{9, 3, 1, 2, 7})
	printTree(root1)

	// 2、从中序遍历和后序遍历构造二叉树，见：https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
	root2 := buildTree2([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	printTree(root2)
}

func buildTree2(inorder []int, postorder []int) *TreeNode {
	// map 存储 inorder 映射，方便获取
	inorderMap := make(map[int]int, len(inorder))
	for k, v := range inorder {
		inorderMap[v] = k
	}

	var dfs func(inorderMap map[int]int, postorder []int, i, l, r int) *TreeNode
	dfs = func(inorderMap map[int]int, postorder []int, i, l, r int) *TreeNode {
		if r-l < 0 {
			return nil
		}
		// 构建根节点
		root := newNode(postorder[i])
		m := inorderMap[root.Val]

		root.Left = dfs(inorderMap, postorder, i-1-(r-m), l, m-1)
		root.Right = dfs(inorderMap, postorder, i-1, m+1, r)

		return root
	}

	root := dfs(inorderMap, postorder, len(postorder)-1, 0, len(postorder)-1)

	return root
}

// 最主要是要事先知道需要定义：i、m、l、r 这四个变量，以及这四个变量的公式转换，还有我需要去做什么。
func buildTree1(preorder []int, inorder []int) *TreeNode {
	// map 存储 inorder 映射，方便获取
	inorderMap := make(map[int]int, len(inorder))
	for k, v := range inorder {
		inorderMap[v] = k
	}

	var dfs func(preorder []int, inorderMap map[int]int, i, l, r int) *TreeNode
	dfs = func(preorder []int, inorderMap map[int]int, i, l, r int) *TreeNode {
		if r-l < 0 {
			return nil
		}
		// 构建根节点
		root := newNode(preorder[i])
		m := inorderMap[root.Val]

		root.Left = dfs(preorder, inorderMap, i+1, l, m-1)
		root.Right = dfs(preorder, inorderMap, i+1+(m-l), m+1, r)

		return root
	}

	root := dfs(preorder, inorderMap, 0, 0, len(preorder)-1)

	return root
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// 中序遍历
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	printTree(root.Left)
	fmt.Println(root.Val)
	printTree(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
