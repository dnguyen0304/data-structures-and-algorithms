package common

// NewRange creates an integer slice of the given length and with sequential
// values.
//
// The time complexity is O(n), where n is the number of elements in the slice.
func NewRange(length int) []int {
	list := make([]int, length)
	for i := range list {
		list[i] = i
	}
	return list
}
