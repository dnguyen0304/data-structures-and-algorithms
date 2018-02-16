package string

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestKMPSearch(t *testing.T) {
	// Set Up
	var cases = []struct {
		name     string
		s        string
		substr   string
		expected bool
	}{
		{"AlwaysFind", "", "", true},
		{"NeverFind", "", "foo", false},
		{"PrefixesNotNeeded", "foobar", "bar", true},
		{"PrefixesNeeded", "abca$abcab", "abcab", true},
	}

	// Test
	msg := "kmpSearch(%q, %q) => expected %t"
	for _, c := range cases {
		t.Run(c.name, func(*testing.T) {
			if found := kmpSearch(c.s, c.substr); found != c.expected {
				t.Errorf(msg, c.s, c.substr, c.expected)
			}
		})
	}
}

func TestKMPSearchTimeComplexity(t *testing.T) {
	for _, n := range []int{10, 100, 1000, 1000000} {
		s := strings.Repeat("abca$", n)
		substr := "abcab"

		start := time.Now()
		kmpSearch(s, substr)
		elapsed := time.Since(start).Seconds()
		fmt.Printf(
			"When n is equal to %7d, the kmpSearch operation takes %.9f seconds.\n",
			n,
			elapsed)
	}
}

func BenchmarkKMPSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kmpSearch("abca$abcab", "abcab")
	}
}

func TestNewPrefixArray(t *testing.T) {
	// Set Up
	var cases = []struct {
		name     string
		s        string
		expected []int
	}{
		{"NoSuffixes", "abc", []int{0, 0, 0}},
		{"ShortSuffix", "abca", []int{0, 0, 0, 1}},
		{"LongSuffix", "abcab", []int{0, 0, 0, 1, 2}},
		{"MultipleSuffixes", "abcababc", []int{0, 0, 0, 1, 2, 1, 2, 3}},
		{"IsDefensive", "", []int{}},
	}

	// Test
	msg := "newPrefixArray(%s) => expected %v but actually %v"
	for _, c := range cases {
		t.Run(c.name, func(*testing.T) {
			prefixes := newPrefixArray(c.s)
			if ok := assertSliceEquals(c.expected, prefixes); !ok {
				t.Errorf(msg, c.s, c.expected, prefixes)
			}
		})
	}
}

func TestNewPrefixArrayTimeComplexity(t *testing.T) {
	for _, n := range []int{10, 100, 1000, 1000000} {
		s := strings.Repeat("foo", n)

		start := time.Now()
		newPrefixArray(s)
		elapsed := time.Since(start).Seconds()
		fmt.Printf(
			"When n is equal to %7d, the newPrefixArray operation takes %.9f seconds.\n",
			n,
			elapsed)
	}
}

func BenchmarkNewPrefixArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newPrefixArray("foo")
	}
}

func assertSliceEquals(actual, expected []int) bool {
	if actual == nil || expected == nil {
		return false
	}
	if len(actual) != len(expected) {
		return false
	}
	for i := range expected {
		if actual[i] != expected[i] {
			return false
		}
	}
	return true
}
