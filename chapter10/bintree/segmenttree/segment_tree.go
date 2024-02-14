package segmenttree

import (
	"math"
)

type SegmentTree struct {
	Array       []int
	SegmentTree []int
}

func (st *SegmentTree) BuildNew() {
	st.SegmentTree = make([]int, 2*len(st.Array)-1)
	st.buildRec(0, 0, len(st.Array)-1)
}

func (st *SegmentTree) buildRec(i, leftCh, rightCh int) {
	if leftCh == rightCh {
		st.SegmentTree[i] = st.Array[leftCh] // leaf node
		return
	}
	mid := (leftCh + rightCh) / 2

	st.buildRec(2*i+1, leftCh, mid)
	st.buildRec(2*i+2, mid+1, rightCh)

	st.SegmentTree[i] = st.SegmentTree[2*i+1] + st.SegmentTree[2*i+2]
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
		return st.SegmentTree[i]
	}
	mid := (leftCh + rightCh) / 2 // split interval to left and right child

	return st.sumRec(2*i+1, leftCh, mid, li, int(math.Min(float64(mid), float64(ri)))) + st.sumRec(2*i+2, mid+1, rightCh, int(math.Max(float64(li), float64(mid+1))), ri)
}
