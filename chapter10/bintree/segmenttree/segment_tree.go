package segmenttree

import (
	"math"
)

type SegmentTree struct {
	Array []int
	Tree  []int
}

func (st *SegmentTree) BuildNew() {
	st.Tree = make([]int, 2*len(st.Array)-1)
	st.buildRec(0, 0, len(st.Array)-1)
}

func (st *SegmentTree) buildRec(i, leftCh, rightCh int) {
	if leftCh == rightCh {
		st.Tree[i] = st.Array[leftCh] // leaf node
		return
	}
	mid := (leftCh + rightCh) / 2

	st.buildRec(2*i+1, leftCh, mid)
	st.buildRec(2*i+2, mid+1, rightCh)

	st.Tree[i] = st.Tree[2*i+1] + st.Tree[2*i+2]
}

// Sum of elements from array, where index boundaries are [li...ri]
func (st *SegmentTree) Sum(i, li, ri int) int {
	return st.sumRec(0, 0, len(st.Array)-1, li, ri)
}

func (st *SegmentTree) sumRec(i, leftCh, rightCh, li, ri int) int {
	if li > ri {
		return 0
	}
	if leftCh == li && rightCh == ri {
		return st.Tree[i]
	}
	mid := (leftCh + rightCh) / 2 // split interval to left and right child

	return st.sumRec(2*i+1, leftCh, mid, li, int(math.Min(float64(mid), float64(ri)))) + st.sumRec(2*i+2, mid+1, rightCh, int(math.Max(float64(li), float64(mid+1))), ri)
}

func (st *SegmentTree) Update(pos, newValue int) {
	st.recUpdate(0, 0, len(st.Array)-1, pos, newValue)
}

func (st *SegmentTree) recUpdate(i, leftCh, rightCh, pos, newValue int) {
	if leftCh > rightCh {
		return
	}
	if leftCh == rightCh {
		st.Tree[i] = newValue
		return
	}
	mid := (leftCh + rightCh) / 2
	if pos <= mid {
		st.recUpdate(2*i+1, leftCh, mid, pos, newValue)
	} else {
		st.recUpdate(2*i+2, mid+1, rightCh, pos, newValue)
	}
	st.Tree[i] = st.Tree[2*i+1] + st.Tree[2*i+2]
}
