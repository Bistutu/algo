package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

func wordBreak(s string, wordDict []string) bool {
	// 初始化 wordDict set，用于快速判断
	wordDictSet := make(map[string]struct{}, len(wordDict))
	for _, v := range wordDict {
		wordDictSet[v] = struct{}{}
	}

	size := len(s)
	dp := make(map[int]bool, size+1)
	dp[0] = true

	for i := 1; i < size+1; i++ {
		for j := 0; j < i; j++ {
			// 判断 dp[j] 和剩下的部分 s[j:i] 是否在 wordDictSet 中，如在则说明 dp[i] 应为 true
			if _, ok := wordDictSet[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[size]
}
