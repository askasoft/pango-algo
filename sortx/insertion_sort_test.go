package sortx

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	f := func(data Interface) {
		InsertionSort(data)
	}
	testSortFunc(t, "InsertionSort", f)
}
