package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFib(t *testing.T) {
	assert.Equal(t, 0, Fib(0))
	assert.Equal(t, 1, Fib(1))
	assert.Equal(t, 1, Fib(2))
	assert.Equal(t, 2, Fib(3))
	assert.Equal(t, 55, Fib(10))
}

func TestFibLine(t *testing.T) {
	assert.Equal(t, 0, Fib(0))
	assert.Equal(t, 1, Fib(1))
	assert.Equal(t, 1, Fib(2))
	assert.Equal(t, 2, Fib(3))
	assert.Equal(t, 55, Fib(10))
}

func TestBinarySearch(t *testing.T) {
	var inputSl1 = []int{1, 3, 5, 6}
	i := BinarySearch(inputSl1, 3)

	assert.Equal(t, inputSl1[i], 3)
}

func TestQuicksort(t *testing.T) {
	var (
		inputSl1 = []int{4, 2, 1}
		inputSl2 = []int{2, 1, 4, 56, 2, 98, 10}

		outputSl1 = []int{1, 2, 4}
		outputSl2 = []int{1, 2, 2, 4, 10, 56, 98}
	)

	Quicksort(inputSl1)
	Quicksort(inputSl2)

	assert.Equal(t, inputSl1, outputSl1)
	assert.Equal(t, inputSl2, outputSl2)
}

func TestSort(t *testing.T) {
	var (
		inputSl1 = []int{4, 2, 1}
		inputSl2 = []int{2, 1, 4, 56, 2, 98, 10}

		outputSl1 = []int{1, 2, 4}
		outputSl2 = []int{1, 2, 2, 4, 10, 56, 98}
	)

	inputSl1 = Sort(inputSl1)
	inputSl2 = Sort(inputSl2)

	assert.Equal(t, inputSl1, outputSl1)
	assert.Equal(t, inputSl2, outputSl2)
}

func TestSortInsert(t *testing.T) {
	var (
		inputSl1 = []int{4, 2, 1}
		inputSl2 = []int{2, 1, 4, 56, 2, 98, 10}

		outputSl1 = []int{1, 2, 4}
		outputSl2 = []int{1, 2, 2, 4, 10, 56, 98}
	)

	SortInsert(inputSl1)
	SortInsert(inputSl2)

	assert.Equal(t, inputSl1, outputSl1)
	assert.Equal(t, inputSl2, outputSl2)
}
