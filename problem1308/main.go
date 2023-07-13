package problem1310

import (
	"math"
)

type Edge struct {
	From string
	To   string
	Cost int
}

func solveWithPrim(in map[string]map[string]int) []Edge {
	finalEdges := make([]Edge, 0)
	seenNodes := map[string]struct{}{
		"plant": {}, // to have somewhere to start
	}

	for len(seenNodes) < len(in) {
		minCost := math.MaxInt
		var minEdge Edge

		for fromNode, edges := range in {
			if _, fromSeen := seenNodes[fromNode]; fromSeen {
				for toNode, cost := range edges {
					if _, toSeen := seenNodes[toNode]; !toSeen && cost < minCost {
						minCost = cost
						minEdge = Edge{From: fromNode, To: toNode, Cost: cost}
					}
				}
			}
		}

		finalEdges = append(finalEdges, minEdge)
		seenNodes[minEdge.To] = struct{}{}
	}

	return finalEdges
}

// Solution : using Prim's algorithm
func Solution(in map[string]map[string]int) []Edge {

	return solveWithPrim(in)
}
