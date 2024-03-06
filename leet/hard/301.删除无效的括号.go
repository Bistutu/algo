package main

import (
	"fmt"
	"math"
)

func main() {
	s := "()())()"
	result := removeInvalidParentheses(s)
	fmt.Println(result)
}

// 回溯遍历所有场景
func removeInvalidParentheses(s string) []string {
	rs := make([]string, 0) // 存储结果
	m := make(map[string]struct{})
	minRemoved := math.MaxInt // 最小移除括号（需要最先计算出该值），初始值为 math.最大值

	// 复杂的 dfs
	var dfs func(index, leftCount, rightCount, removeCount int, path string)
	dfs = func(index, leftCount, rightCount, removeCount int, path string) {

		// 1、如果遍历完整个字符串
		if index == len(s) {
			// 如果左右括号数量相等，那么本次形成了有效字符串
			if leftCount == rightCount {
				// 这里是依题目所求，只取删除符号最小时候的场景
				if removeCount < minRemoved {
					m = make(map[string]struct{})
					m[path] = struct{}{}
					minRemoved = removeCount
				} else if removeCount == minRemoved {
					m[path] = struct{}{}
				}
			}
			return
		}

		// 2、如果还没有遍历完整个字符串
		// 2.1 无论当前字符是什么，我们都有一个选择就是删掉它！
		dfs(index+1, leftCount, rightCount, removeCount+1, path)

		// 2.2 处理左右括号的情况
		if s[index] == '(' {
			// 添加左括号并递归
			dfs(index+1, leftCount+1, rightCount, removeCount, path+s[index:index+1])
		} else if s[index] == ')' {
			if leftCount > rightCount {
				// 添加右括号（仅当左括号数大于右括号数时）
				dfs(index+1, leftCount, rightCount+1, removeCount, path+s[index:index+1])
			}
		} else { // 2.3 如果是其他字符，则直接添加并递归
			dfs(index+1, leftCount, rightCount, removeCount, path+s[index:index+1])
		}
	}

	dfs(0, 0, 0, 0, "")

	for k := range m {
		rs = append(rs, k)
	}
	return rs
}
