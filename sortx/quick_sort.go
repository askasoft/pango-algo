package sortx

// QuickSort quick sort implementation
// Time:  O(n*logn)
// Space: O(logn)
func QuickSort(data Interface) {
	qsort(data, 0, data.Len()-1)
}

func qsort(data Interface, lo, hi int) {
	if lo >= hi {
		return
	}

	pivot := partition(data, lo, hi)
	qsort(data, lo, pivot-1)
	qsort(data, pivot+1, hi)
}

func partition(data Interface, lo, hi int) int {
	pivot := lo
	index := lo + 1

	for i := lo + 1; i <= hi; i++ {
		if data.Less(i, pivot) {
			if i != index {
				data.Swap(i, index)
			}
			index++
		}
	}

	index--
	if pivot != index {
		data.Swap(pivot, index)
	}
	return index
}
