package main

import (
	"fmt"
	"math"
)

// TODO: permutations
// TODO: shortest path
// TODO: getBests

var graph = map[string]map[string]int{
	"a": {"b": 1, "c": 3, "d": 4, "e": 4},
	"b": {"c": 2, "d": 2, "e": 5, "a": 3},
	"c": {"d": 4, "e": 1, "a": 7, "b": 2},
	"d": {"e": 5, "a": 6, "b": 1, "c": 1},
	"e": {"a": 3, "b": 2, "c": 3, "d": 1},
}

func main() {
	fmt.Println(partitionStarts(graph, 4))
	fmt.Println(permDist("dabced", graph))
}

func partitionStarts(graph map[string]map[string]int, parts int) [][]string {
	partitions := make([][]string, parts)
	idx := 0

	for key := range graph {
		partitionIdx := (idx + 1) % parts
		partition := partitions[partitionIdx]
		if partition == nil {
			partitions[partitionIdx] = []string{key}
		} else {
			partitions[partitionIdx] = append(partition, key)
		}
		idx++
	}

	return partitions
}

func permutations(nodes string, start string, graph map[string]map[string]int) int {
	bestPath := int(math.Inf(0))

	// TODO: permutations cannot reference itself when defined like this making
	// recursion problematic
	permutation := func(substr string, perm string) {
		if substr == "" && permDist(start+perm+start, graph) < bestPath {
			bestPath = permDist(start+perm+start, graph)
		} else {
			for i := 0; i < len(substr); i++ {
				permutation(substr[:i]+substr[:i+1], perm+substr[i])
			}
		}
	}

	permutation(nodes, "")

	return bestPath
}

func permDist(perm string, graph map[string]map[string]int) int {
	dist := 0
	for i := 0; i < len(perm)-1; i++ {
		nodeIdx := string(perm[i])
		adjacentNodeIdx := string(perm[i+1])
		distanceFromNodeToAdjacent := graph[nodeIdx][adjacentNodeIdx]
		dist += distanceFromNodeToAdjacent
	}
	return dist
}

func getBests()     {}
func shortestPath() {}
