package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Val   int
	left  *Node
	right *Node
}

func NewNode(v int) *Node {
	return &Node{
		Val:   v,
		left:  nil,
		right: nil,
	}
}

func main() {
	// 创建根节点
	root := NewNode(1)

	// 创建第二层
	root.left = NewNode(2)
	root.right = NewNode(3)

	// 创建第三层
	root.left.left = NewNode(4)
	root.left.right = NewNode(5)
	root.right.left = NewNode(6)
	root.right.right = NewNode(7)

	// 创建第四层
	root.left.left.left = NewNode(8)
	root.left.left.right = NewNode(9)
	root.left.right.left = NewNode(10)
	root.left.right.right = NewNode(11)
	root.right.left.left = NewNode(12)
	root.right.left.right = NewNode(13)
	root.right.right.left = NewNode(14)
	root.right.right.right = NewNode(15)

	BFS(root)
}

// DFS 前序（深度优先）遍历
func DFS(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d\t", root.Val)
	DFS(root.left)
	DFS(root.right)
}

// BFS 广度优先遍历
func BFS(root *Node) {
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		node := l.Remove(l.Front()).(*Node)
		fmt.Printf("%d\t", node.Val)
		if node.left != nil {
			l.PushBack(node.left)
		}
		if node.right != nil {
			l.PushBack(node.right)
		}
	}
}
