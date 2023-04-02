package linkedlist

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	key  T
	prev *Node[T]
	next *Node[T]
}

type Linkedlist[T constraints.Ordered] struct {
	head *Node[T]
	tail *Node[T]
}

func NewNode[T constraints.Ordered](val T) *Node[T] {
	return &Node[T]{key: val}
}

func NewLinkedlist[T constraints.Ordered]() *Linkedlist[T] {
	return &Linkedlist[T]{head: nil, tail: nil}
}

func (ll *Linkedlist[T]) Display() {
	x := ll.head
	fmt.Printf("HEAD--->")
	for x != nil {
		fmt.Printf("| %v |--->", x.key)
		x = x.next
	}
	fmt.Printf("NIL \n\n")
}

// Search in a linked list will return the first value that will match or nil
// O(n)
func (ll *Linkedlist[T]) Search(val T) *Node[T] {
	x := ll.head
	for x != nil && x.key != val {
		x = x.next
	}
	return x
}

// Prepend will make a new head out of n
// O(1)
func (ll *Linkedlist[T]) Prepend(val T) {
	n := NewNode(val)
	n.next = ll.head
	n.prev = nil
	if ll.head != nil {
		ll.head.prev = n
	}
	ll.head = n // make new head
}

// Insert new Node into the linked list right after y Node
func (ll *Linkedlist[T]) Insert(newVal T, y *Node[T]) {
	new := NewNode(newVal)

	// insert first element
	if y == nil && ll.head == nil {
		ll.head = new
		ll.tail = new
		return
	}
	if y == ll.tail {
		ll.tail = new
	}
	new.next = y.next
	new.prev = y
	if y.next != nil {
		y.next.prev = new
	}
	y.next = new
}

// Insert new Node into the linked list right after y Node
func (ll *Linkedlist[T]) InsertLast(value T) {
	new := &Node[T]{key: value}

	// insert first element
	if ll.head == nil {
		ll.head = new
		ll.tail = new
		return
	}
	ll.tail.next = new
	new.prev = ll.tail
	ll.tail = new
}

// @TODO -> not working correctly yet
func (ll *Linkedlist[T]) Delete(val T) {
	n := ll.Search(val)
	if n == nil {
		return
	}
	// switch next pointer
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		ll.head = n.next // either value or nil
	}
	// switch prev pointer
	if n.next != nil {
		n.next.prev = n.prev // either value or nil
	} else {
		ll.tail = n.prev
	}
}

func (ll *Linkedlist[T]) Sort() {
	if ll.head == nil || ll.head.next == nil {
		return
	}

	for x := ll.head.next; x != nil; x = x.next {
		tmpX := x.key
		j := x.prev
		for j != nil && tmpX < j.key {
			j.next.key = j.key
			j = j.prev
		}

		if j == nil {
			ll.head.key = tmpX
		} else {
			j.next.key = tmpX
		}
		fmt.Println("\nTMP SORT RESULT, sorting node", tmpX)
		ll.Display()
	}
}

func (ll *Linkedlist[T]) Equals(ll2 *Linkedlist[T]) bool {
	if ll.head == nil && ll2.head == nil {
		return true
	}
	if ll.head == nil && ll2.head != nil {
		return false
	}
	if ll.head != nil && ll2.head == nil {
		return false
	}

	n1 := ll.head
	n2 := ll2.head
	for n1 != nil {
		if n2 == nil {
			return false // n1 is not nil
		}
		if n1.key != n2.key {
			return false
		}
		n1 = n1.next
		n2 = n2.next
	}
	return n2 == nil // if yes then both are equal
}

// func InsertionSort(a []int) error {
// 	if len(a) == 0 {
// 		return errors.New("ERROR: empty array")
// 	}
// 	for i := 1; i < len(a); i++ {
// 		card := a[i] // card  to be inserted at the correct place
// 		j := i - 1
// 		for j >= 0 && card < a[j] { // while card is less than a[j], shift positions by 1 in arr A
// 			a[j+1] = a[j]
// 			j--
// 		}
// 		a[j+1] = card // insert card at correct place: j+1, j is where the condition was not fulfilled
// 	}
// 	return nil
// }
