package segmenttree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		segmentTree SegmentTree
		expected    []int
	}{
		{
			segmentTree: SegmentTree{Array: []int{1, 2, 3, 4}},
			expected:    []int{10, 3, 7, 1, 2, 3, 4},
		},
		{
			segmentTree: SegmentTree{Array: []int{1, 2}},
			expected:    []int{3, 1, 2},
		},
	}

	for _, tc := range tests {
		tc.segmentTree.BuildNew()
		if !cmp.Equal(tc.segmentTree.Tree, tc.expected) {
			t.Errorf("got %v, expected %v", tc.segmentTree.Tree, tc.expected)
		}
	}
}

func TestSumTree(t *testing.T) {
	tests := []struct {
		segmentTree SegmentTree
		leftIndex   int
		rightIndex  int
		expected    int
	}{
		{
			segmentTree: SegmentTree{Array: []int{1, 2, 3, 4}},
			leftIndex:   0,
			rightIndex:  1,
			expected:    3,
		},
		{
			segmentTree: SegmentTree{Array: []int{1, 2}},
			leftIndex:   1,
			rightIndex:  1,
			expected:    2,
		},
		{
			segmentTree: SegmentTree{Array: []int{22, 2, -1, 4}},
			leftIndex:   1,
			rightIndex:  3,
			expected:    5,
		},
		{
			segmentTree: SegmentTree{Array: []int{1, 2, 3, 4}},
			leftIndex:   0,
			rightIndex:  2,
			expected:    6,
		},
	}

	for _, tc := range tests {
		tc.segmentTree.BuildNew()
		if got := tc.segmentTree.Sum(0, tc.leftIndex, tc.rightIndex); got != tc.expected {
			t.Errorf("got %v, expected %v", got, tc.expected)
		}
	}
}

func TestUpdateTree(t *testing.T) {
	tests := []struct {
		segmentTree SegmentTree
		index       int
		value       int
		expected    []int
	}{
		{
			segmentTree: SegmentTree{Array: []int{1, 2, 3, 4}, Tree: []int{10, 3, 7, 1, 2, 3, 4}},
			index:       1,
			value:       10,
			expected:    []int{18, 11, 7, 1, 10, 3, 4},
		},
		{
			segmentTree: SegmentTree{Array: []int{8, 3, 10, 4}, Tree: []int{25, 11, 14, 8, 3, 10, 4}},
			index:       2,
			value:       2,
			expected:    []int{17, 11, 6, 8, 3, 2, 4},
		},
	}
	for _, tc := range tests {
		tc.segmentTree.Update(tc.index, tc.value)
		if !cmp.Equal(tc.segmentTree.Tree, tc.expected) {
			t.Errorf("got %v, expected %v", tc.segmentTree.Tree, tc.expected)
		}
	}
}
