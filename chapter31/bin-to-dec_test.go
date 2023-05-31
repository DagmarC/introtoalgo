package chapter31

import (
	"testing"
)

func TestBinToDec(t *testing.T) {
	tc := []struct {
		b    string
		want int
	}{
		{b: "110110", want: 54},
		{b: "0001", want: 1},
		{b: "0010", want: 2},
		{b: "0100", want: 4},
		{b: "1000", want: 8},
		{b: "1010", want: 10},
	}
	for _, ts := range tc {
		if got := BinToDec(ts.b); got != ts.want {
			t.Errorf("bin: %s, got %d, want %d", ts.b, got, ts.want)
		}
	}
}
