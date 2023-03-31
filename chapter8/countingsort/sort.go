package countingsort

// Countingsort sorts the incoming array A of len n, k determines the max number that appears in array A, returns new array B, uses tmp array C to count digit occurence
func Countingsort(arr []int, n, k int) []int {

	c := make([]int, k+1) // arr that will count elements from <0-k>, C[0] = 2 means...0 belongs at index 2-1, and then 1-1 (-1 because resultant arr B starts from 0 index, not 1)
	b := make([]int, n)   // resultant array B

	for j := 0; j < len(arr); j++ {
		c[arr[j]]++ // count elements in arr - index in C means elem C[i] means count of elem i
	}
	// elements in arr C are indices where to put elements in arr B
	for i := 1; i < len(c); i++ {
		c[i] = c[i] + c[i-1] // do not forget to decrement index, this counts that resultant arr B is from 1-n, not from 0 to n-1
	}
	// stable sort - from the end of arr A
	for j := len(b) - 1; j >= 0; j-- {
		b[c[arr[j]]-1] = arr[j] // decrement c[arr[j]] cuz B is from 0 to n-1
		c[arr[j]]--             // If there are duplicit elements in arr A, then decrease the index to place another element on the next position, not the same one
	}
	return b
}

