package main

import (
	"fmt"

	"github.com/DagmarC/introtoalgo/chapter6/heapsort"
)

func main() {
	arr := []int{5, 3, 12, 1, 4, 66}
	heapsort.Heapsort(arr)
	fmt.Println("Sorted arr:", arr)
}