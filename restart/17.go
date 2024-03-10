package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	size := len(digits)
	rs := make([]string, 0)
	if size == 0 {
		return rs
	}

	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var backtrack func(index int, now string)
	backtrack = func(index int, now string) {
		if index == size {
			rs = append(rs, now)
			return
		}
		// 获取当前数字
		num := digits[index]
		// 获取当前数字对应的字母
		letters := phoneMap[num]
		for _, v := range letters {
			backtrack(index+1, now+string(v))
		}
	}
	backtrack(0, "")

	return rs
}
