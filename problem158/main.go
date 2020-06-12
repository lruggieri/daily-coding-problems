package problem158

/*
Time: O(n*m), because in the worst case we are going to check every cell
Space: O(n*m), because of the 'explored' support matrix
*/
func Solution(matrix [][]bool) int {
	//build explored matrix
	explored := make([][]Node, len(matrix))
	for i := range explored {
		explored[i] = make([]Node, len(matrix[i]))
	}

	exploreNode(0, 0, matrix, explored)
	return explored[0][0].paths
}

type Node struct {
	visited bool // set this path as already visited, so we don't recompute the calculation
	valid   bool // if it's possible to reach the end from this node
	paths   int  // number of paths from which is possible to reach the end from this node
}

func exploreNode(i, j int, matrix [][]bool, explored [][]Node) {
	// if outside the matrix
	if i >= len(matrix) || j >= len(matrix[i]) {
		return
	}

	// if we already visited this path
	if explored[i][j].visited {
		return
	}

	//set this node as explored
	explored[i][j].visited = true

	if !matrix[i][j] {
		explored[i][j].valid = false
		return
	}

	// if we are at the bottom right corner
	if i == len(matrix)-1 && j == len(matrix[i])-1 {
		explored[i][j].valid = true
		explored[i][j].paths = 1
		return
	}

	// if we can go right
	if i+1 < len(matrix) {
		exploreNode(i+1, j, matrix, explored)
		if explored[i+1][j].valid {
			explored[i][j].paths += explored[i+1][j].paths
		}
	}
	// if we can go down
	if j < len(matrix[i])-1 {
		exploreNode(i, j+1, matrix, explored)
		if explored[i][j+1].valid {
			explored[i][j].paths += explored[i][j+1].paths
		}
	}

	// if we can reach the end from here, then this path is valid. Otherwise it's not.
	if explored[i][j].paths >= 1 {
		explored[i][j].valid = true
	} else {
		explored[i][j].valid = false
	}
}
