package isort

// InsertionSort insertion sort implementation
// Time:  O(n^2)
// Space: O(1)
func InsertionSort(data Interface) {
	n := data.Len()

	for i := 0; i < n; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
