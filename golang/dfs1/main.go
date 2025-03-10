package main

import (
	"fmt"
)

// https://csacademy.com/app/graph_editor/
//
// Algorithm DFS(G, v)
//
//	if v is already visited
//	    return
//	Mark v as visited.
//	// Perform some operation on v.
//	for all neighbors x of v
//	    DFS(G, x)
func main() {
	g := map[int][]int{
		0: {2, 3, 5},
		1: {5},
		2: {3},
		3: {4},
		4: {7},
		6: {0},
	}

	println("start from 0")
	path := dfs(g, 0)
	fmt.Println("path:", path)

	println("start from 1")
	path = dfs(g, 1)
	fmt.Println("path:", path)

	println("start from 3")
	path = dfs(g, 3)
	fmt.Println("path:", path)
}

func dfs(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	return visit(graph, start, visited)
}

func visit(graph map[int][]int, start int, visited map[int]bool) []int {
	visited[start] = true
	path := []int{start}

	for _, node := range graph[start] {
		if !visited[node] {
			p := visit(graph, node, visited)
			path = append(path, p...)
		}
	}

	return path
}
