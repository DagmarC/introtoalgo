package quicksort

import (
	"fmt"
)

// Quicksort sorts arr of size n in situ (in place) with indices p, r, where p < r && both p, r belongs to the  range 0 - (n-1) or (1-n)
// To sort the whole arr start with arr, 0, n-1
func Quicksort(arr *[]int, p, r int) error {
	if p < 0 || r >= len(*arr) {
		return fmt.Errorf("invalid indices of arr %v, p=%d, r=%d, expected in range 0-%d", arr, p, r, len(*arr))
	}
	if p < r {
		q := Partition(arr, p, r) // q is a index of a PIVOT => all elems from p to q-1 are smaller than PIVOT (elem on arr[q]) && all elems from q+1 to r are greater than the PIVOT
		Quicksort(arr, p, q-1)
		Quicksort(arr, q+1, r)
	}
	return nil
}

// Partition will firstly choose the PIVOT (last elem-arr[r]) and then divide the arr[p:r] to 3 parts according to it:
// 1st part: elems LESS than PIVOT arr[q] => arr[p:q-1], index i determines last elem less than pivot
// 2nd part: elem PIVOT on arr[q]
// 3rd part: elems GREATER than PIVOT arr[q] => arr[q+1:r], index j determines last elem greater than pivot
// i is p-1 at the beginning, cuz no element is sorted yet, so no elem is determined to be smaller than the pivot, yet
// all elems from j to r are not sorted, yet (starts at j==p && ends when j==r -> all sorted)
// at the end the SWITCH of the first elem greater than the PIVOT and the PIVOT on the last index arr[r] occurs. Note: 1st greater elem is determined by i+1, since i is the index of last elem smaller than the PIVOT
// returns i+1 (index of the newly switched PIVOT)
// precondition: indices p, r are valid indices of arr A && p < r && r is at most n-1, where n is the len of arr A
func Partition(arr *[]int, p, r int) int {
	pivot := (*arr)[r]
	// fmt.Println("====Partition====")
	// fmt.Println("pivot ", pivot)

	i := p - 1 // i starts before the arr, all elems from arr[p] to arr[i] are smaller than the pivot

	for j := p; j < r; j++ {
		if (*arr)[j] < pivot {
			i++                                         // increase the i and switch tha smaller elem a[j] to a[i]
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i] // SWITCH
		}
	}
	(*arr)[i+1], (*arr)[r] = (*arr)[r], (*arr)[i+1] // SWITCH PIVOT to the "middle", worst case i+1 == r
	// fmt.Printf("arr %?v\n======End======\n", arr)
	return i + 1
}
