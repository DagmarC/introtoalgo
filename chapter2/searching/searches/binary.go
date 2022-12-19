package searches

import "fmt"

func BinarySearch(a []int, key int) bool {
	if len(a) == 0 {
		fmt.Println("INFO: key not found, empty array")
		return false
	}
	res := binaryRec(a, key, 0, len(a))
	return res
}

// binaryRec will return true if the key is present in the array in O(logn) complexity.
// Given array a of lenght n, where start and n belongs to [0:n] interval.
func binaryRec(a []int, key, start, end int) bool {
	if start > end-1 {
		return false
	}
	if start == end-1 {
		return a[start] == key
	}
	mid := (start + end) / 2
	if a[mid] == key {
		return true
	}
	if a[mid] > key {
		// go left
		return binaryRec(a, key, start, mid)
	} else {
		// go right
		return binaryRec(a, key, mid+1, end)
	}
}
