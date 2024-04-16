package segmenttree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		segmentTree SegmentTree
		want        []int
	}{
		{
			segmentTree: SegmentTree{Array: []int{1, 2, 3, 4}},
			want:        []int{10, 3, 7, 1, 2, 3, 4},
		},
		{
			segmentTree: SegmentTree{Array: []int{1, 2}},
			want:        []int{3, 1, 2},
		},
	}

	for _, tc := range tests {
		tc.segmentTree.BuildNew()
		if !cmp.Equal(tc.segmentTree.Tree, tc.want) {
			t.Errorf("got %v, want %v", tc.segmentTree.Tree, tc.want)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		segmentTree SegmentTree
		leftIndex   int
		rightIndex  int
		want        int
	}{
		{
			segmentTree: SegmentTree{Array: []int{1, 2, 3, 4}},
			leftIndex:   0,
			rightIndex:  1,
			want:        3,
		},
		{
			segmentTree: SegmentTree{Array: []int{1, 2}},
			leftIndex:   1,
			rightIndex:  1,
			want:        2,
		},
		{
			segmentTree: SegmentTree{Array: []int{22, 2, -1, 4}},
			leftIndex:   1,
			rightIndex:  3,
			want:        5,
		},
		{
			segmentTree: SegmentTree{Array: []int{1, 2, 3, 4}},
			leftIndex:   0,
			rightIndex:  2,
			want:        6,
		},
	}

	for _, tc := range tests {
		tc.segmentTree.BuildNew()
		if got := tc.segmentTree.Sum(0, tc.leftIndex, tc.rightIndex); got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		segmentTree SegmentTree
		index       int
		value       int
		want        []int
	}{
		{
			segmentTree: SegmentTree{Array: []int{1, 2, 3, 4}, Tree: []int{10, 3, 7, 1, 2, 3, 4}},
			index:       1,
			value:       10,
			want:        []int{18, 11, 7, 1, 10, 3, 4},
		},
		{
			segmentTree: SegmentTree{Array: []int{8, 3, 10, 4}, Tree: []int{25, 11, 14, 8, 3, 10, 4}},
			index:       2,
			value:       2,
			want:        []int{17, 11, 6, 8, 3, 2, 4},
		},
	}
	for _, tc := range tests {
		tc.segmentTree.Update(tc.index, tc.value)
		if !cmp.Equal(tc.segmentTree.Tree, tc.want) {
			t.Errorf("got %v, want %v", tc.segmentTree.Tree, tc.want)
		}
	}
}

func TestQuery(t *testing.T) {
	tests := []struct {
		segmentTree SegmentTree
		start       int
		end         int
		want        int
	}{
		{
			segmentTree: SegmentTree{Array: []int{1, 2, 3, 4}, Tree: []int{10, 3, 7, 1, 2, 3, 4}},
			start:       0,
			end:         1,
			want:        9,
		},
	}
	for _, tc := range tests {
		tc.segmentTree.Query(tc.start, tc.end)
		if got := tc.segmentTree.Query(tc.start, tc.end); got != tc.want {
			t.Errorf("got %v, want %v on tree %v", got, tc.want, tc.segmentTree)
		}
	}
}
