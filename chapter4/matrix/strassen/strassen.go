package strassen

import (
	"fmt"
	"log"
)

const LEAF_SIZE = 1 // IF 1 => pure Strassen algorithm is used.

type Matrix [][]int

// NewMatrix creates 2D array of size nxn
func NewMatrix(n int) Matrix {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	return m
}

func (m Matrix) String() string {
	s := ""
	for i := 0; i < len(m); i++ {
		s += fmt.Sprintln(m[i])
	}
	return s
}

// add two matrices C=A+B
func (m Matrix) add(b Matrix) Matrix {
	c := NewMatrix(len(m))
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			c[i][j] = m[i][j] + b[i][j]

		}
	}
	return c
}

// substract two matrices C=A-B
func (m Matrix) substract(b Matrix) Matrix {
	c := NewMatrix(len(m))
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			c[i][j] = m[i][j] - b[i][j]
		}
	}
	return c
}

// combineFinal will fill the matrix m with matrix cp only in rango of specified indice which are is, ie, js, je
func (m Matrix) combineFinal(is, ie, js, je int, cp Matrix) {
	var k, p int
	for i := is; i < ie; i++ {
		for j := js; j < je; j++ {

			// Check for correct indices range
			if i >= len(m) || j >= len(m[i]) {
				fmt.Println("ERROR: bad range indices is, ie, js, je are greater than the specified matrix m.", i, j)
				return
			}
			// Check for correct indices range
			if k >= len(cp) || p >= len(cp[k]) {
				fmt.Println("ERROR: bad range indices k, p are greater than the specified matrix cp.", i, j)
				return
			}

			m[i][j] = cp[k][p]
			p++
		}
		k++
		p = 0
	}
}

// fillParts will fillParts the matrix m with matrix cp, indices in matrix cp are is, ie, js, je, In comparison to combine matrix m is the smaller matrix and cp is the larger one where the boundaries are indices
func (m Matrix) fillParts(is, ie, js, je int, cp Matrix) {
	for i, k := is, 0; i < ie; i, k = i+1, k+1 {
		for j, p := js, 0; j < je; j, p = j+1, p+1 {

			// Check for correct indices range
			if i >= len(cp) || j >= len(cp[i]) {
				fmt.Println("ERROR: bad range indices is, ie, js, je are greater than the given matrix cp.", i, j)
				return
			}
			// Check for correct indices range
			if k >= len(m) || p >= len(m[k]) {
				fmt.Println("ERROR: bad range indices k, p are greater than the specified matrix m.", i, j)
				return
			}

			m[k][p] = cp[i][j]
		}
	}

}

// ======================STRASSEN======================
// ====================================================

// prereq: matrices A, B are divisible by 2ˆn, if not then fill with 0s (nearest power 2ˆx - n) 
func StrassenMultip(a, b Matrix) Matrix {
	// Are matrices divisible by 2ˆn? if not fill it with 0s
	a = matrix2n(a)
	b = matrix2n(b)
	c := strassenRec(a, b)
	return c
}

func matrix2n(origin Matrix) Matrix {
	if IsPower(len(origin)) {
		return origin
	}
	np, err := NearestPower(len(origin))
	if err != nil {
		log.Fatal(err)
	}
	az := NewMatrix(np)
	az.fillParts(0, len(origin), 0, len(origin), origin)

	return az
}

func strassenRec(a, b Matrix) Matrix {
	n := len(a)
	c := NewMatrix(n)

	if n <= LEAF_SIZE {
		return BruteForce(a, b)
	}
	// Divide large matrices A, B into 4 sub-parts and init new matrices
	a11 := NewMatrix(n / 2)
	a12 := NewMatrix(n / 2)
	a21 := NewMatrix(n / 2)
	a22 := NewMatrix(n / 2)

	b11 := NewMatrix(n / 2)
	b12 := NewMatrix(n / 2)
	b21 := NewMatrix(n / 2)
	b22 := NewMatrix(n / 2)

	// Fill in the newly divided matrices with numbers from the original A, B matrices
	a11.fillParts(0, n/2, 0, n/2, a)
	a12.fillParts(0, n/2, n/2, n, a)
	a21.fillParts(n/2, n, 0, n/2, a)
	a22.fillParts(n/2, n, n/2, n, a)

	b11.fillParts(0, n/2, 0, n/2, b)
	b12.fillParts(0, n/2, n/2, n, b)
	b21.fillParts(n/2, n, 0, n/2, b)
	b22.fillParts(n/2, n, n/2, n, b)

	// Calculate the sum matrices (addition and subtraction, sum is not that complicated as the product is)- Strassen's method
	s1 := b12.substract(b22)
	s2 := a11.add(a12)
	s3 := a21.add(a22)
	s4 := b21.substract(b11)
	s5 := a11.add(a22)
	s6 := b11.add(b22)
	s7 := a12.substract(a22)
	s8 := b21.add(b22)
	s9 := a11.substract(a21)
	s10 := b11.add(b12)

	// Calculate thge product matrices - (multiplication - Notice that there are only 7 - better than the brute forces 8)
	p1 := strassenRec(a11, s1)
	p2 := strassenRec(s2, b22)
	p3 := strassenRec(s3, b11)
	p4 := strassenRec(a22, s4)
	p5 := strassenRec(s5, s6)
	p6 := strassenRec(s7, s8)
	p7 := strassenRec(s9, s10)

	// Calculate the final product sub matrices
	c11 := p4.add(p5).add(p6).substract(p2)
	c12 := p1.add(p2)
	c21 := p3.add(p4)
	c22 := p1.add(p5).substract(p3).substract(p7)

	// Combine C parts together
	c.combineFinal(0, n/2, 0, n/2, c11)
	c.combineFinal(0, n/2, n/2, n, c12)
	c.combineFinal(n/2, n, 0, n/2, c21)
	c.combineFinal(n/2, n, n/2, n, c22)

	return c
}

func BruteForce(a, b Matrix) Matrix {
	n := len(a)

	// init C
	c := NewMatrix(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}
