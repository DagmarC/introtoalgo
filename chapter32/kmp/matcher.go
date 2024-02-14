package kmp

import "fmt"

// KMP - Knuth-Morris-Pratt algorithm
// precomputes pi-prefix array of len m (len of pattern P) rather than whole transition function
func KMPMatcher(t, p string) []int {
	pi := computePrefix(p)
	shifts := make([]int, 0)
	n := len(t)
	m := len(p)
	q := -1                  // number of characters matched, q starts from -1 cuz it is also the index in the pattern being matched, starting from q+1 and we need to cover the first char that starts at q+1=0--p[0], so q need to be -1
	for i := 0; i < n; i++ { // scan the text from left to right
		for q > -1 && p[q+1] != t[i] {
			q = pi[q] // next character does not match
		}
		if p[q+1] == t[i] {
			q++ // next character matches
		}
		if q == m-1 {
			fmt.Printf("pattern occurs with shift %d\n", i-m+1)
			shifts = append(shifts, i-m+1) // when appending q, add +1 since q started from -1 matches
			q = pi[q]                      // look for the next match
		}
	}
	return shifts
}

// computePrefix will count the number of matches against itself and returns the array of it
// so when the incorrect match in chars in KMP matcher occurs it can look to this array if there is sub-match p=abab T=aaabaabb
func computePrefix(p string) []int {
	m := len(p)
	pi := make([]int, m)
	pi[0] = -1
	k := -1
	for q := 1; q < m; q++ {
		for k > -1 && p[k+1] != p[q] {
			k = pi[k]
		}
		if p[k+1] == p[q] {
			k++
		}
		pi[q] = k
	}
	return pi
}
