package sorts

import "errors"

// Consider sorting nn numbers stored in array AA by first finding the smallest element of AA and exchanging it
// with the element in A[1]A[1]. Then find the second smallest element of AA, and exchange it with A[2]A[2].
//Continue in this manner for the first n-1n−1 elements of AA. Write pseudocode for this algorithm,
// which is known as selection sort. What loop invariant does this algorithm maintain? Why does it need to run
// for only the first n-1n−1 elements, rather than for all nn elements? Give the best-case and
// worst-case running times of selection sort in \ThetaΘ-notation.

func SelectionSort(a []int) error {
	if len(a) == 0 {
		return errors.New("no elements in arr")
	}
	// elements in arr a[0, n-1] -> a[:i] are already sorted
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] && j != min {
				min = j // Set min index to j
			}
		}
		if min != i {
			a[i], a[min] = a[min], a[i] // swap elements
		}
	}
	return nil
}
