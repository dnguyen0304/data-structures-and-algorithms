package heap_binary_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/dnguyen0304/data_structures_and_algorithms/common"
	"github.com/dnguyen0304/data_structures_and_algorithms/heap_binary"
)

func TestHeapSize(t *testing.T) {
	expected := 1
	output := heap_binary.NewMax(make([]int, expected)).Size()
	if expected != output {
		t.Fail()
	}
}

func TestHeapIsEmpty(t *testing.T) {
	size := 0
	heap := heap_binary.NewMax(make([]int, size))
	if !heap.IsEmpty() {
		t.Fail()
	}
}

func TestHeapIsNotEmpty(t *testing.T) {
	size := 1
	heap := heap_binary.NewMax(make([]int, size))
	if heap.IsEmpty() {
		t.Fail()
	}
}

func TestHeapPop(t *testing.T) {
	length := 10
	expected := length - 1
	heap := heap_binary.NewMax(common.NewRange(length))
	if expected != heap.Pop() || expected != heap.Size(){
		t.Fail()
	}
}

func TestHeapPeek(t *testing.T) {
	length := 10
	expected := length - 1
	output := heap_binary.NewMax(common.NewRange(length)).Peek()
	if expected != output {
		t.Fail()
	}
}

func TestHeapSort(t *testing.T) {
	input, expected := []int{2, 4, 0, 1, 3}, []int{0, 1, 2, 3, 4}
	heap := heap_binary.NewMax(input)
	heap.Sort()
	// reflect.DeepEqual should not be used in production. It is approximately
	// 100 times slower than loops.
	if !reflect.DeepEqual(expected, input) {
		t.Fail()
	}
}

func TestHeapSortTimeComplexity(t *testing.T) {
	// Rhodes' "Thousand-Million" Thought Experiment
	for _, length := range []int{10, 100, 1000, 1000000} {
		array := common.NewRandomRange(length)
		heap := heap_binary.NewMax(array)
		start := time.Now()
		heap.Sort()
		elapsed := time.Since(start).Seconds()
		fmt.Printf(
			"When n is equal to %8d, the Sort operation takes %.9f seconds.\n",
			length,
			elapsed)
	}
}

func TestNewMaxTimeComplexity(t *testing.T) {
	for _, length := range []int{10, 100, 1000, 1000000} {
		array := common.NewRandomRange(length)
		start := time.Now()
		heap_binary.NewMax(array)
		elapsed := time.Since(start).Seconds()
		fmt.Printf(
			"When n is equal to %8d, the NewMax operation takes %.9f seconds.\n",
			length,
			elapsed)
	}
}

func TestGetParent(t *testing.T) {
	input, expected := 1, 0
	output := heap_binary.GetParent(input)
	if expected != output {
		t.Fail()
	}
}

func TestGetLeftChild(t *testing.T) {
	input, expected := 0, 1
	output := heap_binary.GetLeftChild(input)
	if expected != output {
		t.Fail()
	}
}

func TestGetRightChild(t *testing.T) {
	input, expected := 0, 2
	output := heap_binary.GetRightChild(input)
	if expected != output {
		t.Fail()
	}
}

func TestAssertMaxHeapPropertySucceeds(t *testing.T) {
	array := []int{2, 0, 1}
	heap_binary.NewMax(array)
	if heap_binary.AssertMaxHeapProperty(array) != nil {
		t.Fail()
	}
}

func TestAssertMaxHeapPropertyFails(t *testing.T) {
	array := []int{0, 1, 2}
	if heap_binary.AssertMaxHeapProperty(array) == nil {
		t.Fail()
	}
}

func TestMaxObservesMaxHeapProperty(t *testing.T) {
	length := 10
	array := common.NewRandomRange(length)
	heap_binary.NewMax(array)
	if heap_binary.AssertMaxHeapProperty(array) != nil {
		t.Fail()
	}
}
