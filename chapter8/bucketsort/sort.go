package bucketsort

import (
	"fmt"

	ll "github.com/DagmarC/introtoalgo/chapter10/linkedlist"
)

// Bucket sort precondition is that all elements are from interval <0.0 - 1)
// n is the len arr
func Bucketsort(arr []float64, n int) {
	b := make([]ll.Linkedlist[float64], n) // each position is linked list
	for _, el := range arr {
		j := int(float64(n) * el)
		b[j].InsertLast(el) // n*A[i], 0 <= (A[i]==el) < 1
	}
	for i := 0; i < n; i++ {
		// b[i] -> sort the linked list elements via insertion sort
		b[i].Sort()
		// b[i].Display()
	}
	fmt.Println("==================CONNECT LLS==================")
	ll.ConnectateLls(b...)
	//Result: Connectate all the B[0], B[1], B[n] lists together in order
	// Return CONNECTATED LIST
}
