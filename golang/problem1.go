/*
I looped through each index in the string. At each index, I checked if stringB
matched stringA starting from that index. If it did, I added it to the indices slice.

Time complexity: O(nm)
Space complexity: O(n)

Examples used to test the code:
findSubstring("googlygoggles", "googly")
findSubstring("googlygoggles", "oo")
findSubstring("googlygoggles", "g")
findSubstring("mmmmm", "mm")
findSubstring("asdf", "")
findSubstring(None, None)
findSubstring("", "hello")
findSubstring("", "")
findSubstring("aaaaaa","a")
findSubstring("baaaaa","a")
findSubstring("aaaaab","a")
*/

package golang

// Returns a slice containing the indices that have matches
func FindSubstring(stringA, stringB string) (indices []int) {
	// If stringB is empty, return an empty slice
	if len(stringB) == 0 {
		return
	}

	// Iterates through the indexes up to n - m + 1
	for i := 0; i < len(stringA) - len(stringB) + 1; i++ {
		flag := true

		// For each index, check if each place afterwards is the same
		for j := 0; j < len(stringB); j++ {
			// If its not, stop checking and continue to the next index of stringA
			if stringA[i + j] != stringB[j] {
				flag = false
				break
			}
		}

		// If the substring matches, add the index to indices
		if flag {
			indices = append(indices, i)
		}
	}

	// Returns indices
	return
}