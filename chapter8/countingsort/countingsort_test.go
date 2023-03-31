package countingsort

import (
	"testing"

	"github.com/DagmarC/introtoalgo/utils"
)

func TestCountingsort(t *testing.T) {
	tc := []struct {
		arr      []int
		expected []int
		k        int // max elem in arr
	}{
		{arr: []int{1, 4, 13, 15, 23, 3, 5, 9}, expected: []int{1, 3, 4, 5, 9, 13, 15, 23}, k: 23},
		{arr: []int{1, 2, 3, 2, 4, 5}, expected: []int{1, 2, 2, 3, 4, 5}, k: 5},
		{arr: []int{0, 3, 0, 10, 3, 2, 0, 0, 2, 6, 8}, expected: []int{0, 0, 0, 0, 2, 2, 3, 3, 6, 8, 10}, k: 10},
		{arr: []int{0, 0}, expected: []int{0, 0}, k: 0},
		{arr: []int{}, expected: []int{}, k: 0},
	}
	for _, test := range tc {
		res := Countingsort(test.arr, len(test.arr), test.k)
		if !utils.EqualIntArr(test.expected, res) {
			t.Errorf("error, got %v, expected %v", res, test.expected)
		}
	}
}
