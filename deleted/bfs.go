package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func main() {
	root := &TreeNode{val: 1, left: &TreeNode{val: 2, left: &TreeNode{val: 4}, right: &TreeNode{val: 5}}, right: &TreeNode{val: 3, left: &TreeNode{val: 6}, right: &TreeNode{val: 7}}}
	var bfs func(node *TreeNode)
	bfs = func(node *TreeNode) {
		queue := list.List{}
		queue.PushBack(node)

		for queue.Len() != 0 {
			// 取出元素、移除并访问
			back := queue.Front()
			queue.Remove(back)
			fmt.Println("访问：", back.Value.(*TreeNode).val)

			// 检测左侧与右侧元素
			if left := back.Value.(*TreeNode).left; left != nil {
				queue.PushBack(left)
			}
			if right := back.Value.(*TreeNode).right; right != nil {
				queue.PushBack(right)
			}
		}
	}

	bfs(root) // 调用
}
