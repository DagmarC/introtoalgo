package indexedtree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		arr      []int
		want []int
	}{
		{
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: []int{-100, 1, 3, 3, 10, 5, 11, 7, 36},
		},
	}
	for _, tc := range tests {
		ft := NewFenwickTree(tc.arr)
		if !cmp.Equal(ft.bitTree, tc.want) {
			t.Errorf("got %v, want %v", ft.bitTree, tc.want)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		ft   FenwickTree
		i    int
		want int
	}{
		{
			ft:   FenwickTree{size: 8, arr: []int{1, 2, 3, 4, 5, 6, 7, 8}, bitTree: []int{-100, 1, 3, 3, 10, 5, 11, 7, 36}},
			i:    3,
			want: 6,
		},
		{
			ft:   FenwickTree{size: 8, arr: []int{1, 2, 3, 4, 5, 6, 7, 8}, bitTree: []int{-100, 1, 3, 3, 10, 5, 11, 7, 36}},
			i:    4,
			want: 10,
		},
		{
			ft:   FenwickTree{size: 8, arr: []int{1, 2, 3, 4, 5, 6, 7, 8}, bitTree: []int{-100, 1, 3, 3, 10, 5, 11, 7, 36}},
			i:    8,
			want: 36,
		},
		{
			ft:   FenwickTree{size: 8, arr: []int{1, 2, 3, 4, 5, 6, 7, 8}, bitTree: []int{-100, 1, 3, 3, 10, 5, 11, 7, 36}},
			i:    5,
			want: 15,
		},
	}
	for _, tc := range tests {
		if got := tc.ft.Sum(tc.i); got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		ft   FenwickTree
		i    int
		inc  int
		want []int
	}{
		{
			ft:   FenwickTree{size: 8, arr: []int{1, 2, 3, 4, 5, 6, 7, 8}, bitTree: []int{-100, 1, 3, 3, 10, 5, 11, 7, 36}},
			i:    3,
			inc:  3,
			want: []int{-100, 1, 3, 6, 13, 5, 11, 7, 39},
		},
		{
			ft:   FenwickTree{size: 8, arr: []int{1, 2, 3, 4, 5, 6, 7, 8}, bitTree: []int{-100, 1, 3, 3, 10, 5, 11, 7, 36}},
			i:    1,
			inc:  10,
			want: []int{-100, 11, 13, 3, 20, 5, 11, 7, 46},
		},
		{
			ft:   FenwickTree{size: 8, arr: []int{1, 2, 3, 4, 5, 6, 7, 8}, bitTree: []int{-100, 1, 3, 3, 10, 5, 11, 7, 36}},
			i:    8,
			inc:  4,
			want: []int{-100, 1, 3, 3, 10, 5, 11, 7, 40},
		},
		{
			ft:   FenwickTree{size: 8, arr: []int{1, 2, 3, 4, 5, 6, 7, 8}, bitTree: []int{-100, 1, 3, 3, 10, 5, 11, 7, 36}},
			i:    7,
			inc:  3,
			want: []int{-100, 1, 3, 3, 10, 5, 11, 10, 39},
		},
		{
			ft:   FenwickTree{size: 8, arr: []int{1, 2, 3, 4, 5, 6, 7, 8}, bitTree: []int{-100, 1, 3, 3, 10, 5, 11, 7, 36}},
			i:    5,
			inc:  55,
			want: []int{-100, 1, 3, 3, 10, 60, 66, 7, 91},
		},
	}
	for _, tc := range tests {
		tc.ft.Add(tc.i, tc.inc)
		if !cmp.Equal(tc.ft.bitTree, tc.want) {
			t.Errorf("got %v, want %v", tc.ft.bitTree, tc.want)
		}
	}
}
