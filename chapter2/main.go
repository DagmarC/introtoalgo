package main

import (
	"fmt"
	"log"

	"github.com/DagmarC/introtoalgo/chapter2/sorting"
)

func main() {
	a := []int{5, 1, 7, 3, 11, 4, 2, 9}
	err := sorting.MergeSort(a, 0, len(a))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INFO: Merge sort result: ", a)
}
