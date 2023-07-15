package problem1310

import (
	"fmt"
)

type Graph []*Node

func (g Graph) GetNode(value string) *Node {
	for _, n := range g {
		if n.value == value {
			return n
		}
	}

	return nil
}

type Node struct {
	value       string
	connections []*Node
}

func (n *Node) ConnectedNodes(exclusions ...string) int {
	return n.connectedNodes(map[string]struct{}{}, exclusions...)
}

func (n *Node) connectedNodes(seen map[string]struct{}, exclusions ...string) int {
	seen[n.value] = struct{}{}
	connectedNodes := 1 // self

	isEdgeExcluded := func(edge string) bool {
		for _, e := range exclusions {
			if e == edge {
				return true
			}
		}

		return false
	}

	for _, connection := range n.connections {
		if _, ok := seen[connection.value]; !ok && !isEdgeExcluded(n.value+connection.value) {
			connectedNodes += connection.connectedNodes(seen, exclusions...)
		}
	}

	return connectedNodes
}

func (n *Node) Print() {
	n.print(map[string]struct{}{})
}

func (n *Node) print(seen map[string]struct{}) {
	if _, ok := seen[n.value]; !ok {
		seen[n.value] = struct{}{}
		fmt.Println(string(n.value))
		for _, c := range n.connections {
			fmt.Println("\t->", string(c.value))
		}

		for _, c := range n.connections {
			c.print(seen)
		}
	}
}

func buildGraphFromConfig(config map[string][]string) (g Graph) {
	seen := map[string]*Node{}

	for from, connections := range config {
		var fromNode *Node

		if t, ok := seen[from]; ok {
			fromNode = t
		} else {
			fromNode = &Node{value: from}
			seen[from] = fromNode
			g = append(g, fromNode)
		}

		for _, to := range connections {
			var toNode *Node
			if t, ok := seen[to]; ok {
				toNode = t
			} else {
				toNode = &Node{value: to}
				g = append(g, toNode)

				seen[to] = toNode
			}

			fromNode.connections = append(fromNode.connections, toNode)
		}
	}

	return g
}

// Solution : O(n^2). Not the best solution. Tarjan’s Algorithm performs in O(V+E)
func Solution(in Graph) []string {
	bridges := make([]string, 0)

	maxConnectedNodes := func(exclusions ...string) int {
		var maxConnectedNodes int
		for _, n := range in {
			connectedNodes := n.ConnectedNodes(exclusions...)
			if connectedNodes > maxConnectedNodes {
				maxConnectedNodes = connectedNodes
			}
		}

		return maxConnectedNodes
	}

	initialMaxConnectedNodes := maxConnectedNodes()

	for _, node := range in {
		for _, connection := range node.connections {
			edge := node.value + connection.value
			connectedNodesWithoutEdge := maxConnectedNodes(edge)

			if initialMaxConnectedNodes != connectedNodesWithoutEdge {
				bridges = append(bridges, edge)
			}
		}
	}

	return bridges
}
