package indexedtree

import "fmt"

// Fenwick Tree or Binary Indexed Tree: https://www.youtube.com/watch?v=uSFzHCZ4E-8
type FenwickTree struct {
	size    int   // n: length of the array
	arr     []int // array: the input array on which queries are made
	bitTree []int // bit: store the sum of range
}

// The computation of   $g(i)$  is defined using the following simple operation:
// we replace all trailing  $1$  bits in the binary representation of  $i$  with  $0$  bits.
func NewFenwickTree(arr []int) *FenwickTree {
	ft := FenwickTree{
		size:    len(arr) + 1,
		arr:     arr,
		bitTree: make([]int, len(arr)+1),
	}
	ft.bitTree = append([]int{-100}, arr...) // copy original elements from arr to bitTree array, ignore 1st element

	for i := 1; i < ft.size; i++ {
		ft.addDirectParent(i, ft.bitTree[i])
	}
	return &ft
}

// Sum of elements from array in range [0, r]
// indexed from 1, arr [1,2,3,4] i=3 -> elements 1, 2, 3
// 7 	= 00111
// -7	= 11001 /2s complement => switch all bits in 7 and add 1
// 00111 & 11001 = 00001
// 00111 - 00001 = 00110 "child"
func (ft *FenwickTree) Sum(i int) int {
	sum := 0
	for i > 0 {
		sum += ft.bitTree[i]
		i -= (i & -i) // flip the last set bit
	}
	return sum
}

// Add the value in the array on index i with increment inc
// 7 	= 00111
// -7	= 11001 /2s complement => switch all bits in 7 and add 1
// 00111 & 11001 = 00001
// 00111 + 00001 = 01000 PARENT IDX
func (ft *FenwickTree) Add(i, inc int) {
	fmt.Printf("ADD: i=%d\tinc=%d\t\tft.size=%d\n", i, inc, ft.size)
	for i <= ft.size {
		ft.bitTree[i] += inc
		i += (i & -i) // add last set bit
	}
}

func (ft *FenwickTree) addDirectParent(i, inc int) {
	parent := i + (i & -i)
	if parent < ft.size {
		ft.bitTree[parent] = ft.bitTree[parent] + ft.bitTree[i]
	}
}
