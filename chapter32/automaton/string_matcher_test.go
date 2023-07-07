package automaton

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSTRM(t *testing.T) {
	tests := []struct {
		alphabet []rune
		pattern  string
		text     string
		matches  []int
	}{
		{
			alphabet: []rune{'a', 'b'},
			pattern:  "aabab",
			text:     "aaababaabaababaab",
			matches:  []int{1, 9},
		},
		{
			alphabet: []rune{'a', 'b', 'c'},
			pattern:  "ababaca",
			text:     "abababacabac",
			matches:  []int{2},
		},
		{
			alphabet: []rune{'a', 'r', 'i', 's', 'k', 'o', 'g'},
			pattern:  "risko",
			text:     "arisisriskorisgriskoriskisk",
			matches:  []int{6, 15},
		},
	}
	for _, tc := range tests {
		transFn := ComputeTransFn(tc.pattern, tc.alphabet)
		got := FinAutStringMatcher(tc.text, transFn, tc.pattern)
		if !cmp.Equal(tc.matches, got) {
			t.Errorf("want %v, got %v\n", tc.matches, got)
		}
		fmt.Printf("PASS: text=%s && 1st match %s T[%d:%d]\n", tc.text, tc.text[got[0]:got[0]+len(tc.pattern)], got[0], got[0]+len(tc.pattern))
	}
}
