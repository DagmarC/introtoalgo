package quicksort

import (
	"testing"

	"github.com/DagmarC/introtoalgo/utils"
)

func TestQuicksort(t *testing.T) {
	tc := []struct {
		arr      []int
		expected []int
	}{
		{arr: []int{1, 4, 13, 15, 23, 3, 5, 9}, expected: []int{1, 3, 4, 5, 9, 13, 15, 23}},
		{arr: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{arr: []int{3, 2}, expected: []int{2, 3}},
		{arr: []int{0}, expected: []int{0}},
		{arr: []int{}, expected: []int{}},
	}
	for _, test := range tc {
		err := Quicksort(&test.arr, 0, len(test.arr)-1)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !utils.EqualIntArr(test.arr, test.expected) {
			t.Errorf("error, got %v, expected %v", test.arr, test.expected)
		}
	}
}
