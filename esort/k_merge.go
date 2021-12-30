package esort

import "io"

// KmergeIO K-Merge IO interface
type KmergeIO interface {
	Less(a, b int) bool
	Read(i int) error
	Write(i int) error
}

// Kmerge a K-Merge struct
type Kmerge struct {
	lotr *LoserTree // loser tree
	eofs []bool     // End-Of-File slice
	kmio KmergeIO   // IO interface
}

// NewKmerge create a K-Merge instance
func NewKmerge(k int, kmio KmergeIO) *Kmerge {
	km := &Kmerge{eofs: make([]bool, k), kmio: kmio}

	km.lotr = NewLoserTree(k, func(a, b int) bool {
		if km.eofs[a] {
			return false
		}
		if km.eofs[b] {
			return true
		}
		return km.kmio.Less(a, b)
	})
	return km
}

func (km *Kmerge) read(i int) error {
	err := km.kmio.Read(i)
	if err == io.EOF {
		km.eofs[i] = true
	} else if err != nil {
		return err
	}
	return nil
}

// Merge do k-merge
func (km *Kmerge) Merge() error {
	// init read
	for i := 0; i < len(km.eofs); i++ {
		if err := km.read(i); err != nil {
			return err
		}
	}

	// construct the loser tree
	km.lotr.Construct()

	// do k-merge
	for !km.eofs[km.lotr.Root()] {
		if err := km.kmio.Write(km.lotr.Root()); err != nil {
			return err
		}

		if err := km.read(km.lotr.Root()); err != nil {
			return err
		}

		km.lotr.Adjust()
	}

	return nil
}
