// Package heap_binary implements binary heap data structures.
//
// Common usages include sorting and priority queues.
//
// # What type of heap is this? Would it be acceptable to assume the heap is a
//   maximum binary heap?
// # What operations should the heap support?
// # What type of data structure should the heap be backed by? Would it be
//   acceptable to use a dynamic array rather than a binary tree?
//
// # What type of data does the data structure contain?
// # Am I allowed to use other data structures?
// # Am I allowed to implement other data structures?
package heap_binary

import "errors"

type Heap interface {
	Size() int
	IsEmpty() bool
	Push(int)
	Pop() int
	Peek() int
	Sort()
}

// Max is an integer max binary heap data structure backed by an underlying
// array.
type Max struct {
	array []int
	size int
}

// Size counts the number of elements in the heap.
//
// The time complexity is O(1).
func (heap *Max) Size() int {
	// This must not use len(heap.array) because the size can be a subset of
	// the underlying array.
	return heap.size
}

// IsEmpty checks if the heap contains elements.
//
// The time complexity is O(1).
func (heap *Max) IsEmpty() bool {
	return heap.Size() == 0
}

// Push (MAX-HEAP-INSERT) adds a new element to the heap.
//
// The time complexity is O(log n), where n is the number of elements in the heap.
func (heap *Max) Push(int) {
}

// Pop (HEAP-EXTRACT-MAX) removes the root element from the heap.
//
// The time complexity is O(log n), where n is the number of elements in the heap.
func (heap *Max) Pop() int {
	root, last := 0, heap.Size() - 1
	heap.array[root], heap.array[last] = heap.array[last], heap.array[root]
	heap.size--
	heap.heapifyDown(root)
	return heap.array[last]
}

// Peek (HEAP-MAXIMUM) gets but does not remove the root element of the heap.
//
// The time complexity is O(1).
func (heap *Max) Peek() int {
	root := 0
	return heap.array[root]
}

// Sort (HEAPSORT) sorts the heap.
//
// The operation is applied in-place. The time complexity is O(n log n), where
// n is the number of elements in the heap.
func (heap *Max) Sort() {
	for last := heap.Size() - 1; last >= 1; last-- {
		// This should not remove elements from the heap.
		heap.Pop()
	}
}

// heapifyDown (MAX-HEAPIFY) restores the max heap property for the binary
// tree whose root is at the given index.
//
// As a precondition, this assumes the left and right subtrees observe the max
// heap property. The time complexity is O(log n), where is the number of nodes
// in the binary tree. The time complexity can also be represented as O(n),
// where n is the height of the binary tree.
func (heap *Max) heapifyDown(i int) {
	// Base Case: out of bounds parent index
	if i < 0 || i >= heap.Size() {
		return
	}

	largest, left, right := i, GetLeftChild(i), GetRightChild(i)

	// Check the left child.
	if left < heap.Size() && heap.array[left] > heap.array[i] {
		largest = left
	}
	// Check the right child.
	if right < heap.Size() && heap.array[right] > heap.array[largest] {
		largest = right
	}
	if largest != i {
		// Swap the parent with the larger child.
		heap.array[i], heap.array[largest] = heap.array[largest], heap.array[i]
		// Recursive Case
		heap.heapifyDown(largest)
	}
}

// NewMax (BUILD-MAX-HEAP) creates an integer max binary heap.
//
// The operation is applied in-place. The time complexity is O(n), where n is
// the number of elements in the array.
func NewMax(array []int) *Max {
	heap := &Max{array: array, size: len(array)}
	// This must iterate in reverse because as a precondition, heapifyDown
	// assumes the left and right subtrees observe the max heap property.
	// This should not heapify the indices greater than len(array)/2 because
	// those subtrees are leaf nodes and therefore already observe the max
	// heap property.
	for i := len(array)/2; i >= 0; i-- {
		heap.heapifyDown(i)
	}
	return heap
}

// GetParent gets the index of the parent for the child at the given index.
//
// The time complexity is O(1).
func GetParent(i int) int {
	return (i-1)/2
}

// GetLeftChild gets the index of the left child for the parent at the given
// index.
//
// The time complexity is O(1).
func GetLeftChild(i int) int {
	return 2*i + 1
}

// GetRightChild gets the index of the right child for the parent at the given
// index.
//
// The time complexity is O(1).
func GetRightChild(i int) int {
	return 2*i + 2
}

// AssertMaxHeapProperty asserts if the array observes the max heap property.
//
// The value of parents is always larger than those of its children. The time
// complexity is O(n), where n is the number of elements in the array.
func AssertMaxHeapProperty(array []int) error {
	var err error
	message := "The array does not observe the max heap property."

	for i, left, right := 0, 0, 0; i < len(array)/2; i++ {
		left, right = GetLeftChild(i), GetRightChild(i)

		// no children
		// This assumption is valid because heaps behave similarly to complete
		// binary trees (breadth-first search).
		if left >= len(array) {
			break
		}
		// larger left child
		if array[left] > array[i] {
			err = errors.New(message)
			break
		}
		// no right child
		if right >= len(array) {
			break
		}
		// larger right child
		if array[right] > array[i] {
			err = errors.New(message)
			break
		}
	}
	return err
}
