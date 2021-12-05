package sortx

// HeapSort heap sort implementation
// Time:  O(n*logn)
// Space: O(1)
func HeapSort(data Interface) {
	n := data.Len()

	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(data, i, n)
	}

	// One by one extract an element from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		data.Swap(0, i)

		// call max heapify on the reduced heap
		heapify(data, 0, i)
	}
}

func heapify(data Interface, i, n int) {
	root := i       // Initialize largest as root
	node := 2*i + 1 // left child = 2*i + 1

	for node < n {
		// If left child is larger than root
		if data.Less(root, node) {
			root = node
		}

		node++
		// If right child is larger than largest so far
		if node < n && data.Less(root, node) {
			root = node
		}

		// If largest is root
		if root == i {
			return
		}

		data.Swap(i, root)

		// Recursively heapify the affected sub-tree
		i = root
		node = 2*i + 1
	}
}
