package string

// kmpSearch searches for the substring ("needle") in the string ("haystack").
//
// It is backed by the Knuth–Morris–Pratt algorithm. The goal is to avoid
// backtracking in the haystack.
//
// The time complexity is O(m+n), where m and n are the number of characters
// in the string and substring, respectively.
func kmpSearch(s, substr string) bool {
	// Base Case
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}

	// Iterative Case
	i := 0
	j := 0
	prefixes := newPrefixArray(substr)

	for i < len(s) && j < len(substr) {
		if s[i] != substr[j] {
			if j != 0 {
				j = prefixes[j-1]
				continue
			}
			i++
			continue
		}
		i++
		j++
	}
	if j == len(substr) {
		return true
	}
	return false
}

// newPrefixArray creates a new array to track where substring suffixes are
// the same as the prefixes.
//
// The time complexity is O(n), where n is the number of characters in the
// string.
func newPrefixArray(s string) []int {
	prefixes := make([]int, len(s))

	if len(prefixes) == 0 {
		return prefixes
	}
	prefixes[0] = 0

	for i, j := 1, 0; i < len(prefixes); {
		if s[i] != s[j] {
			if j > 0 {
				j = prefixes[j-1]
				continue
			}
			i++
			continue
		}
		prefixes[i] = j + 1
		i++
		j++
	}
	return prefixes
}
