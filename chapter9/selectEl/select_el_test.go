package selectEl

import (
	"testing"

	"github.com/DagmarC/introtoalgo/utils"
)

var errEl int = -1

func TestSelectElFromArr(t *testing.T) {
	arr, sortedArr := utils.TestIntArr()
	tmpArr := make([]int, len(arr))
	copy(tmpArr, arr)
	for i := range arr {
		el := RandomizedSelect(&tmpArr, 0, len(tmpArr)-1, i+1) // ith+1 smallest el, means that in resultedsorted arr it is at index i
		if el != sortedArr[i] {
			t.Errorf("iteration n.%d, got %d, want %d", i+1, el, sortedArr[i])
		}
		copy(tmpArr, arr)
	}
}

func TestSelectElFromArrEmpty(t *testing.T) {
	elIn := 3
	arr, _ := utils.TestIntArrEmpty()
	el := RandomizedSelect(&arr, 0, len(arr)-1, elIn)
	if el != errEl {
		t.Errorf("got %d, want %d", el, errEl)
	}
}

func TestSelectElFromArrSorted(t *testing.T) {
	elIn := 2
	arr, sortedArr := utils.TestIntArrSorted()
	el := RandomizedSelect(&arr, 0, len(arr)-1, elIn)
	if el != sortedArr[elIn-1] {
		t.Errorf("got %d, want %d", el, sortedArr[elIn-1])
	}
}

func TestSelectElFromArrOneEl(t *testing.T) {
	elIn := 1
	arr, sortedArr := utils.TestIntArrEl()
	el := RandomizedSelect(&arr, 0, len(arr)-1, elIn)
	if el != sortedArr[elIn-1] {
		t.Errorf("got %d, want %d", el, sortedArr[elIn-1])
	}
}

func TestSelectElFromArrBackwards(t *testing.T) {
	elIn := 3
	arr, sortedArr := utils.TestIntArrBackwards()
	el := RandomizedSelect(&arr, 0, len(arr)-1, elIn)
	if el != sortedArr[elIn-1] {
		t.Errorf("got %d, want %d", el, sortedArr[elIn-1])
	}
}

func TestSelectElFromArrDup(t *testing.T) {
	elIn := 3
	arr, sortedArr := utils.TestIntArrDup()
	el := RandomizedSelect(&arr, 0, len(arr)-1, elIn)
	if el != sortedArr[elIn-1] {
		t.Errorf("got %d, want %d", el, sortedArr[elIn-1])
	}
}
