package main

import (
	"fmt"
	"log"

	"github.com/DagmarC/introtoalgo/chapter2/sorting"
)

func main() {
	// Selection sort
	a := []int{5, 1, 7, 3, 11, 4, 2, 9}
	err := sorting.SelectionSort(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INFO: Selection sort result: ", a)

	// Merge sort
	b := []int{5, 1, 7, 3, 11, 4, 2, 9}
	err = sorting.MergeSort(b, 0, len(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INFO: Merge sort result: ", b)
}
