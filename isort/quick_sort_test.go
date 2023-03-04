package isort

import (
	"fmt"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	f := func(data Interface) {
		QuickSort(data)
	}
	testSortFunc(t, "QuickSort", f)
}

func TestQuickSortDbg(t *testing.T) {
	a := []int{4, 3, 5, 7, 1, 8, 6, 5}

	fmt.Println("0 <-> 0 : ", a)
	QuickSort(&IntSliceDebug{Array: a, Verbose: true})

	if !sort.IntsAreSorted(a) {
		t.Fatalf("%s [%d] - NG: %v", "QuickSort", 0, a)
	}
}
