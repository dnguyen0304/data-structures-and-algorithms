package common

import (
	"math/rand"
	"time"
)

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

// NewRandomRange creates an integer slice of the given size and with random
// values.
//
// The time complexity is O(n), where n is the number of elements in the slice.
func NewRandomRange(length int) []int {
	seconds := time.Now().UTC().Unix()
	// This is global state. However, the standard library does provide APIs
	// for creating local sources and random number generators.
	rand.Seed(seconds)

	list := make([]int, length)
	for i := range list {
		list[i] = rand.Intn(length)
	}
	return list
}
