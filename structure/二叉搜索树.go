package main

import "fmt"

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
	root := NewNode(8) // 根节点

	// 第二层
	root.left = NewNode(4)
	root.right = NewNode(12)

	// 第三层
	root.left.left = NewNode(2)
	root.left.right = NewNode(6)
	root.right.left = NewNode(10)
	root.right.right = NewNode(14)

	// 第四层
	root.left.left.left = NewNode(1)
	root.left.left.right = NewNode(3)
	root.left.right.left = NewNode(5)
	root.left.right.right = NewNode(7)
	root.right.left.left = NewNode(9)
	root.right.left.right = NewNode(11)
	root.right.right.left = NewNode(13)
	root.right.right.right = NewNode(15)

	//root = nil
	root = root.insert(0)
	printBFS(root)
	fmt.Println()
	root = root.delete(3)
	root = root.delete(2)
	printBFS(root)
}
func (n *Node) delete(val int) *Node {
	var pre *Node // 用于保存当前节点的父节点
	cur := n      // 当前要检查的节点
	if n == nil {
		return nil // 如果树是空的，直接返回nil
	}
	// 寻找目标节点以及其父节点
	for cur != nil && cur.Val != val {
		pre = cur // 更新父节点
		if cur.Val < val {
			cur = cur.right // 如果当前节点值小于目标，去右子树中找
		} else {
			cur = cur.left // 否则，去左子树中找
		}
	}
	// 如果找不到目标值，直接返回原节点
	if cur == nil {
		return n
	}
	// 如果目标节点有两个子节点
	if cur.left != nil && cur.right != nil {
		// 寻找右子树中的最小值节点
		minNode := cur.right
		preOfMinNode := cur // 保存最小值节点的父节点
		for minNode.left != nil {
			preOfMinNode = minNode
			minNode = minNode.left // 一直向左走，直到找到最小值
		}
		// 用最小值节点的值替换目标节点的值
		cur.Val = minNode.Val
		// 删除最小值节点
		if preOfMinNode.left == minNode {
			preOfMinNode.left = minNode.right
		} else {
			preOfMinNode.right = minNode.right
		}
		return n
	} else {
		// 如果目标节点有一个或零个子节点
		var child *Node
		if cur.left != nil {
			child = cur.left
		} else {
			child = cur.right
		}
		// 如果要删除的节点不是根节点
		if pre != nil {
			if pre.left == cur {
				pre.left = child
			} else {
				pre.right = child
			}
		} else { // 如果要删除的节点是根节点
			n = child
		}
	}
	return n // 返回新的根节点（如果有的话）
}

// 搜索（应该得到 Node）
func (n *Node) search(num int) *Node {
	if n == nil {
		return nil
	}
	if n.Val > num {
		return n.left.search(num)
	} else if n.Val < num {
		return n.right.search(num)
	} else {
		return n
	}
}

// 插入
func (n *Node) insert(num int) *Node {
	var pre *Node
	var cur = n
	if n == nil {
		return &Node{Val: num, left: nil, right: nil}

	}
	for cur != nil {
		if cur.Val == num {
			return n
		}
		pre = cur
		if num > cur.Val {
			cur = cur.right
		} else {
			cur = cur.left
		}
	}
	if num > pre.Val {
		pre.right = &Node{Val: num}
	} else {
		pre.left = &Node{Val: num}
	}
	return n
}

// 深度优先中序遍历
func printBFS(root *Node) {
	if root == nil {
		return
	}
	printBFS(root.left)
	fmt.Printf("%d\t", root.Val)
	printBFS(root.right)
}
