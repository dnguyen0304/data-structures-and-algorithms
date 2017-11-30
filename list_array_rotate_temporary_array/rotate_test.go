package list_array_rotate_temporary_array_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/dnguyen0304/data-structures-and-algorithms/common"
	"github.com/dnguyen0304/data-structures-and-algorithms/list_array_rotate_temporary_array"
)

func TestLeftRotateTimeComplexity(t *testing.T) {
	seconds := time.Now().UTC().Unix()
	// This is global state. However, the standard library does provide APIs
	// for creating local sources and random number generators.
	rand.Seed(seconds)

	// Rhodes' "Thousand-Million" Thought Experiment
	for _, n := range []int{10, 1000, 1000000, 10000000} {
		list := common.NewRange(n)
		factor := rand.Intn(n)
		start := time.Now()
		list_array_rotate_temporary_array.LeftRotate(list, factor)
		elapsed := time.Since(start).Seconds()
		// Width assigns a fixed width whereas precision determines at least
		// how many digits are displayed after the decimal point. The latter
		// may pad zeroes to whole numbers.
		fmt.Printf(
			"When n is equal to %8d, the operation takes %.9f seconds.\n",
			n,
			elapsed)
	}
}
