package sorting

import (
	"errors"
)

// 1. Divide - Given arr A[p..r], Divide arr boundaries to q=(p+r)/2 to create arrays L[p, q], R[q+1, r].
// 2. Conquer - By sorting two arrays R and L using recursively the merge procedure. Base case is the arr of 1 element - already sorted one.
// 3. Merge - L and R array merge and sort together. Prereq: Both L and R are already sorted.

// MergeSort function will recursively calculates median from part of array A until one element arr from array A[p:r+1] has left and then merge it together.
// indices p and r, r = n
func MergeSort(a []int, p, r int) error {
	if p > r || len(a) < r || p < 0 {
		return errors.New("ERROR: invalid boundaries of arr A")
	}
	if p >= r-1 {
		return nil // base case - 1 element in arr A
	}
	q := (p + r) / 2   // median of arr A
	MergeSort(a, p, q) // a[q] is not present here in merge procedure
	MergeSort(a, q, r) // a[q] is present here in merge procedure

	merge(a, p, q, r)
	// fmt.Println("AFTER MERGE", a)
	return nil
}

func merge(a []int, p, q, r int) {

	nl := q - p // 1, 5, 3, 3, 5, 6, 7 .. 3, 5, 7 -> 5-3 = 2  [3, 5]
	nr := r - q // 7-5 = 2 [6, 7]

	// Create L and R arrs
	left := make([]int, nl)
	for i := 0; i < nl; i++ {
		left[i] = a[p+i]
	}
	// Create L and R arrs
	right := make([]int, nr)
	for i := 0; i < nr; i++ {
		right[i] = a[q+i]
	}
	// fmt.Printf("\nMERGE a=%v, p=%d, q=%d, r=%d, left=%v, right=%v\n", a, p, q, r, left, right)

	var i int // indicates the smallest element left in arr L
	var j int // indicates the smalles eleemtn left in arr R

	// k indicates the position of already sorted elements in arr A that starts from p and ends on r
	k := p
	for ; k < r; k++ {

		if i >= nl || j >= nr {
			break
		}
		// L arr has smaller element on top
		if left[i] <= right[j] {
			a[k] = left[i] // Add to arr a from L arr
			i++
		} else {
			a[k] = right[j] // Add to arr a from R arr
			j++
		}
	}
	// Rest of L arr
	for i < nl {
		a[k] = left[i] // Add to arr a from L arr
		i++
		k++
	}
	// Rest of R arr
	for j < nr {
		a[k] = right[j] // Add to arr a from L arr
		j++
		k++
	}
}
