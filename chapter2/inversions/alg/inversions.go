package alg

//Inversions

// Let A[1 \ldots n]A[1…n] be an array of nn distinct numbers. If i < ji<j and A[i] > A[j]A[i]>A[j],
// then the pair (i, j)(i,j) is called an inversion of AA.

// Give an algorithm that determines the number of inversions in any permutation on nn elements in
// \Theta(n \lg n)Θ(nlgn) worst-case time. (Hint: Modify merge sort.)

// NumInversions is in fact the merge sort that counts when th inversion is done
func NumInversions(a []int, p, r int) int {
	if p > r || len(a) < r || p < 0 {
		return 0
	}
	if p >= r-1 {
		return 0 // base case - 1 element in arr A
	}
	q := (p + r) / 2 // median of arr A

	res := NumInversions(a, p, q) + NumInversions(a, q, r)
	return res + inversions(a, p, q, r)
}

func inversions(a []int, p, q, r int) int {

	nl := q - p // 1, 5, 3, 3, 5, 6, 7 .. 3, 5, 7 -> 5-3 = 2  [3, 5]
	nr := r - q // 7-5 = 2 [6, 7]

	// Create L and R arrs
	left := make([]int, nl)
	for i := 0; i < nl; i++ {
		left[i] = a[p+i]
	}
	// Create L and R arrs
	right := make([]int, nr)
	for i := 0; i < nr; i++ {
		right[i] = a[q+i]
	}
	// fmt.Printf("\nMERGE a=%v, p=%d, q=%d, r=%d, left=%v, right=%v\n", a, p, q, r, left, right)

	var i int   // indicates the smallest element left in arr L
	var j int   // indicates the smalles eleemtn left in arr R
	var inv int // indicates number of inversions done

	// k indicates the position of already sorted elements in arr A that starts from p and ends on r
	k := p
	for ; k < r; k++ {
		if i >= nl || j >= nr {
			break
		}
		// L arr has smaller element on top
		if left[i] <= right[j] {
			a[k] = left[i] // Add to arr a from L arr
			i++
		} else {
			a[k] = right[j]      // Add to arr a from R arr
			inv += len(left[i:]) // it applies for the rest of left (SORTED) arr
			j++
		}
	}
	// Rest of L arr
	for i < nl {
		a[k] = left[i] // Add to arr a from L arr
		i++
		k++
	}
	// Rest of R arr
	for j < nr {
		a[k] = right[j] // Add to arr a from L arr
		j++
		k++
	}
	return inv
}
