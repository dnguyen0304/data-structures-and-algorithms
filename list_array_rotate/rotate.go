// package list_array_rotate implements rotation algorithms for array data
// structures.
//
// Using a temporary array is approximately 100 times faster than reversing
// in-place.
//
// # Should the array be rotated to the left or to the right? Would it be
//   acceptable if I use left rotation?
// # Should the array be rotated in-place?
// # Should iteration or recursion be used? Would it be acceptable if I use
//   iteration?
// # What type of data do the array's elements contain?
// # Am I allowed to use other data structures?
// # Does the entire array fit into memory?
package list_array_rotate

import "github.com/dnguyen0304/data-structures-and-algorithms/list_array_reverse"

// LeftByReversing left rotates an integer slice by the given factor.
//
// The operation is applied in-place. The time complexity is O(n), where n is
// the number of elements in the slice.
func LeftByReversing(list []int, factor int) {
	list_array_reverse.Recursive(list[:factor])
	list_array_reverse.Recursive(list[factor:])
	list_array_reverse.Recursive(list)
}

// RightByReversing right rotates an integer slice by the given factor.
//
// The operation is applied in-place. The time complexity is O(n), where n is
// the number of elements in the slice.
func RightByReversing(list []int, factor int) {
	LeftByReversing(list, len(list)-factor)
}

// LeftByTemporaryArray left rotates an integer slice by the given factor.
//
// The operation is applied in-place. The time complexity is O(n), where n is
// the number of elements in the slice.
func LeftByTemporaryArray(list []int, factor int) {
	// Base Case
	if list == nil {
		return
	}
	// Base Case
	// Only pointer-types (i.e. pointers, slices, channels, interfaces, maps,
	// and functions) may be compared to nil.
	if factor == 0 {
		return
	}
	// Iterative Case
	temporary := make([]int, factor)
	copy(temporary, list[:factor])

	for i := 0; i < len(list)-factor; i++ {
		list[i] = list[i+factor]
	}

	copy(list[len(list)-factor:], temporary)
}

// RightByTemporaryArray right rotates an integer slice by the given factor.
//
// The operation is applied in-place. The time complexity is O(n), where n is
// the number of elements in the slice.
func RightByTemporaryArray(list []int, factor int) {
	LeftByTemporaryArray(list, len(list)-factor)
}
