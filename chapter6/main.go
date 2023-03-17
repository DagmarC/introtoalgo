package main

import (
	"fmt"

	"github.com/DagmarC/introtoalgo/chapter6/heap"
)

func main() {
	arr := []int{5, 3, 12, 1, 4, 66}
	heap.Heapsort(arr)
	fmt.Println("Sorted arr:", arr)
}
