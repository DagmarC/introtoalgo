package maxheap

import (
	"errors"
	"fmt"
	"math"
)

type HeapEl interface {
	Key() int
	SetKey(key int)
	HeapIndex() int
	UpdateHeapIndex(i int)
}

type MaxHeap struct {
	arr  *[]HeapEl // pointer to arr data structure with T (pointer values, that contain key attribute)
	size int
}

func BuildMaxHeap(arr []HeapEl) *MaxHeap {
	maxHeap := &MaxHeap{arr: &arr, size: 1}
	arr[0].UpdateHeapIndex(0) // update the mapping to the heap index

	// 1st elem is max heap on his own, start with the second one
	for i := 1; i < len(arr); i++ {
		maxHeap.Insert(arr[i])
	}
	return maxHeap
}

// Left returns the index in the heap arr of the left children, do not check if the index is out of range
func (h *MaxHeap) Left(i int) int {
	return 2*i + 1
}

// Right returns the index in the heap arr of the right children, do not check if the index is out of range
func (h *MaxHeap) Right(i int) int {
	return 2*i + 2
}

// Right returns the index in the heap arr of the right children, do not check if the index is out of range
func (h *MaxHeap) Parent(i int) int {
	return i / 2
}

// switchEl switches the elements in the heap arr on position i, j and update mapping of the elements accordingly
func (h *MaxHeap) switchEl(i, j int) {
	// fmt.Println(i, j)
	if i >= h.size || j >= h.size {
		return
	}
	(*h.arr)[i], (*h.arr)[j] = (*h.arr)[j], (*h.arr)[i]

	(*h.arr)[i].UpdateHeapIndex(i) // Update mapping of heap indexing
	(*h.arr)[j].UpdateHeapIndex(j) // Update mapping of heap indexing
}

// MaxHeapify checks if the element in the heap array rooted at index i fullfills the heap property and that is that it is greater than both of its children.
// If not then the switch occurs and the procedure is called again to propagate down the tree.
// It counts that all subtrees below already fullfills the conditions of the max heap
func (h *MaxHeap) MaxHeapify(i int) {
	if i >= h.size {
		return // index i is larger than the size of the heap array
	}
	l := h.Left(i)
	r := h.Right(i)
	largest := i // the index of the largest of elements of parent at i, left(i), right(i), default is index i

	if l < h.size && (*h.arr)[l].Key() > (*h.arr)[i].Key() {
		largest = l
	}
	if r < h.size && (*h.arr)[r].Key() > (*h.arr)[largest].Key() {
		largest = r
	}
	// if i index == largerst index -> parent at index i is greater than both of its children, so heap condition is not violated
	// if not then SWITCH elements in the heap arrL i and largest indices and repeat the step again with largest index to propagate down the heap
	if i != largest {
		h.switchEl(i, largest) // switch and update mapping - elements position in heap
		h.MaxHeapify(largest)
	}

}

// ================ PRIORITY QUEUE FN ====================
// =======================================================

func (h *MaxHeap) Max() (HeapEl, error) {
	if h.size < 1 {
		return nil, errors.New("heap underflow")
	}
	return (*h.arr)[0], nil
}

func (h *MaxHeap) ExtractMax() (HeapEl, error) {
	max, err := h.Max()
	if err != nil {
		return nil, err
	}
	(*h.arr)[0] = (*h.arr)[h.size-1] // last elem to the front
	h.size--                         // decrease the heap size
	h.MaxHeapify(0)                  // correct the placement of the first element

	return max, nil
}

// IncreaseKey will increase an existing element in heap array at position i with a new key value, key must be greater the the existing one
func (h *MaxHeap) IncreaseKey(i, key int) error {
	if i >= h.size || i < 0 {
		return fmt.Errorf("index error, shuld be from 0-%d and is %d", h.size-1, i)
	}

	el := (*h.arr)[i]
	if key <= el.Key() {
		return fmt.Errorf("existing key=%d is already greater than the new key=%d", el.Key(), key)
	}
	el.SetKey(key)
	for i > 0 && (*h.arr)[h.Parent(i)].Key() < (*h.arr)[i].Key() {
		h.switchEl(i, h.Parent(i)) // switch elems and update mapping - elements position in heap
		i = h.Parent(i)            // update index i with the parent index

	}
	return nil
}

// IncreaseKey will increase an existing element in heap array at position i with a new key value, key must be greater the the existing one
func (h *MaxHeap) DecreaseKey(i, key int) error {
	if i >= h.size || i < 0 {
		return fmt.Errorf("index error, shuld be from 0-%d and is %d", h.size-1, i)
	}
	el := (*h.arr)[i]
	if key >= el.Key() {
		return fmt.Errorf("existing key=%d is already smaller than the new key=%d", el.Key(), key)
	}
	el.SetKey(key)
	h.MaxHeapify(i)
	return nil
}

func (h *MaxHeap) Insert(el HeapEl) error {
	h.size++ // increase heap size

	key := el.Key()         // store key, IncreaseKey method will update it to this key
	el.SetKey(-math.MaxInt) // needed, cuz IncreaseKey method will otherwise fail cuz key would be smaller or equal to the existing key

	(*h.arr)[h.size-1] = el
	(*h.arr)[h.size-1].UpdateHeapIndex(h.size - 1) // map to index heap size in the array
	return h.IncreaseKey(h.size-1, key)
}

func (h *MaxHeap) Print(i int) {
	if i >= h.size {
		return
	}
	fmt.Printf("key %d, heap index %d at i=%d \n", (*h.arr)[i].Key(), (*h.arr)[i].HeapIndex(), i)
	h.Print(h.Left(i))
	h.Print(h.Right(i))
}
