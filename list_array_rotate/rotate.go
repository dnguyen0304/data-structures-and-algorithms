// The list_array_rotate package implements rotation algorithms for array data
// structures.
//
// # Should the array be rotated to the left or to the right?
// # Should the array be rotated in-place?
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
