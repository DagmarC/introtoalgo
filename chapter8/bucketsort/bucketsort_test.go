package bucketsort

import "testing"

var arr []float64 = []float64{0.13, 0.91, 0.44, 0.65, 0.11, 0.10, 0.98, 0.99, 0.55, 0.68, 0.74, 0.76, 0.15}
func TestBucketSort(t *testing.T) {
	Bucketsort(arr, len(arr))
}