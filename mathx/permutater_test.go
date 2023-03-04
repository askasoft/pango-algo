package mathx

import (
	"fmt"
	"testing"
)

func TestPermutate5(t *testing.T) {
	testPermutate(t, 5)
}

func TestPermutate6(t *testing.T) {
	testPermutate(t, 6)
}

func testPermutate(t *testing.T, n int) {
	cnt := 0

	hash := map[string]bool{}
	iter := func(a []int) bool {
		cnt++
		fmt.Println(cnt, ": ", a)
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
		exp := PermutationCount(n, m)

		hash = map[string]bool{}
		Permutate(n, m, iter)
		if exp != cnt {
			t.Errorf("Permutate count - expect: %d, actual: %d", exp, cnt)
		}
	}
}

func TestPermutate51(t *testing.T) {
	testPermutate1(t, 5)
}

func TestPermutate61(t *testing.T) {
	testPermutate1(t, 6)
}

func testPermutate1(t *testing.T, n int) {
	cnt := 0

	iter := func(a []int) bool {
		cnt++
		fmt.Println(cnt, ": ", a)
		return false
	}

	for m := 1; m <= n; m++ {
		fmt.Println("------ n: ", n, " ---------- m: ", m)

		cnt = 0
		exp := 1

		Permutate(n, m, iter)
		if exp != cnt {
			t.Errorf("Permutate count - expect: %d, actual: %d", exp, cnt)
		}
	}
}
