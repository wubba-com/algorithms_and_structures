package dll

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_AddFirst(t *testing.T) {
	var (
		l     = New()
		tests = []int{
			1,
			2,
			3,
		}
	)

	for _, val := range tests {
		t.Run(fmt.Sprintf("value-%d", val), func(t *testing.T) {
			l.AddFirst(val)
		})
	}

	var empty *Node

	// next link
	assert.Equal(t, tests[0], l.last.Value())            // 1
	assert.Equal(t, tests[2], l.first.Value())           // 3
	assert.Equal(t, tests[1], l.first.next.Value())      // 2
	assert.Equal(t, tests[0], l.first.next.next.Value()) // 1
	assert.Equal(t, empty, l.first.next.next.next)       // nil

	// reverse link
	last := l.first.next.next
	assert.Equal(t, l.last.Value(), last.Value())     // 1
	assert.Equal(t, tests[0], last.Value())           // 1
	assert.Equal(t, tests[1], last.prev.Value())      // 2
	assert.Equal(t, tests[2], last.prev.prev.Value()) // 3
	assert.Equal(t, empty, last.prev.prev.prev)       // nil
}

func TestLinkedList_AddLast(t *testing.T) {
	var (
		l     = New()
		tests = []int{
			1,
			2,
			3,
		}
	)

	for _, val := range tests {
		t.Run(fmt.Sprintf("value-%d", val), func(t *testing.T) {
			l.AddLast(val)
		})
	}

	var empty *Node
	// next link
	assert.Equal(t, tests[2], l.last.Value())            // 3
	assert.Equal(t, tests[0], l.first.Value())           // 1
	assert.Equal(t, tests[1], l.first.next.Value())      // 2
	assert.Equal(t, tests[2], l.first.next.next.Value()) // 3
	assert.Equal(t, empty, l.first.next.next.next)       // nil

	// reverse link
	last := l.first.next.next
	assert.Equal(t, last, l.last)                     // 3
	assert.Equal(t, tests[2], last.Value())           // 3
	assert.Equal(t, tests[1], last.prev.Value())      // 2
	assert.Equal(t, tests[0], last.prev.prev.Value()) // 1
	assert.Equal(t, empty, last.prev.prev.prev)       // nil
}

func TestLinkedList_Find(t *testing.T) {
	var (
		l     = New()
		tests = []int{
			1,
			2,
			3,
		}
	)

	for _, val := range tests {
		t.Run(fmt.Sprintf("last-%d", val), func(t *testing.T) {
			l.AddLast(val)
		})
	}

	n := l.Find(0)
	assert.Equal(t, tests[0], n.Value()) // 1
	assert.Equal(t, l.first, n)          // 1
	n = l.Find(1)
	assert.Equal(t, tests[1], n.Value()) // 2
	n = l.Find(2)
	assert.Equal(t, tests[2], n.Value()) // 3
	assert.Equal(t, l.last, n)           // 3

	// reverse
	n = l.Find(-1)
	assert.Equal(t, tests[2], n.Value()) // 3
	assert.Equal(t, l.last, n)           // 3
	n = l.Find(-2)
	assert.Equal(t, tests[1], n.Value()) // 2
	n = l.Find(-3)
	assert.Equal(t, tests[0], n.Value()) // 1
	assert.Equal(t, l.first, n)          // 1

	// обновим список для добавления элементов с новым методом
	l = New()
	for _, val := range tests {
		t.Run(fmt.Sprintf("first-%d", val), func(t *testing.T) {
			l.AddFirst(val)
		})
	}

	// next
	n = l.Find(0)
	assert.Equal(t, tests[2], n.Value()) // 3
	assert.Equal(t, l.first, n)          // 3
	n = l.Find(1)
	assert.Equal(t, tests[1], n.Value()) // 2
	n = l.Find(2)
	assert.Equal(t, tests[0], n.Value()) // 1
	assert.Equal(t, l.last, n)           // 1

	// reverse
	n = l.Find(-1)
	assert.Equal(t, tests[0], n.Value()) // 1
	assert.Equal(t, l.last, n)           // 1
	n = l.Find(-2)
	assert.Equal(t, tests[1], n.Value()) // 2
	n = l.Find(-3)
	assert.Equal(t, tests[2], n.Value()) // 3
	assert.Equal(t, l.first, n)          // 3
}
