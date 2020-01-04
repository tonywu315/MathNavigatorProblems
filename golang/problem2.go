/*
I created 2 maps, which contain the values in sliceA and sliceB. I iterated through
both maps one at a time and added the keys to either same or diff, depending on
whether the keys were in both or just one slice.

Time complexity: O(n + m)
Space complexity: O(n + m)

Examples used to test the code:
split([1,2,3,4], [1,3,5])
split([1,2,3], [])
split([], [1,2,3])
split([],[])
split([1,2,3],[1,2,3])
*/

package golang

// Returns an array containing 2 slices that contain the mutual and different elements
func Split(sliceA, sliceB []int) [2][]int {
	same := make([]int, 0)
	diff := make([]int, 0)

	// The value is irrelevant since we only need to lookup. struct{} is 0 bytes, which saves space
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})

	// Adds each element in sliceA to the map
	for _, num := range sliceA {
		// Sets the key to a struct of size 0
		m1[num] = struct{}{}
	}

	// Repeats for sliceB
	for _, num := range sliceB {
		m2[num] = struct{}{}
	}

	// Loops through all the keys in m1
	for key := range m1 {
		// If the key is in both, add it to same
		if _, ok := m2[key]; ok {
			same = append(same, key)
		} else {	// Otherwise add it to diff
			diff = append(diff, key)
		}
	}

	// If a key is in m2 but not in m1, add it to diff
	for key := range m2 {
		if _, ok := m1[key]; !ok {
			diff = append(diff, key)
		}
	}

	return [2][]int{same, diff}
}