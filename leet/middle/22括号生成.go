package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))

}

func generateParenthesis(n int) []string {
	var res []string
	dfs(n, n, "", &res)
	return res
}

func dfs(left, right int, str string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
	}
	if left > 0 {
		dfs(left-1, right, str+"(", res)
	}
	if right > left {
		dfs(left, right-1, str+")", res)
	}
}
