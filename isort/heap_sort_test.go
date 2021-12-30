package isort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	f := func(data Interface) {
		HeapSort(data)
	}
	testSortFunc(t, "HeapSort", f)
}
