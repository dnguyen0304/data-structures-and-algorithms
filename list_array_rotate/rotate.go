// The list_array_rotate_reversing package implements a rotation algorithm for
// array data structures.
//
// # Should the array be rotated to the left or to the right?
// # Should the array be rotated in-place?
// # What type of data do the array's elements contain?
// # Am I allowed to use other data structures?
// # Does the entire array fit into memory?
package list_array_rotate_reversing

import "github.com/dnguyen0304/data-structures-and-algorithms/list_array_reverse_recursive"

// LeftRotate left rotates an integer slice by the given factor.
//
// The operation is applied in-place. The time complexity is O(n), where n is
// the number of elements in the slice.
func LeftRotate(list []int, factor int) {
	list_array_reverse_recursive.Reverse(list[:factor])
	list_array_reverse_recursive.Reverse(list[factor:])
	list_array_reverse_recursive.Reverse(list)
}

// RightRotate right rotates an integer slice by the given factor.
//
// The operation is applied in-place. The time complexity is O(n), where n is
// the number of elements in the slice.
func RightRotate(list []int, factor int) {
	LeftRotate(list, len(list)-factor)
}
