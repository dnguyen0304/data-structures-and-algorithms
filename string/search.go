package string

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
