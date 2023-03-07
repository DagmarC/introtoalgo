package heapsort

// Heapsort returns the sorted arr a1 <= a2 <= ... <= an
// 1. build max heap from the array -> max el at the top
// 2. go from the bottom up to 2nd el ->
// .... switch first-max el and the last el
// .... decrease h.size to cut out the alredy sorted max el at the bottom after switch
// .... correct the heap property on the 1st (switched el) - MaxHeapify procedure
func Heapsort(arr []int) {
	h := BuildMaxHeap(&arr)
	for i := len(arr) - 1; i > 0; i-- {
		h.switchEl(0, i) // exchange 1st element (current max) with the last element, starts with the last one and continue down up to 2nd element
		h.decreaseSize() // decrease heap size to cut out the last el. -> current last el is already switched and it is the max el
		h.MaxHeapify(0)  // fix the heap property of the switched element from the bottom
	}
}
