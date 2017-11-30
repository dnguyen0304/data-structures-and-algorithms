package list_array_reverse_recursive_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dnguyen0304/data-structures-and-algorithms/common"
	"github.com/dnguyen0304/data-structures-and-algorithms/list_array_reverse_recursive"
)

func TestReverseTimeComplexity(t *testing.T) {
	// Rhodes' "Thousand-Million" Thought Experiment
	for _, length := range []int{10, 100, 1000, 1000000} {
		list := common.NewRandomRange(length)
		start := time.Now()
		list_array_reverse_recursive.Reverse(list)
		elapsed := time.Since(start).Seconds()
		// Width assigns a fixed width whereas precision determines at least
		// how many digits are displayed after the decimal point. The latter
		// may pad zeroes to whole numbers.
		fmt.Printf(
			"When n is equal to %8d, the operation takes %.9f seconds.\n",
			length,
			elapsed)
	}
}
