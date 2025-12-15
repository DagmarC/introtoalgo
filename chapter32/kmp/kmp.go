package kmp

// matchSubstring we have input and substring to be matched inside input. It returns the first index of the match.
// If no match was found - return -1
// Time Complexity: O(N + M) where N = len(input), M = len(substring)
func MatchSubstring(input, substring string) []int {
	result := make([]int, 0)

	if len(substring) == 0 {
		return result
	}

	lps := computeLPS(substring)

	j := 0 // pointer on the substring match
	for i := 0; i < len(input); {
		if (input[i] == substring[j]) {
			i++
			j++
		} else if j == 0 {
			// Mismatch at the very first character: move input pointer
			i++ 
		} else {
			// Mismatch at the input[i] != substring[j]
			j = lps[j - 1]
		}

		if j == len(substring) {
			// match was found, since pointer j exceed its length, return the first index of the match
			result = append(result, i - j)
			j = lps[j-1] // continue looking the next match
		}
	}
	return result
}

// computeLPS builds the Longest Prefix Suffix array.
// lps[i] stores the length of the longest proper prefix of substring[0...i]
// that is also a suffix of substring[0...i]
// Time Complexity: O(M)
func computeLPS(substring string) []int {
	lps := make([]int, len(substring))
	
	lps[0] = 0  
	prev_lps := 0  // previous lps index

	for i := 1; i < len(substring); {
		if substring[prev_lps] == substring[i] {
			// prefix matches suffix
			lps[i] = prev_lps + 1
			prev_lps++
			i++
		} else if prev_lps == 0 {
			// no previous prefix matches suffix (example substring "AB" at index i = 1 and prev_lps = 0)
			lps[i] = 0
			i++
		} else {
			// prev_lps is differrent from 0, so change its value to the previous lps pointer from the lps array
			// Firstly AAAC = A(1)A(2)A(3)C  you are comparing if A(3) matches C -> NO so go back to prev_lps = lsp[prev_lps-1]
			// Compare if  A(2) matches C -? NO prev_lps = lsp[prev_lps-1]
			// Compare if  A(2) matches C -? NO prev_lps = lsp[prev_lps-1] .. etc until prev_lps is 0
			prev_lps = lps[prev_lps - 1]
		}
	}
	return lps
}