package main

import (
	"fmt"

	"github.com/DagmarC/introtoalgo/chapter4/matrix/strassen"
)

func main() {
	a := strassen.Matrix([][]int{
		{1, 3},
		{7, 5},
	})
	b := strassen.Matrix([][]int{
		{6, 8},
		{4, 2},
	})

	fmt.Println("==========A========")
	fmt.Println(a)
	fmt.Println("==========B========")
	fmt.Println(b)

	fmt.Println("==========BRUTE FORCE========")
	fmt.Println(strassen.BruteForce(a, b))

	fmt.Println("==========STRASSEN V.========")
	fmt.Println(strassen.StrassenMultip(a, b))

	a = strassen.Matrix([][]int{
		{11, 3, 2, 1},
		{7, 5, 0, 4},
		{7, 1, 2, 4},
		{9, 5, 5, 2},
	})
	b = strassen.Matrix([][]int{
		{2, 2, 3, 1},
		{7, 5, 2, 4},
		{3, 1, 0, 4},
		{5, 0, 1, 6},
	})

	fmt.Println("==========A========")
	fmt.Println(a)
	fmt.Println("==========B========")
	fmt.Println(b)

	fmt.Println("==========BRUTE FORCE========")
	fmt.Println(strassen.BruteForce(a, b))

	fmt.Println("==========STRASSEN V.========")
	fmt.Println(strassen.StrassenMultip(a, b))

	a = strassen.Matrix([][]int{
		{11, 3, 2, 1, 2},
		{7, 5, 0, 4, 3},
		{7, 1, 2, 4, 4},
		{9, 5, 5, 2, 1},
		{5, 0, 1, 6, 1},
	})
	b = strassen.Matrix([][]int{
		{2, 2, 3, 1, 3},
		{7, 5, 2, 4, 0},
		{3, 1, 0, 4, 1},
		{5, 0, 1, 6, 1},
		{5, 0, 1, 6, 1},

	})

	fmt.Println("==========A========")
	fmt.Println(a)
	fmt.Println("==========B========")
	fmt.Println(b)

	fmt.Println("==========BRUTE FORCE========")
	fmt.Println(strassen.BruteForce(a, b))

	fmt.Println("==========STRASSEN V.========")
	fmt.Println(strassen.StrassenMultip(a, b))
}
