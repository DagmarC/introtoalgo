package main

import (
	"fmt"

	"github.com/DagmarC/introtoalgo/chapter2/searching/searches"
)

func main() {
	// Binary Search
	fmt.Println("\n====================Binary Search=========================")
	fmt.Println("==========================================================")


	a := []int{1, 2, 5, 8, 10, 12, 100}
	key := 7
	fmt.Printf("Is key %d present in arr a=%v? \t%v\n", key, a, searches.BinarySearch(a, key))
	key = 12
	fmt.Printf("Is key %d present in arr a=%v? %v\n", key, a, searches.BinarySearch(a, key))
	b := []int{11}
	fmt.Printf("Is key %d present in arr a=%v? \t\t%v\n", key, b, searches.BinarySearch(b, key))

	// Sum Search
	fmt.Println("\n=====================Sum Search========================")
	fmt.Println("=======================================================")

	a = []int{22, 11, 3, 4, 7, 8}
	key = 26
	fmt.Printf("For key=%d, there are a+b = key in arr a=%v \t%v\n\n", key, a, searches.SumSearch(a, key))
	a = []int{22, 11, 3, 4, 7, 8}
	key = 20
	fmt.Printf("For key=%d, there are a+b = key in arr a=%v \t%v\n\n", key, a, searches.SumSearch(a, key))

}
