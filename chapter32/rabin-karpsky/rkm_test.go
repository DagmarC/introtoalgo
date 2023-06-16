package rabinkarpsky

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRKM(t *testing.T) {
	cases := []struct {
		text    string
		pattern string
		q       int
		d       int
		want    []int
	}{
		{text: "abbabababbba", pattern: "ab", want: []int{0, 3, 5, 7}, q: 13, d: 10},
		{text: "leopardleo", pattern: "par", want: []int{3}, q: 7, d: 26},
		{text: "3141592653589793", pattern: "26", want: []int{6}, q: 11, d: 10},
		{text: "2359023141526739921", pattern: "31415", want: []int{6}, q: 13, d: 10},

	}
	for _, tc := range cases {
		if got := RabinKarpskyMatcher(tc.text, tc.pattern, tc.d, tc.q); !cmp.Equal(got, tc.want) {
			t.Errorf("want %v, got %v...text %s, pattern %s", got, tc.want, tc.text, tc.pattern)
		}
	}
}
