package utils

func TestIntArr() ([]int, []int) {
	return []int{3, 1, 6, 12, 10, 17, 14, 14, 20, 2}, []int{1, 2, 3, 6, 10, 12, 14, 14, 17, 20}
}

func TestIntArrEmpty() ([]int, []int) {
	return []int{}, []int{}
}

func TestIntArrSorted() ([]int, []int) {
	return []int{3, 5, 7}, []int{3, 5, 7}
}

func TestIntArrEl() ([]int, []int) {
	return []int{3}, []int{3}
}

func TestIntArrBackwards() ([]int, []int) {
	return []int{3, 2, 1}, []int{1, 2, 3}
}

func TestIntArrDup() ([]int, []int) {
	return []int{3, 3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3, 3}
}
