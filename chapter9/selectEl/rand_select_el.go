package selectEl

import (
	"github.com/DagmarC/introtoalgo/chapter7/quicksort"
)

// RandomizedSelect worst case O(nˆ2) but in general it is much better, depends on the pivot choosing,
// p and r are the boundaries of the array (please pass at the beginning <0, n-1>, not <1-n>)
// i stands for the i-th smallest number in array, i belongs to the interval <1, n> (1st smallest is i=1)
func RandomizedSelect(arr *[]int, p, r, i int) int {
	if p > r {
		return -1
	}
	if i > len(*arr) || i < 0 {
		return -1
	}
	if p == r {
		return (*arr)[p] // 1 <= i <= r-p+1 means that i=1
	}
	q := quicksort.Partition(arr, p, r) // Use partition (or better randomized partition) to divide the arr according the pivot
	k := q - p + 1                      // number of elements in the low side of partition plus 1 for the Pivot element

	if i == k {
		return (*arr)[q] // the pivot is the answer
	} else if i < k {
		return RandomizedSelect(arr, p, q-1, i)
	} else {
		return RandomizedSelect(arr, q+1, r, i-k)
	}
}
