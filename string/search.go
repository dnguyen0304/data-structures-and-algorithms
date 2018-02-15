package string

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
