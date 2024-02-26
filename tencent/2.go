package main

import (
	"fmt"
	"math"
)

func main() {
	// 构造一棵示例二叉树
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}
	// 计算好节点的数量并打印结果
	fmt.Println("Number of good nodes:", goodNodes(root)) // 输出：3
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	count := 0
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		if isPrime(sum) {
			count++
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, 0)
	return count
}

// 判断质数
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
