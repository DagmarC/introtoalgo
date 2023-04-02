package linkedlist

import (
	"fmt"
	"testing"
)

func testData() *Linkedlist[float64] {
	ll := NewLinkedlist[float64]()
	ll.InsertLast(0.21)
	ll.InsertLast(0.34)
	ll.InsertLast(0.32)
	ll.InsertLast(0.11)

	return ll
}

func TestLinkedlistInsertSearch(t *testing.T) {
	ll := testData()

	tc := []struct {
		val      float64
		expexted Node[float64]
	}{
		{val: 0.21, expexted: Node[float64]{key: 0.21}},
		{val: 0.11, expexted: Node[float64]{key: 0.11}},
		{val: 0.32, expexted: Node[float64]{key: 0.32}},
	}

	for _, tt := range tc {
		if n := ll.Search(tt.val); n == nil || n.key != tt.expexted.key {
			t.Errorf("got %v, expected %v", n, tt.expexted)
		}
	}
	ll.Display()
}

func TestLinkedlistInsertSearchNil(t *testing.T) {
	ll := testData()

	tc := []struct {
		val      float64
		expexted *Node[float64] // anything that could be nil
	}{
		{val: 0.22, expexted: nil},
		{val: 0.99, expexted: nil},
		{val: 0.0, expexted: nil},
	}

	for _, tt := range tc {
		if n := ll.Search(tt.val); n != tt.expexted {
			t.Errorf("got %v, expected %v", n, tt.expexted)
		}
	}
	ll.Display()
}

func TestLinkedlistPrepend(t *testing.T) {
	ll := testData()

	val := 0.78
	tmpHead := ll.head.key

	ll.Prepend(val)
	if ll.head.key != val {
		t.Errorf("got %f, want %f", ll.head.key, val)
	}
	if ll.head.next.key != tmpHead {
		t.Errorf("invalid head next pointer, should point to previous head %f but points to %f", tmpHead, ll.head.next.key)
	}

	if ll.head.next.prev.key != val {
		t.Errorf("invalid head prev pointer, old head should point to new head %f but points to %f", val, ll.head.next.prev.key)
	}

	ll.Display()
}

func TestLinkedlistInsert(t *testing.T) {
	ll := testData()

	val := 0.87
	y := ll.head.next
	expNext := y.next.key

	ll.Insert(val, y)

	if y.next.key != val {
		t.Errorf("expected y.next %.2f, got %.2f", val, y.next.key)
	}
	if y.next.prev.key != y.key {
		t.Errorf("expected y.next.prev to be y %.2f, got %.2f", y.key, y.next.prev.key)
	}
	if y.next.next.key != expNext {
		t.Errorf("expected new.next %.2f, got %.2f", expNext, y.next.next.key)
	}
	if y.next.next.prev.key != val {
		t.Errorf("expected new %.2f, got %.2f", val, y.next.next.prev.key)
	}
	ll.Display()
}

func TestLinkedlistInsertLast(t *testing.T) {
	ll := testData()

	newTail := 0.87
	expPrev := ll.tail.key

	ll.InsertLast(newTail)

	if ll.tail.key != newTail {
		t.Errorf("expected tail %.2f, got %.2f", newTail, ll.tail.key)
	}
	if ll.tail.prev.key != expPrev {
		t.Errorf("expected prev %.2f, got %.2f", expPrev, ll.tail.prev.key)
	}
	if ll.tail.prev.next.key != newTail {
		t.Errorf("expected tail.prev.next to be new tail %.2f, got %.2f", newTail, ll.tail.prev.next.key)
	}
	ll.Display()
}

func TestLinkedlistDeleteMiddle(t *testing.T) {
	ll := testData()

	val := ll.head.next.key
	expNext := 0.32

	ll.Delete(val)
	if n := ll.Search(val); n != nil {
		t.Errorf("node with val %f should not be found", val)
	}

	if ll.head.next.key != expNext {
		t.Errorf("invalid next pointer, should point to  %.2f but points to %22f", expNext, ll.head.next.key)
	}
	if ll.head.next.prev != ll.head {
		t.Errorf("invalid prev pointer, should point to  %.2f but points to %.2f", ll.head.key, ll.head.next.prev.key)
	}
	ll.Display()
}

func TestLinkedlistDeleteHead(t *testing.T) {
	ll := testData()

	head := ll.head.key
	newHead := ll.head.next.key
	expNext := ll.head.next.next.key // ll has at least 3 elements

	ll.Delete(head)
	if n := ll.Search(head); n != nil {
		t.Errorf("node with val %f should not be found", head)
	}
	if ll.head.key != newHead {
		fmt.Errorf("new head expected value is %f, got %f", newHead, ll.head.key)
	}
	if ll.head.next.key != expNext {
		t.Errorf("invalid next pointer, should point to  %.2f but points to %22f", expNext, ll.head.next.key)
	}
	if ll.head.next.prev != ll.head {
		t.Errorf("invalid prev pointer, should point to  %.2f but points to %.2f", ll.head.key, ll.head.next.prev.key)
	}
	if ll.head.prev != nil {
		t.Errorf("invali head prev pointer, should point to NIL, but points to %.2f", ll.head.prev.key)
	}
	ll.Display()
}

func TestLinkedlistDeleteTail(t *testing.T) {
	ll := testData()

	tail := ll.tail.key
	newTail := ll.tail.prev.key
	expPrev := ll.tail.prev.prev.key

	ll.Delete(tail)
	if n := ll.Search(tail); n != nil {
		t.Errorf("node with val %f should not be found\n", tail)
	}
	if ll.tail.key != newTail {
		fmt.Errorf("new head expected value is %f, got %f\n", newTail, ll.tail.key)
	}
	if ll.tail.prev.key != expPrev {
		t.Errorf("invalid prev pointer, should point to  %.2f but points to %22f", expPrev, ll.tail.prev.key)
	}
	if ll.tail.prev.next.key != newTail {
		t.Errorf("invalid next pointer, should point to  %.2f but points to %22f", newTail, ll.tail.prev.next.key)
	}
	if ll.tail.next != nil {
		t.Errorf("tail next pointer should points to NIL, but has %.2f", ll.tail.next.key)
	}

	// @TODO pointers test
	ll.Display()
}

func TestLinkedlistSort(t *testing.T) {
	ll := testData()

	expLl := NewLinkedlist[float64]()
	expLl.InsertLast(0.11)
	expLl.InsertLast(0.21)
	expLl.InsertLast(0.32)
	expLl.InsertLast(0.34)

	ll.Sort()
	if !ll.Equals(expLl) {
		fmt.Println("\nERROR: List is not sorted correctly, got:")
		ll.Display()
		fmt.Println("want:")
		expLl.Display()
		t.Errorf("error")
	}
}
