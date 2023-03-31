package utils

func EqualIntArr(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, elA := range a {
		if elA != b[i] {
			return false
		}
	}
	return true
}