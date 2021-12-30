package esort

// Less less function for loser tree
// Should return a bool:
//    true , if a < b
//    false, if a >= b
type Less func(a, b int) bool

// LoserTree array based loser tree
type LoserTree struct {
	nodes []int
	less  Less
}

// NewLoserTree create a loser tree instance
func NewLoserTree(size int, less Less) *LoserTree {
	return &LoserTree{nodes: make([]int, size), less: less}
}

// Construct construct the loser tree
func (lt *LoserTree) Construct() {
	for i := len(lt.nodes) - 1; i >= 0; i-- {
		lt.nodes[i] = -1
	}
	for i := len(lt.nodes) - 1; i >= 0; i-- {
		lt.adjust(i)
	}
}

// Root get the root node (winner)
func (lt *LoserTree) Root() int {
	return lt.nodes[0]
}

func (lt *LoserTree) adjust(q int) {
	t := (q + len(lt.nodes)) >> 1
	p := lt.nodes[t]

	// nodes[t] = -1 --> minimum
	for t > 0 {
		if p == -1 || (q != -1 && lt.less(p, q)) {
			q, lt.nodes[t] = lt.nodes[t], q
		}
		t = t >> 1
		p = lt.nodes[t]
	}
	lt.nodes[0] = q
}

// Adjust adjust the loser tree from the root
func (lt *LoserTree) Adjust() {
	q := lt.nodes[0]
	t := (q + len(lt.nodes)) >> 1
	p := lt.nodes[t]

	for t > 0 {
		if lt.less(p, q) {
			q, lt.nodes[t] = lt.nodes[t], q
		}
		t = t >> 1
		p = lt.nodes[t]
	}
	lt.nodes[0] = q
}
