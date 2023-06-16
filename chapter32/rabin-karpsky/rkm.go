package rabinkarpsky

import (
	"fmt"
	"math"
	"unicode/utf8"
)

// RabinKarpskyMatcher uses modulo q to calculate the substring one by one, t is text, p is pattern, q is a prime and d is usually taken from |Sigma alphabet|
// modulo/reminder prooblem: Go impl differs to Py impl: https://github.com/golang/go/issues/448
func RabinKarpskyMatcher(t string, p string, d int, q int) []int {
	matches := make([]int, 0) // sapce for at least 4 matches, attaches shifts (indices/ints) in text t where match occurs

	n := utf8.RuneCountInString(t)
	m := utf8.RuneCountInString(p)
	h := math.Pow(float64(d), float64(m-1))

	ph := 0  // 'hash'of a pattern substring
	ts := 0  // 'hash' of a current substring in text t with shift s
	hmc := 0 // count of all hash matches of ph and ts

	fmt.Printf("BEGIN----Text %s, Pattern %s\n", t, p)
	fmt.Printf("BEGIN----h=%.0f, n=%d, m=%d d=%d q=%d\n\n", h, n, m, d, q)

	// preprocesisng
	for i, pi := range p {
		ti, _ := utf8.DecodeRuneInString(t[i:]) // returns first rune in string from substring t[i:]
		// We need to truncate integer division toward negative infinity
		// 1. way:  ((m % n) + n) % n
		// 2. way: temp = m % n; if temp < 0: temp += n
		if ph = (d*ph + int(pi)) % q; ph < 0 {
			ph += q
		}
		if ts = (d*ts + int(ti)) % q; ts < 0 {
			ts += q
		}
	}

	for s := -1; s < n-m; s++ {
		fmt.Println("SUBSTRING TEXT: ", t[s+1:s+m+1])
		if ph == ts {
			fmt.Printf("HASH MATCH: p=%s, substring=%s, s+1=%d\n", p, t[s+1:s+m+1], s+1)
			hmc++
			// check if it is a valid modulos match or just collision - spurious hit
			if p == t[s+1:s+m+1] {
				fmt.Println("MATCH OK", s+1)
				matches = append(matches, s+1)
			}
		}
		if s < n-m {
			// calculate next position hash in text, delete first digit and add the next one
			to, _ := utf8.DecodeRuneInString(t[s+1:])   // old rune to be removed from hash
			tn, _ := utf8.DecodeRuneInString(t[s+m+1:]) // next rune to be added to a hash

			trem := int(to) * int(h) // digit to be substracted
			// Modulo/ Reminder hack
			if ts = ((d * (ts - trem)) + int(tn)) % q; ts < 0 {
				ts += q
			}
		}
	}
	fmt.Println("N. OF SPURIOUS HITS (collisions that are not actually match): ", hmc-len(matches))
	fmt.Println()
	return matches
}
