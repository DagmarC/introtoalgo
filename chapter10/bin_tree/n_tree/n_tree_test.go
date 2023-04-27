package ntree

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type MyInt int

func (mi MyInt) String() string {
	return strconv.Itoa(int(mi))
}

func TestNTreePrint(t *testing.T) {
	n1 := NewNnode[MyInt](1)
	n2 := NewNnode[MyInt](2)
	n3 := NewNnode[MyInt](3)
	n4 := NewNnode[MyInt](4)
	n5 := NewNnode[MyInt](5)
	n6 := NewNnode[MyInt](6)
	n7 := NewNnode[MyInt](7)
	n8 := NewNnode[MyInt](8)
	n9 := NewNnode[MyInt](9)
	n10 := NewNnode[MyInt](10)
	n11 := NewNnode[MyInt](11)

	nt := NewNTree[MyInt]()

	nt.root = n1
	n1.leftCh = n2

	n2.rightSib = n3
	n2.leftCh = n7

	n3.leftCh = n4
	n4.rightSib = n5
	n5.rightSib = n6
	n5.leftCh = n11

	n7.rightSib = n8
	n7.leftCh = n9
	n9.rightSib = n10

	want := "N-ary tree: 1 2 7 9 10 8 3 4 5 11 6 "
	got := nt.String()

	if !cmp.Equal(want, got) {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestNTreePrint2(t *testing.T) {
	n1 := NewNnode[MyInt](1)
	n2 := NewNnode[MyInt](2)
	n3 := NewNnode[MyInt](3)
	n4 := NewNnode[MyInt](4)
	n5 := NewNnode[MyInt](5)
	n6 := NewNnode[MyInt](6)
	n7 := NewNnode[MyInt](7)
	n8 := NewNnode[MyInt](8)
	n9 := NewNnode[MyInt](9)
	n10 := NewNnode[MyInt](10)
	n11 := NewNnode[MyInt](11)

	nt := NewNTree[MyInt]()

	nt.root = n1
	n1.leftCh = n2

	n2.rightSib = n3
	n3.rightSib = n4
	n3.leftCh = n5
	n4.rightSib = n6

	n5.rightSib = n11

	n6.leftCh = n8
	n8.rightSib = n9
	n8.leftCh = n10

	n10.leftCh = n7

	want := "N-ary tree: 1 2 3 5 11 4 6 8 10 7 9 "
	got := nt.String()

	if !cmp.Equal(want, got) {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestNTreePrintEmpty(t *testing.T) {
	nt := NewNTree[MyInt]()

	want := "Empty N-ary tree"
	got := nt.String()

	if !cmp.Equal(want, got) {
		t.Errorf("want %s, got %s", want, got)
	}
}
