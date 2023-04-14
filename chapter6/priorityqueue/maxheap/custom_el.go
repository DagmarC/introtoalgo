package maxheap

type CustomEl struct {
	key       int
	heapIndex int
	name      string
}

func (ce *CustomEl) Key() int {
	return ce.key
}

func (ce *CustomEl) SetKey(key int) {
	ce.key = key
}

func (ce *CustomEl) HeapIndex() int {
	return ce.heapIndex
}

func (ce *CustomEl) UpdateHeapIndex(i int) {
	ce.heapIndex = i
}
