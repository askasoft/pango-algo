package mathx

// PermutationCount P(n, m)
func PermutationCount(n, m int) (count int) {
	if n < 1 || m < 1 || m > n {
		return
	}

	count = 1
	k := n - m
	for i := n; i > k; i-- {
		count *= i
	}

	return
}

// Permutate iterate P(n, m) with function iter
func Permutate(n, m int, iter func([]int) bool) {
	sw := swapper{iter: iter}
	Combinate(n, m, func(a []int) bool {
		sw.nums = a
		return sw.swap(0, m)
	})
}

type swapper struct {
	nums []int
	iter func([]int) bool
}

func (sw *swapper) swap(m, n int) bool {
	if m >= n-1 {
		return sw.iter(sw.nums)
	}

	if !sw.swap(m+1, n) {
		return false
	}

	for i := m + 1; i < n; i++ {
		sw.nums[m], sw.nums[i] = sw.nums[i], sw.nums[m]

		if !sw.swap(m+1, n) {
			return false
		}

		sw.nums[m], sw.nums[i] = sw.nums[i], sw.nums[m]
	}

	return true
}
