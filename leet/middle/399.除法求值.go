package main

import "fmt"

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}

	result := calcEquation(equations, values, queries)

	for _, val := range result {
		fmt.Println(val)
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 1、初始化图
	graph := make(map[string]map[string]float64)
	for k, v := range equations {
		a, b := v[0], v[1]
		// 构建 a -> b 的边
		if graph[a] == nil {
			graph[a] = make(map[string]float64)
		}
		graph[a][b] = values[k]

		// 构建 b -> a 的边
		if graph[b] == nil {
			graph[b] = make(map[string]float64)
		}
		graph[b][a] = 1 / values[k]
	}

	// 2、定义 dfs，查找图中 start -> end 的路径
	var dfs func(start, end string, visited map[string]bool) float64
	dfs = func(start, end string, visited map[string]bool) float64 {
		if v, ok := graph[start][end]; ok {
			return v
		}
		// 遍历 start 节点的所有邻居节点
		for neighbor := range graph[start] {
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			// 递归的在邻居节点进行 dfs
			temp := dfs(neighbor, end, visited)
			if temp != -1.0 {
				return graph[start][neighbor] * temp
			}
			// 回退（回溯）
			visited[neighbor] = false
		}
		return -1
	}

	// 3、开始对 queries 进行结果遍历
	rs := make([]float64, len(queries))
	for k, v := range queries {
		start, end := v[0], v[1]
		if graph[start] == nil || graph[end] == nil {
			rs[k] = -1.0
			continue
		}
		if start == end {
			rs[k] = 1.0
			continue
		}
		visited := make(map[string]bool)
		rs[k] = dfs(start, end, visited)
	}

	return rs
}
