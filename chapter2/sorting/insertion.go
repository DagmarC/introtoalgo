package sorting

import (
	"errors"
)

// Insertion sort can be compared to having pile of cards in one R hand and inserting them one by one into another L hand.
// During insertion to that L hand you will sort the card to the right place, so it means that all cards in the left card are sorted.

func InsertionSort(a []int) error {
	if len(a) == 0 {
		return errors.New("ERROR: empty array")
	}
	for i := 1; i < len(a); i++ {
		card := a[i] // card  to be inserted at the correct place
		j := i - 1
		for j >= 0 && card < a[j] { // while card is less than a[j], shift positions by 1 in arr A
			a[j+1] = a[j]
			j--
		}
		a[j+1] = card // insert card at correct place: j+1, j is where the condition was not fulfilled
	}
	return nil
}

func InsertionSortR(a *[]int) error {
	if len(*a) == 0 {
		return nil
	}
	insertRec(a, len((*a))-2, len(*a)-1)
	return nil
}

// insertRec function i: represents currently A[0:i] sorted arr, and n represents element A[n] being inserted
func insertRec(a *[]int, i, n int) {
	if n <= 0 {
		return
	}
	insertRec(a, i-1, n-1)

	// Insert key into the already sorted arr a[0:i] via shifting greater numbers until proper position for the key is found
	key := (*a)[n]
	for i >= 0 && (*a)[i] > key {
		(*a)[i+1] = (*a)[i]
		i--
	}
	(*a)[i+1] = key // Insert the key, i+1 bc. at i the for cycle above stops
}
