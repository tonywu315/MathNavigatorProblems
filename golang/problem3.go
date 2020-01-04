/*
I created a dictionary to store the number of votes each candidate got.
To figure out which candidate won, I interated through the dictionary
and replaced the variables "most" and "winner" with the candidates that had the
most votes.

Time complexity: O(n)
Space complexity: O(n)

Examples used to test the code:
printCandidateInfo(["A","B","B","B","C","C","D"])
printCandidateInfo(["A"])
*/

package golang

import "fmt"

// Prints each candidate's votes, the winner, and the additional votes needed for each candidate to win
func PrintCandidateInfo(votes []string) {
	counts := make(map[string]int)

	// Stores the votes each candidate received in the map
	for _, candidate := range votes {
		counts[candidate] += 1
	}

	var most int
	var winner string

	for candidate := range counts {
		// Prints the number of votes each candidate received and addresses if vote is plural
		if counts[candidate] == 1 {
			fmt.Println("Candidate", candidate, "received", counts[candidate], "vote")
		} else {
			fmt.Println("Candidate", candidate, "received", counts[candidate], "votes")
		}

		// If one of the candidates has more votes, replace most and winner with that one
		if counts[candidate] > most {
			most = counts[candidate]
			winner = candidate
		}
	}

	// Prints the candidate with the most votes and addresses if vote is plural
	if most == 1 {
		fmt.Println("\nCandidate", winner, "won with", most, "vote\n")
	} else {
		fmt.Println("\nCandidate", winner, "won with", most, "votes\n")
	}

	for candidate := range counts {
		// Other candidates need 1 more than the candidate with the most votes to win
		difference := most - counts[candidate] + 1

		// The winner does not need 1 extra vote
		if difference == 1 {
			difference = 0
		}

		// Prints the additional votes needed to win and addresses if vote is plural
		if difference == 1 {
			fmt.Println("Candidate", candidate, "needed", difference, "more vote to win")
		} else {
			fmt.Println("Candidate", candidate, "needed", difference, "more votes to win")
		}
	}
}