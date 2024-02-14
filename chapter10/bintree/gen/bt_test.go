package gen

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBinTree(t *testing.T) {
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)
	n5 := NewNode(5)
	n6 := NewNode(6)
	n7 := NewNode(7)

	bt := NewBinTree()
	bt.root = n6
	bt.root.left = n5
	bt.root.right = n4
	n5.right = n3
	n4.left = n2
	n2.right = n1
	n1.right = n7

	want := "Binary Tree: 5 3 6 2 1 7 4 "
	got := bt.String()

	if !cmp.Equal(got, want) {
		t.Errorf("want %s, got %s", got, want)
	}
	err := bt.printBTstack()
	if err != nil {
		fmt.Println(err)
	}
}

func TestBinTreeEmpty(t *testing.T) {
	bt := NewBinTree()

	want := "Binary Tree: Empty\n"
	got := bt.String()

	if !cmp.Equal(got, want) {
		t.Errorf("want %s, got %s", got, want)
	}
	err := bt.printBTstack()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func TestBinTreeRight(t *testing.T) {
	n3 := NewNode(3)
	n4 := NewNode(4)
	n5 := NewNode(5)
	n6 := NewNode(6)

	bt := NewBinTree()
	bt.root = n6
	n6.right = n5
	n5.right = n4
	n4.left = n3

	want := "Binary Tree: 6 5 3 4 "
	got := bt.String()

	if !cmp.Equal(got, want) {
		t.Errorf("want %s, got %s", got, want)
	}
	err := bt.printBTstack()
	if err != nil {
		fmt.Println(err)
	}
}

func TestBinTreeRoor(t *testing.T) {
	n3 := NewNode(3)

	bt := NewBinTree()
	bt.root = n3

	want := "Binary Tree: 3 "
	got := bt.String()

	if !cmp.Equal(got, want) {
		t.Errorf("want %s, got %s", got, want)
	}
	err := bt.printBTstack()
	if err != nil {
		fmt.Println(err)
	}
}
