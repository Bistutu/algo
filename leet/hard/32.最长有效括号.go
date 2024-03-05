package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

// 秘诀：有个基准元素，当前下标 - 基准下标 = 本次长度
func longestValidParentheses(s string) int {
	maxLength := 0
	stack := []int{-1} // 初始化入栈，将 -1 作为哨兵元素压入栈

	for k, v := range s {
		if v == '(' {
			stack = append(stack, k)
		} else {
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
			// 如果此时栈为空，则说明没有与“闭括号”匹配的开括号
			if len(stack) == 0 {
				stack = append(stack, k)
			} else {
				maxLength = max(maxLength, k-stack[len(stack)-1])
			}
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
遍历字符串 s = ")()())"：
	第一个字符 )：栈中只有 -1。弹出 -1 并压入当前索引 0。此时栈：[0]。
	第二个字符 (：是开括号，压入其索引 1。此时栈：[0, 1]。
	第三个字符 )：是闭括号，弹出栈顶元素 1。计算长度：2 (当前索引) - 0 (栈顶元素) = 2。更新最大长度为 2。此时栈：[0]。
	第四个字符 (：是开括号，压入其索引 3。此时栈：[0, 3]。
	第五个字符 )：是闭括号，弹出栈顶元素 3。计算长度：4 (当前索引) - 0 (栈顶元素) = 4。更新最大长度为 4。此时栈：[0]。
	第六个字符 )：是闭括号，弹出栈顶元素 0。栈变空，压入当前索引 5。此时栈：[5]。
*/
