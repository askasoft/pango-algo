package isort

// BubbleSort bubble sort implementation
// Time:  O(n^2)
// Space: O(1)
func BubbleSort(data Interface) {
	n := data.Len()

	for i := n - 1; i > 0; i-- {
		swap := false
		for j := 0; j < i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}
