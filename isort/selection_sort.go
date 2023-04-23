package isort

// SelectionSort selection sort implementation
// Time:  O(n^2)
// Space: O(1)
func SelectionSort(data Interface) {
	n := data.Len()

	for i := 0; i < n; i++ {
		m := i
		for j := i + 1; j < n; j++ {
			if data.Less(j, m) {
				m = j
			}
		}
		if m != i {
			data.Swap(m, i)
		}
	}
}
