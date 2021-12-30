package isort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

// IntSliceDebug attaches the methods of Interface to []int, sorting in increasing order.
type IntSliceDebug struct {
	Array     []int
	LessCount int
	SwapCount int
	Verbose   bool
}

func (x *IntSliceDebug) Len() int { return len(x.Array) }
func (x *IntSliceDebug) Less(i, j int) bool {
	x.LessCount++
	return x.Array[i] < x.Array[j]
}

func (x *IntSliceDebug) Swap(i, j int) {
	x.SwapCount++
	x.Array[i], x.Array[j] = x.Array[j], x.Array[i]
	if x.Verbose {
		fmt.Println(i, "<->", j, ": ", x.Array)
	}
}

func testSortFunc(t *testing.T, m string, f func(Interface)) {
	for i := 1; i < 100; i++ {
		a := make([]int, i)
		for j := 0; j < i; j++ {
			a[j] = rand.Intn(i)
		}

		a2 := make([]int, i)
		copy(a2, a)
		sort.Ints(a2)

		f(IntSlice(a))
		if !sort.IntsAreSorted(a) {
			t.Fatalf("%s [%d] - NG: %v", m, i, a)
		}

		if !reflect.DeepEqual(a, a2) {
			t.Fatalf("%s [%d] - NE: %v, want %v", m, i, a, a2)
		}
	}
}
