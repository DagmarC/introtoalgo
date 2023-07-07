package automaton

import (
	"fmt"
	"strings"
)

type Transition struct {
	q  int  // from state
	r  rune // via
	to int  // to state
}

// Finite automaton that uses the transition function to match the pattern P in string T (1st occurrence).
// t is a given string, transFn  is a transition fn from QxAlphabet of M on pattern P, p is the Pattern P,  m is the length of the Pattern P
func FinAutStringMatcher(t string, transFn func(int, rune) int, p string) []int {
	matches := make([]int, 0)
	fmt.Println("----STRING MATCHER PATTERN----")
	m := len(p)
	q := 0 // start state 0, only accepting state is when q == m (lenght of the pattern P)
	for i, r := range t {
		// new q is the new state when T[i] (r) is read, when q = 0, it means that 0 chars of P matches in the string T
		// when q = 2 -> means that first two chars of P matches in T
		// from := q
		q = transFn(q, r)
		// fmt.Printf("T%d: %d->%c=%d\n", i, from, r, q)
		if q == m {
			fmt.Printf("First match of pattern P %s in string T %s found with shift %d\n", p, t, i-m+1)
			matches = append(matches, i-m+1)
		}
	}
	return matches // no match found
}

func ComputeTransFn(p string, alphabet []rune) func(int, rune) int {
	fmt.Println("----COMPUTE TRANSITION FUNCTION----")
	transitions := make([]Transition, 0)
	m := len(p)
	// q is a state from 0 to m
	for q := 0; q <= m; q++ {
		for _, a := range alphabet {
			k := min(m, q+1)
			// fmt.Println(p[:q]+string(a), p[:k])
			for !strings.HasSuffix(p[:q]+string(a), p[:k]) {
				k--
			}
			transitions = append(transitions, Transition{q: q, r: a, to: k})
		}
	}
	transFn := func(state int, r rune) int {
		for _, t := range transitions {
			if t.q == state && t.r == r {
				return t.to
			}
		}
		return 0 // state or r not found
	}
	return transFn
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
