package isort

// ShellSort shell sort implementation
// Time:  O(n*logn)
// Space: O(1)
func ShellSort(data Interface) {
	n := data.Len()

	// Start with a big gap, then reduce the gap
	for gap := n >> 1; gap > 0; gap >>= 1 {
		// Do a gapped insertion sort for this gap size.
		// The first gap elements a[0..gap-1] are already in gapped order
		// keep adding one more element until the entire array is gap sorted
		for i := gap; i < n; i++ {
			// shift earlier gap-sorted elements up until the correct
			// location for a[i] is found
			for j := i; j >= gap && data.Less(j, j-gap); j -= gap {
				data.Swap(j, j-gap)
			}
		}
	}
}

// ShellKunthSort  kunth shell sort implementation
// Time:  O(n^1.5)
// Space: O(1)
func ShellKunthSort(data Interface) {
	n := data.Len()

	// calculate gap
	// O(n^1.5)) by Knuth,1973>: 1, 4, 13, 40, 121, ...
	gap := 1
	for gap < n/3 {
		gap = gap*3 + 1
	}

	for ; gap > 0; gap /= 3 {
		for i := gap; i < n; i++ {
			for j := i; j >= gap && data.Less(j, j-gap); j -= gap {
				data.Swap(j, j-gap)
			}
		}
	}
}
