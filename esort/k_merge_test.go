package esort

import (
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

type testKmergeIO struct {
	nums [][]int
	ins  []int
	outs []int
	exps []int
}

func newTestKmergeIORandom(k int) *testKmergeIO {
	tk := &testKmergeIO{}

	tk.nums = make([][]int, k)
	for i := 0; i < k; i++ {
		ns := make([]int, rand.Intn(100))
		for j := 0; j < len(ns); j++ {
			ns[j] = rand.Intn(100)
		}
		tk.nums[i] = ns
	}

	tk.init()

	return tk
}

func (tk *testKmergeIO) init() {
	k := len(tk.nums)

	tk.ins = make([]int, k)

	// make expects
	for i := 0; i < k; i++ {
		ns := tk.nums[i]
		sort.Ints(ns)
		tk.exps = append(tk.exps, ns...)
	}
	sort.Ints(tk.exps)
}

func (tk *testKmergeIO) Less(a, b int) bool {
	return tk.ins[a] < tk.ins[b]
}

func (tk *testKmergeIO) Read(i int) error {
	ns := tk.nums[i]
	if len(ns) == 0 {
		return io.EOF
	}
	tk.ins[i] = ns[0]
	tk.nums[i] = ns[1:]
	return nil
}

func (tk *testKmergeIO) Write(i int) error {
	tk.outs = append(tk.outs, tk.ins[i])
	return nil
}

func TestKMergeRandom(t *testing.T) {
	for k := 2; k < 100; k++ {
		tk := newTestKmergeIORandom(k)
		in := fmt.Sprint(tk.nums)

		km := NewKmerge(k, tk)
		err := km.Merge()
		if err != nil {
			t.Errorf("%d %v", k, err)
			continue
		}

		if !sort.IntsAreSorted(tk.outs) {
			t.Errorf("%d Not sorted\n Inputs: %s\nOutputs: %v", k, in, tk.outs)
		}

		if !reflect.DeepEqual(tk.exps, tk.outs) {
			t.Errorf("%d Not equal\nExpects: %v\nOutputs: %v", k, tk.exps, tk.outs)
		}
	}
}

func TestKMergeSimple(t *testing.T) {
	cs := [][][]int{
		{{1, 2}},
		{{}, {}},
		{{1}, {2}},
		{{1}, {2, 3}},
		{{1, 2, 3}, {4, 5}, {2, 3}},
	}

	for i, c := range cs {
		tk := &testKmergeIO{}
		tk.nums = c
		tk.init()

		in := fmt.Sprint(tk.ins)

		km := NewKmerge(len(c), tk)
		err := km.Merge()
		if err != nil {
			t.Errorf("%d %v", i, err)
			continue
		}

		if !sort.IntsAreSorted(tk.outs) {
			t.Errorf("%d Not sorted\n Inputs: %s\nOutputs: %v", i, in, tk.outs)
		}

		if !reflect.DeepEqual(tk.exps, tk.outs) {
			t.Errorf("%d Not equal\nExpects: %v\nOutputs: %v", i, tk.exps, tk.outs)
		}
	}
}
