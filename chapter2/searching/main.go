package main

import "github.com/DagmarC/introtoalgo/chapter2/searching/searches"

func main() {
	a := []int{1, 2, 5, 8, 10, 12, 100}
	searches.BinarySort(a, 7)
	searches.BinarySort(a, 12)

	b := []int{11}
	searches.BinarySort(b, 12)
	searches.BinarySort(b, 11)




}
