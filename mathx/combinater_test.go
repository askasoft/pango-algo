package mathx

import (
	"fmt"
	"sort"
	"testing"
)

func TestCombinate5(t *testing.T) {
	testCombinate(t, 5)
}

func TestCombinate6(t *testing.T) {
	testCombinate(t, 6)
}

func testCombinate(t *testing.T, n int) {
	cnt := 0

	hash := map[string]bool{}
	iter := func(a []int) bool {
		cnt++
		fmt.Println(cnt, ": ", a)
		if !sort.IntsAreSorted(a) {
			t.Error("Not sorted: ", a)
		}
		s := fmt.Sprint(a)
		if _, ok := hash[s]; ok {
			t.Error("Duplicated: ", a)
		}
		hash[s] = true
		return true
	}

	for m := 1; m <= n; m++ {
		fmt.Println("------ n: ", n, " ---------- m: ", m)

		cnt = 0
		exp := CombinationCount(n, m)

		hash = map[string]bool{}
		Combinate(n, m, iter)
		if exp != cnt {
			t.Errorf("Combinate count - expect: %d, actual: %d", exp, cnt)
		}
	}
}

func TestCombinate51(t *testing.T) {
	testCombinate1(t, 5)
}

func TestCombinate61(t *testing.T) {
	testCombinate1(t, 6)
}

func testCombinate1(t *testing.T, n int) {
	cnt := 0

	iter := func(a []int) bool {
		cnt++
		fmt.Println(cnt, ": ", a)
		if !sort.IntsAreSorted(a) {
			t.Error("Not sorted: ", a)
		}
		return false
	}

	for m := 1; m <= n; m++ {
		fmt.Println("------ n: ", n, " ---------- m: ", m)

		cnt = 0
		exp := 1

		Combinate(n, m, iter)
		if exp != cnt {
			t.Errorf("Combinate count - expect: %d, actual: %d", exp, cnt)
		}
	}
}
