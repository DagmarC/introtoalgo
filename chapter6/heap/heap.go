package heap

// Heap in this case is the Max Heap data structure -> parent > left(i) & right(i), so the root is the greatest of all
type Heap struct {
	arr  *[]int // pointer to the array have possibility to sort the array insitu (in place)
	size int
}

func (h *Heap) decreaseSize() {
	h.size--
}

// Left returns the index in the heap arr of the left children, do not check if the index is out of range
func (h *Heap) Left(i int) int {
	return 2 * i
}

// Right returns the index in the heap arr of the right children, do not check if the index is out of range
func (h *Heap) Right(i int) int {
	return 2*i + 1
}

// switchEl switches the elements in the heap arr on position i, j
func (h *Heap) switchEl(i, j int) {
	if i >= h.size || j >= h.size {
		return
	}
	(*h.arr)[i], (*h.arr)[j] = (*h.arr)[j], (*h.arr)[i]
}

// BuildMaxHeap returns the max heap data structure from the given array...BOTTOM-UP
// Note: n/2 + 1 are leaves, so they are 1-element heaps to begin with --> no need fix them
func BuildMaxHeap(arr *[]int) *Heap {
	h := &Heap{
		arr:  arr,
		size: len(*arr),
	}
	for i := (h.size / 2) - 1; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return h

}

// MaxHeapify checks if the element in the heap array rooted at index i fullfills the heap property and that is that it is greater than both of its children.
// If not then the switch occurs and the procedure is called again to propagate down the tree.
// It counts that all subtrees below already fullfills the conditions of the max heap
func (h *Heap) MaxHeapify(i int) {
	if i >= h.size {
		return // index i is larger than the size of the heap array
	}
	l := h.Left(i)
	r := h.Right(i)
	largest := i // the index of the largest of elements of parent at i, left(i), right(i), default is index i

	if l < h.size && (*h.arr)[l] > (*h.arr)[i] {
		largest = l
	}
	if r < h.size && (*h.arr)[r] > (*h.arr)[largest] {
		largest = r
	}
	// if i index == largerst index -> parent at index i is greater than both of its children, so heap condition is not violated
	// if not then SWITCH elements in the heap arrL i and largest indices and repeat the step again with largest index to propagate down the heap
	if i != largest {
		h.switchEl(i, largest)
		h.MaxHeapify(largest)
	}
}
