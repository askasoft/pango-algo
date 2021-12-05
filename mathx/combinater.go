package mathx

// CombinationCount C(n, m)
func CombinationCount(n, m int) (count int) {
	if n < 1 || m < 1 || m > n {
		return
	}

	count = 1
	k := n - m
	if k < m {
		k = m
	}
	for i := n; i > k; i-- {
		count = count * i / (n - i + 1)
	}

	return
}

// Combinate iterate C(n, m) with function iter
func Combinate(n, m int, iter func([]int) bool) {
	if n < 1 || m < 1 || m > n {
		return
	}

	a := make([]int, m)
	combinate(n, m, a, iter)
}

func combinate(n, m int, a []int, iter func([]int) bool) bool {
	for i := n; i >= m; i-- {
		a[m-1] = i
		if m > 1 {
			if !combinate(i-1, m-1, a, iter) {
				return false
			}
		} else {
			if !iter(a) {
				return false
			}
		}
	}
	return true
}
