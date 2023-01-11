package queue

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	dll "github.com/wubba-com/algorithms_and_structures/structure/do_linked_list"
	"testing"
)

func TestQueue_Unshift(t *testing.T) {
	q := New()

	var tests = []int{
		1,
		2,
		3,
	}

	for _, val := range tests {
		t.Run(fmt.Sprintf("%d", val), func(t *testing.T) {
			q.Unshift(val)
			node := q.PeekShift()

			assert.Equal(t, val, node.Value())
		})
	}
}

func TestQueue_ShiftAndUnshift(t *testing.T) {
	q := New()

	var tests = []int{
		1,
		2,
		3,
	}

	for _, val := range tests {
		t.Run(fmt.Sprintf("%d", val), func(t *testing.T) {
			q.Push(val)
			node, err := q.Shift()
			if err != nil {
				return
			}

			assert.Equal(t, val, node.Value())
		})
	}
}

func TestQueue_Push(t *testing.T) {
	q := New()

	var tests = []int{
		1,
		2,
		3,
	}

	for _, val := range tests {
		t.Run(fmt.Sprintf("%d", val), func(t *testing.T) {
			q.Push(val)
			node := q.PeekPop()

			assert.Equal(t, val, node.Value())
		})
	}
}

func TestQueue_PopAndPush(t *testing.T) {
	q := New()

	var tests = []int{
		1,
		2,
		3,
	}

	for _, val := range tests {
		t.Run(fmt.Sprintf("%d", val), func(t *testing.T) {
			q.Push(val)
			node, err := q.Pop()
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, val, node.Value())
		})
	}

	var empty *dll.Node
	node, err := q.Pop()
	if err == nil {
		assert.Error(t, errors.New("error should not be equal nil"))
	}

	// проверка, что список пуст
	assert.Equal(t, err, EmptyList)
	assert.Equal(t, node, empty)
	assert.Equal(t, q.PeekShift(), empty)
	assert.Equal(t, q.PeekPop(), empty)

	var (
		expectedValue = 4
	)

	q.Push(expectedValue)

	var actNode *dll.Node
	actNode, err = q.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, expectedValue, actNode.Value())

	node, err = q.Pop()
	if err != nil {
		assert.Equal(t, err, EmptyList)
	} else {
		assert.Error(t, errors.New("error should not be equal nil"))
	}
}
