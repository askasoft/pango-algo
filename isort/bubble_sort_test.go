package isort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	f := func(data Interface) {
		BubbleSort(data)
	}
	testSortFunc(t, "BubbleSort", f)
}

func testBubbleSortAsc(t *testing.T, m string, a1, a2 []int, lc, sc int) {
	intdbg := &IntSliceDebug{Array: a1}
	BubbleSort(intdbg)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("%s - NE: %v, want %v", m, a1, a2)
	}

	if intdbg.LessCount != lc {
		t.Errorf("%s: Wrong Less Count %d, want %d", m, intdbg.LessCount, lc)
	}

	if intdbg.SwapCount != sc {
		t.Errorf("%s: Wrong Swap Count %d, want %d", m, intdbg.SwapCount, sc)
	}
}

func TestBubbleSortAsc0(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := []int{1, 2, 3, 4, 5}

	testBubbleSortAsc(t, "BubbleSortAsc", a1, a2, len(a1)-1, 0)
}

func TestBubbleSortAsc1(t *testing.T) {
	a1 := []int{2, 1, 3, 4, 5}
	a2 := []int{1, 2, 3, 4, 5}

	testBubbleSortAsc(t, "BubbleSortAsc1", a1, a2, len(a1)-1+len(a1)-2, 1)
}
