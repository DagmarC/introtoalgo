package segmenttree

import "fmt"

type SegmentTree struct {
	Array []int
	Tree  []int
}

// This code snippet is implementing the functionality to build a segment tree data structure. Here's a
// breakdown of what the code is doing:
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

// This `Query` function in the `SegmentTree` struct is used to find the sum of elements within a given
// range `[li...ri]` in the original array. It calls the `queryRec` function recursively to traverse
// the segment tree and calculate the sum based on the specified range.
func (st *SegmentTree) Query(li, ri int) int {
	return st.queryRec(0, 0, len(st.Array)-1, li, ri)
}

func (st *SegmentTree) queryRec(i, leftCh, rightCh, li, ri int) int {
	if li > ri {
		return 0
	}
	if leftCh == li && rightCh == ri {
		return st.Tree[i]
	}
	mid := (leftCh + rightCh) / 2
	fmt.Println(mid, i, leftCh, rightCh, li, ri)
	return min(st.queryRec(2*i+1, leftCh, mid, li, min(mid, ri)), st.queryRec(2*i+2, mid+1, rightCh, max(li, mid+1), ri))
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

	return st.sumRec(2*i+1, leftCh, mid, li, min(mid, ri)) + st.sumRec(2*i+2, mid+1, rightCh, max(li, mid+1), ri)
}

// This `Update` function in the `SegmentTree` struct is used to update the value of a specific element
// in the original array and then propagate the changes to update the corresponding nodes in the
// segment tree.
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
