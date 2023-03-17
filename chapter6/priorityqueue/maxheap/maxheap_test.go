package maxheap

import "testing"

func getArr() []HeapEl {
	var arr []HeapEl
	var ce1 HeapEl = &CustomEl{key: 3, name: "ala"}
	var ce2 HeapEl = &CustomEl{key: 5, name: "bala"}
	var ce3 HeapEl = &CustomEl{key: 1, name: "mini"}
	var ce4 HeapEl = &CustomEl{key: 30, name: "maxi"}
	var ce5 HeapEl = &CustomEl{key: 16, name: "mid"}
	var ce6 HeapEl = &CustomEl{key: 10, name: "omid"}

	arr = append(arr, ce1, ce2, ce3, ce4, ce5, ce6)
	return arr
}

func TestBuildHeapArr(t *testing.T) {
	// Insert must work correctly in order to make the BuildMaxHeap fn work.

	expectedV, expectedI := 30, 0
	arr := getArr()
	mh := BuildMaxHeap(arr)
	max, err := mh.Max()
	if err != nil {
		t.Errorf("error %v, expected key %d at index %d\n", err, expectedV, expectedI)
	}
	if max.Key() != expectedV || max.HeapIndex() != expectedI {
		t.Errorf("got key %d at index %d, expected key %d at index %d\n", max.Key(), max.HeapIndex(), expectedV, expectedI)
	}
	expL, li := 16, 1
	expR, ri := 10, 2
	left := (*mh.arr)[mh.Left(0)]
	right := (*mh.arr)[mh.Right(0)]

	if left.Key() != expL || left.HeapIndex() != li {
		t.Errorf("got key %d at index %d, expected key %d at index %d\n", left.Key(), left.HeapIndex(), expL, li)
	}
	if right.Key() != expR || right.HeapIndex() != ri {
		t.Errorf("got key %d at index %d, expected key %d at index %d\n", right.Key(), right.HeapIndex(), expR, ri)
	}
}

func TestExtractMax(t *testing.T) {
	expectedV, expectedI := 30, 0
	arr := getArr()
	mh := BuildMaxHeap(arr)
	max, err := mh.ExtractMax()
	if err != nil {
		t.Errorf("error %v, expected key %d at index %d\n", err, expectedV, expectedI)
	}
	if max.Key() != expectedV || max.HeapIndex() != expectedI {
		t.Errorf("got extracted key %d at index %d, expected key %d at index %d\n", max.Key(), max.HeapIndex(), expectedV, expectedI)
	}

	expectedV, expectedI = 16, 0
	max, err = mh.Max()
	if err != nil {
		t.Errorf("error %v, expected key %d at index %d\n", err, expectedV, expectedI)
	}
	if max.Key() != expectedV || max.HeapIndex() != expectedI {
		t.Errorf("got key %d at index %d, expected key %d at index %d\n", max.Key(), max.HeapIndex(), expectedV, expectedI)
	}
}

func TestIncreaseKey(t *testing.T) {
	expKey, index, expIndex := 40, 3, 0
	arr := getArr()
	mh := BuildMaxHeap(arr)
	err := mh.IncreaseKey(index, expKey)
	if err != nil {
		t.Errorf("error %v, expected increase key at index %d to value %d to be at index %d\n", err, index, expKey, expIndex)
	}
	max, err := mh.Max()
	if err != nil {
		t.Errorf("error %v, expected key %d at index %d\n", err, expKey, expIndex)
	}
	if max.Key() != expKey || max.HeapIndex() != expIndex {
		t.Errorf("got key %d at index %d, expected key %d at index %d\n", max.Key(), max.HeapIndex(), expKey, expIndex)
	}
	// Control the switch of parents and indices
	expL, li := 30, 1
	expLL, lli := 16, 3
	left := (*mh.arr)[mh.Left(0)]
	lleft := (*mh.arr)[mh.Left(1)]

	if left.Key() != expL || left.HeapIndex() != li {
		t.Errorf("got key %d at index %d, expected key %d at index %d\n", left.Key(), left.HeapIndex(), expL, li)
	}
	if lleft.Key() != expLL || lleft.HeapIndex() != lli {
		t.Errorf("got key %d at index %d, expected key %d at index %d\n", lleft.Key(), lleft.HeapIndex(), expLL, lli)
	}
}

func TestDecreaseKey(t *testing.T) {
	// MaxHeapify fn is tested with DecreaseKey fn as well
	expMax, expKey, index, expIndex := 16, 2, 0, 3
	arr := getArr()
	mh := BuildMaxHeap(arr)
	err := mh.DecreaseKey(index, expKey) // change key 30[0] to 2[3]
	if err != nil {
		t.Errorf("error %v, expected decrease key at index %d to value %d to be at index %d\n", err, index, expKey, expIndex)
	}
	max, err := mh.Max()
	if err != nil {
		t.Errorf("error %v, expected key %d at index %d\n", err, expKey, expIndex)
	}
	// 1st elem
	if max.Key() != expMax || max.HeapIndex() != index {
		t.Errorf("got key %d at index %d, expected key %d at index %d\n", max.Key(), max.HeapIndex(), expMax, index)
	}
	// Control the switch of parents and indices
	expL, li := 3, 1
	expLL, lli := 2, 3
	left := (*mh.arr)[mh.Left(0)]
	lleft := (*mh.arr)[mh.Left(1)]

	if left.Key() != expL || left.HeapIndex() != li {
		t.Errorf("got key %d at index %d, expected key %d at index %d\n", left.Key(), left.HeapIndex(), expL, li)
	}
	if lleft.Key() != expLL || lleft.HeapIndex() != lli {
		t.Errorf("got key %d at index %d, expected key %d at index %d\n", lleft.Key(), lleft.HeapIndex(), expLL, lli)
	}
	mh.Print(0)
}

func TestDecreaseKeyErr(t *testing.T) {
	expKey, index, key := 10, 5, 5
	arr := getArr()
	mh := BuildMaxHeap(arr)
	err := mh.DecreaseKey(index, expKey) // change key 5[5] to 10[] -> err
	if err == nil {
		t.Errorf("Decrease of key %d to key %d at index %d should not be successfull", key, expKey, index)
	}
	el := (*mh.arr)[index] // see if no change has occurred
	if el.Key() != key && el.HeapIndex() != index {
		t.Errorf("Key %d at i %d should not be changed, expected key %d at i %d", el.Key(), el.HeapIndex(), key, index)
	}
}
