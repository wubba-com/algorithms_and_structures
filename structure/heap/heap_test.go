package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeap_Push(t *testing.T) {
	h := NewMax()
	h.Push(10, nil)

	h.Push(20, nil)
	assert.Equal(t, 20, h.sl[0].priority)

	h.Push(30, nil)
	assert.Equal(t, 3, h.heapSize)
	assert.Equal(t, 3, h.size())
	assert.Equal(t, 30, h.sl[0].priority)
	assert.Equal(t, 20, h.sl[h.heapSize-1].priority)
}

func TestHeap_Pop(t *testing.T) {
	h := NewMax()
	h.Push(10, nil)
	h.Push(20, nil)
	h.Push(30, nil)
	h.Push(40, nil)
	h.Push(50, nil)
	h.Push(60, nil)
	h.Push(70, nil)
	h.Push(80, nil)
	n := h.Pop()
	assert.Equal(t, 80, n.priority)
	assert.Equal(t, 70, h.sl[0].priority)

	n = h.Pop()
	assert.Equal(t, 70, n.priority)
	assert.Equal(t, 60, h.sl[0].priority)

	n = h.Pop()
	assert.Equal(t, 60, n.priority)
	assert.Equal(t, 50, h.sl[0].priority)

	n = h.Pop()
	assert.Equal(t, 50, n.priority)
	assert.Equal(t, 40, h.sl[0].priority)

	n = h.Pop()
	assert.Equal(t, 40, n.priority)
	assert.Equal(t, 30, h.sl[0].priority)

	n = h.Pop()
	assert.Equal(t, 30, n.priority)
	assert.Equal(t, 20, h.sl[0].priority)

	n = h.Pop()
	assert.Equal(t, 20, n.priority)
	assert.Equal(t, 10, h.sl[0].priority)

	n = h.Pop()
	assert.Equal(t, 10, n.priority)
	assert.Equal(t, 0, len(h.sl))

	n = h.Pop()
	var empty *node
	assert.Equal(t, empty, n)
	assert.Equal(t, 0, len(h.sl))
}
