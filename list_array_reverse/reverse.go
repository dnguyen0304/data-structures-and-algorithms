// The list_array_reverse package implements reversal algorithms for array
// data structures.
//
// # Should the array be reversed in-place?
// # What type of data do the array's elements contain?
// # Am I allowed to use other data structures?
// # Does the entire array fit into memory?
package list_array_reverse

// Reverse reverses an integer slice.
//
// The operation is applied in-place. The time complexity is O(n), where n is
// the number of elements in the slice.
func Recursive(list []int) {
	recursive(list, 0, len(list)-1)
}

func recursive(list []int, left int, right int) {
	// Base Case
	if list == nil {
		return
	}
	// Base Case
	if left >= right {
		return
	}
	// Recursive Case
	list[left], list[right] = list[right], list[left]
	recursive(list, left+1, right-1)
}
