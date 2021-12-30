package isort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	f := func(data Interface) {
		SelectionSort(data)
	}
	testSortFunc(t, "SelectionSort", f)
}
