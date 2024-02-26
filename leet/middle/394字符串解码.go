package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
}

func decodeString(s string) string {
	stack := make([]string, 0)
	num := 0

	for _, v := range s {
		if '0' <= v && v <= '9' {
			num = num*10 + int(v-'0')
		} else if v == '[' {
			stack = append(stack, strconv.Itoa(num))
			stack = append(stack, "[")
			num = 0 // 重新归零（这个 num 用的很巧妙）
		} else if v == ']' {
			temp := ""
			for len(stack) > 0 && stack[len(stack)-1] != "[" {
				temp = stack[len(stack)-1] + temp
				stack = stack[:len(stack)-1]
			}
			// 弹出 [
			stack = stack[:len(stack)-1]
			// 弹出 数字
			repeatTimes, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			// 解码并重新压入栈
			decodedString := strings.Repeat(temp, repeatTimes)
			stack = append(stack, decodedString)
		} else {
			// 字母时
			stack = append(stack, string(v))
		}
	}
	return strings.Join(stack, "")
}
