package graph_loops

import "github.com/lruggieri/daily-coding-problems/util"

// Solution : starts from a random node and proceeds depth first
// O(N * V), can be reduced to O(N + V)
func Solution(g map[int][]int) [][]int {
	loops := make([][]int, 0)
	queue := make([]int, 0)

	var firstNode int
	for firstNode = range g {
		// TODO can be improved by maintaining a map of nodes that were already explored
		solveDepthFirst(g, firstNode, queue, &loops)
	}

	return loops
}

func getLoop(queue []int, nodeValue int) []int {
	for i := 0; i < len(queue)-1; i++ {
		if queue[i] == nodeValue {
			return queue[i:]
		}
	}

	return nil
}

func loopsContain(loops [][]int, newLoop []int) bool {
	for _, l := range loops {
		if util.SliceSameUnordered(l, newLoop) {
			return true
		}
	}

	return false
}

func solveDepthFirst(g map[int][]int, node int, queue []int, loops *[][]int) {
	if loop := getLoop(queue, node); loop != nil {
		if !loopsContain(*loops, loop) {
			*loops = append(*loops, loop)
		}
	} else {
		queue = append(queue, node)
		for _, connection := range g[node] {
			solveDepthFirst(g, connection, queue, loops)
		}
	}
}
