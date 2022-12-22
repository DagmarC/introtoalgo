package main

import (
	"fmt"

	"github.com/DagmarC/introtoalgo/chapter2/inversions/alg"
)

func main() {
	a := []int{2, 3, 8, 6, 1}
	fmt.Println("Number of inversions in arr a: ", a, alg.NumInversions(a, 0, len(a)))
	a = []int{5, 4, 3, 2, 1}
	fmt.Println("Number of inversions in arr a: ", a, alg.NumInversions(a, 0, len(a)))
	a = []int{2, 1}
	fmt.Println("Number of inversions in arr a: ", a, alg.NumInversions(a, 0, len(a)))
	a = []int{1, 2}
	fmt.Println("Number of inversions in arr a: ", a, alg.NumInversions(a, 0, len(a)))
}