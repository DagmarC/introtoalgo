package factorial

import "testing"

func TestFactorial(t *testing.T) {
	tc := []struct {
		in       int
		expected int
	}{
		{in: 1, expected: 1},
		{in: 2, expected: 2},
		{in: 5, expected: 120},
		{in: 4, expected: 24},
	}
	for _, test := range tc {
		if out := factorial(test.in); out != test.expected {
			t.Errorf("factorial(%v) got %d expected %d\n", test.in, out, test.expected)
		}
	}
}
