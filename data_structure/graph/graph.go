package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("1", "2")
	addEdge("1", "3")
	addEdge("2", "3")
	fmt.Println(hasEdge("1", "2"))
	fmt.Println(hasEdge("1", "3"))
	fmt.Println(hasEdge("2", "3"))
	fmt.Println(hasEdge("3", "1"))
}
