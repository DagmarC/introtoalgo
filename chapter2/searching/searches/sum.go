package searches

import (
	"fmt"

	"github.com/DagmarC/introtoalgo/chapter2/sorting/sorts"
)

// SumSearch finds whether in arr A there exists a. b int such that a+b=key in O(nlgn) complexity.
func SumSearch(arr []int, key int) bool {
	// Sort whole arr a via merge sort (nlgn)
	sorts.MergeSort(arr, 0, len(arr)) 

	// For each a search b such thah key-a=b. O(nlgn) - n times searching for b, searching b is lgn
	for _, a := range arr {
		if BinarySearch(arr, key-a) {
			fmt.Printf("a=%d + b=%d = key=%d\n", a, key-a, key)
			return true
		}
	}
	return false
}
