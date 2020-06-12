package countriesInMap

/*
Time: O(n), given by the fact that sooner or later we will have to visit each country
Space: O(n), given by alreadyVisitedMap
*/
func Solution(A [][]int) int {
	//creating a map signaling already visited elements with A's size
	alreadyVisitedMap := make([][]bool, len(A))
	for i := range alreadyVisitedMap {
		alreadyVisitedMap[i] = make([]bool, len(A[i]))
	}

	differentCountries := 0
	for i := range A {
		for j := range A[i] {
			if !alreadyVisitedMap[i][j] {
				differentCountries++
				visitCountry(i, j, A, alreadyVisitedMap)
			}
		}
	}
	return differentCountries
}

/*
visit region i,j and each region belonging to the same country
*/
func visitCountry(i, j int, A [][]int, alreadyVisitedMap [][]bool) {
	if alreadyVisitedMap[i][j] {
		return
	} //already visited? don't revisit it
	alreadyVisitedMap[i][j] = true // we visited this country "region" now

	/*
		From specification, once we get to a new region we already know that we visited everything north.
		So we only need to visit south, east and west.
	*/

	//can we travel east?
	if j < len(A[i])-1 {
		//is it the same country?
		if A[i][j+1] == A[i][j] {
			visitCountry(i, j+1, A, alreadyVisitedMap)
		}
	}
	//can we travel west?
	if j > 0 {
		//is it the same country?
		if A[i][j-1] == A[i][j] {
			visitCountry(i, j-1, A, alreadyVisitedMap)
		}
	}
	//can we travel south?
	if i < len(A)-1 {
		//is it the same country?
		if A[i+1][j] == A[i][j] {
			visitCountry(i+1, j, A, alreadyVisitedMap)
		}
	}
}
