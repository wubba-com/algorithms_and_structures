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
			q.PushFront(val)
			node := q.PeekFront()

			assert.Equal(t, val, node.Value())
		})
	}
}

func TestQueue_PushBackAndPopFront(t *testing.T) {
	q := New()

	var tests = []int{
		1,
		2,
		3,
	}

	for _, val := range tests {
		t.Run(fmt.Sprintf("%d", val), func(t *testing.T) {
			q.PushBack(val)
			node, err := q.PopFront()
			if err != nil {
				return
			}

			assert.Equal(t, val, node.Value())
		})
	}
}

func TestQueue_PushBack(t *testing.T) {
	q := New()

	var tests = []int{
		1,
		2,
		3,
	}

	for _, val := range tests {
		t.Run(fmt.Sprintf("%d", val), func(t *testing.T) {
			q.PushBack(val)
			node := q.PeekBack()

			assert.Equal(t, val, node.Value())
		})
	}
}

func TestQueue_PopBackAndPushBack(t *testing.T) {
	q := New()

	var tests = []int{
		1,
		2,
		3,
	}

	for _, val := range tests {
		t.Run(fmt.Sprintf("%d", val), func(t *testing.T) {
			q.PushBack(val)
			node, err := q.PopBack()
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, val, node.Value())
		})
	}

	var empty *dll.Node
	node, err := q.PopBack()
	if err == nil {
		assert.Error(t, errors.New("error should not be equal nil"))
	}

	// проверка, что список пуст
	assert.Equal(t, err, EmptyList)
	assert.Equal(t, node, empty)
	assert.Equal(t, q.PeekFront(), empty)
	assert.Equal(t, q.PeekBack(), empty)

	var (
		expectedValue = 4
	)

	q.PushBack(expectedValue)

	var actNode *dll.Node
	actNode, err = q.PopBack()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, expectedValue, actNode.Value())

	node, err = q.PopBack()
	if err != nil {
		assert.Equal(t, err, EmptyList)
	} else {
		assert.Error(t, errors.New("error should not be equal nil"))
	}
}
