package main

import "fmt"

/*
Range game

You are at the range. You have N targets labeled 1 to N from left to right.
Hit the target by performing the following operation:
• In each round starting from the first element and moving from left to right, hit every alternate target.
• Repeat the above step until only one target is left.
Note: You will get i points when you hit a target labeled i.

Calculate the score you will obtain by continuously hitting targets until there is only one target remaining.

Function description
Complete the Solve() function. This function takes the following argument and returns the total number of points you will get by hitting targets until only one target is left:
• N: Represents the total number of targets

Input format
• The first line contains an integer N denoting the total number of targets.

Output format
Print the total points you will get by hitting the targets by following the specified operation.

Constraints
1 ≤ N≤ 10€
*/

func solve(n int) int {
	if n <= 1 {
		return 0
	}

	sum := 0

	hits := 0
	targets := make([]bool, n)

	for {
		nextIsHit := true

		for i := 1; i <= n; i++ {
			if hits == n-1 {
				return sum
			}

			if !targets[i-1] {
				if nextIsHit {
					sum += i
					hits++
					targets[i-1] = true
				}

				nextIsHit = !nextIsHit
			}
		}
	}
}

func main() {
	var out int = solve(5)
	fmt.Print(out)
}
