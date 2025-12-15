# KMP Algorithm (Knuth-Morris-Pratt)

## Features

- **Time Complexity:** Linear time complexity $O(N + M)$.
- **Space Complexity** $O(M)$
- **Full Matches:** Returns all matched starting indices of the substring in the input.

## Main Idea

- **`computeLPS`**: This function calculates the number of prefix matches against the pattern itself. It traverses the substring of length `m`.
  - The index `i` runs from 1 to `m`.
  - The index `prev_lps` increases (or decreases) at most `m` times. 
  - Overall, this results in at most `3 * m` operations, which simplifies to **O(M)**.

- **`MatchSubstring`**: This function returns all matched indices in the input of length `n`. It traverses the input with the index `i` from 0 to `n`.
  - The index `j` points to the current position in the substring.
  - Since the index `i` never decreases, `j` can only increase (or decrease) at most `n` times in total.
  - Overall, this results in at most `3 * n` operations, which is **O(N)**.

- **Result**: By combining both phases, the total time complexity is **O(N + M)**.

- **Space Complexity**: I allocate a single integer array `lps` of size `m` (length of the substring) so it is **O(M)**. Pointers `i`, `j`, `prev_lps` takes a constant extra space, but we do not take constants into consideration. 

- **Note on Output:** There is also the `result` array to store the matches. In the worst case (e.g., finding "A" in "AAAA"), this can grow to **O(N)**, but this is considered *Output Space*, not the algorithm's internal complexity.

## Usage

```go
package main

import (
	"fmt"
	"your-module/kmp"
)

func main() {
	input := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	// Returns a slice of starting indices: [10]
	matches := kmp.MatchSubstring(input, pattern)

	if len(matches) > 0 {
		fmt.Printf("Pattern found at indices: %v\n", matches)
	} else {
		fmt.Println("Pattern not found")
	}
}