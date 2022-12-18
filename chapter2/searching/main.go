package main

import "github.com/DagmarC/introtoalgo/chapter2/searching/searches"

func main() {
	a := []int{1, 2, 5, 8, 10, 12, 100}
	searches.BinarySearch(a, 7)
	searches.BinarySearch(a, 12)

	b := []int{11}
	searches.BinarySearch(b, 12)
	searches.BinarySearch(b, 11)




}
