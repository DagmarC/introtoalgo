package main

import (
	"fmt"
	"log"

	"github.com/DagmarC/introtoalgo/chapter2/sorting/sorts"
)

func main() {
	// Insertion sort
	a := []int{100, 1, 7, 22, 11, 4, 102, 9}
	err := sorts.InsertionSort(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INFO: Insertion sort result: ", a)

	// Insertion sort recursive
	a = []int{100, 11, 7, 22, 11, 4, 102, 9}
	err = sorts.InsertionSortR(&a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INFO: Insertion sort REC result: ", a)

	// Selection sort
	a = []int{5, 1, 7, 3, 11, 4, 2, 9}
	err = sorts.SelectionSort(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INFO: Selection sort result: ", a)

	// Merge sort
	b := []int{5, 1, 7, 3, 11, 4, 2, 9}
	err = sorts.MergeSort(b, 0, len(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INFO: Merge sort result: ", b)
}
