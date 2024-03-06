package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	rs := make([]string, 0)
	// 前序遍历
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			rs = append(rs, "null")
			return
		}
		rs = append(rs, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(rs, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	// 前序遍历
	nodes := strings.Split(data, ",")
	num := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if nodes[num] == "null" {
			num++
			return nil
		}
		val, _ := strconv.Atoi(nodes[num])
		num++

		node := &TreeNode{Val: val}
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}

func main() {
	// 构造一个测试用的二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	// 初始化 Codec
	codec := Constructor()

	// 序列化二叉树
	serializedData := codec.serialize(root)
	fmt.Println("Serialized Data:", serializedData)

	// 反序列化二叉树
	deserializedTree := codec.deserialize(serializedData)

	// 再次序列化，以验证反序列化的结果
	reSerializedData := codec.serialize(deserializedTree)
	fmt.Println("Re-Serialized Data:", reSerializedData)
}
