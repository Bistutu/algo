package main

import "fmt"

func main() {
	fmt.Println(hanota([]int{2, 1, 0}, []int{}, []int{}))
}

// 汉诺塔
func hanota(A []int, B []int, C []int) []int {
	var dfs func(num int, src *[]int, buf *[]int, tar *[]int)
	dfs = func(num int, src *[]int, buf *[]int, tar *[]int) {
		if num == 1 {
			move(src, tar)
			return
		}
		dfs(num-1, src, tar, buf)
		move(src, tar)
		dfs(num-1, buf, src, tar)
	}
	dfs(len(A), &A, &B, &C)
	return C
}

// 移动元素
func move(src *[]int, tar *[]int) {
	temp := (*src)[len(*src)-1]
	*src = (*src)[:len(*src)-1]
	*tar = append(*tar, temp)
}
