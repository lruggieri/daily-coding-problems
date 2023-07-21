package problem1052

/*
enemies = {
    0: [3],
    1: [2],
    2: [1, 4],
    3: [0, 4, 5],
    4: [2, 3],
    5: [3]
}


	0	1	2	3	4	5

0	0	0	0	1	0	0

1	0	0	1	0	0	0

2	0	1	0	0	1	0

3	1	0	0	0	1	1

4	0	0	1	1	0	0

5	0	0	0	1	0	0
*/

// Solution : O(S + E) where S is the number of students and E is the number of enemy relationships
func Solution(playersEnemies map[int][]int) (teamA, teamB []int, solvable bool) {
	teams := map[int]bool{} // false: teamA, true: teamB

	for player1 := range playersEnemies {
		if _, ok := teams[player1]; ok {
			continue
		}

		teams[player1] = false

		studentsToCheck := []int{player1}

		for len(studentsToCheck) > 0 {
			currentStudent := studentsToCheck[0]
			studentsToCheck = studentsToCheck[1:]

			currentStudentEnemies := playersEnemies[currentStudent]

			for _, enemy := range currentStudentEnemies {
				if enemyTeam, ok := teams[enemy]; !ok {
					studentsToCheck = append(studentsToCheck, enemy)
					teams[enemy] = !teams[currentStudent]
				} else if enemyTeam == teams[currentStudent] {
					return nil, nil, false
				}
			}
		}
	}

	for player, team := range teams {
		if !team {
			teamA = append(teamA, player)
		} else {
			teamB = append(teamB, player)
		}
	}

	return teamA, teamB, true
}
