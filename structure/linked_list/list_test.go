package ll

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Count(t *testing.T) {
	l := New()
	assert.Equal(t, l.Size(), uint32(0))
}

func TestLinkedList_Add(t *testing.T) {
	l := New()
	l.Add(1)
	assert.Equal(t, l.First().Value(), int32(1))
	assert.Equal(t, l.End().Value(), int32(1))

	l.Add(2)
	assert.Equal(t, l.First().Value(), int32(1))
	assert.Equal(t, l.End().Value(), int32(2))

	l.Add(3)
	assert.Equal(t, l.First().Value(), int32(1))
	assert.Equal(t, l.End().Value(), int32(3))

	assert.Equal(t, l.First().Value(), int32(1))
	assert.Equal(t, l.First().Next().Value(), int32(2))
	assert.Equal(t, l.First().Next().Next().Value(), int32(3))
	assert.Equal(t, l.End().Value(), int32(3))
	assert.Equal(t, l.Size(), uint32(3))

	l.Add(4)
	l.Add(5)
	l.Add(6)

	l.Add(7)
	l.Add(8)
	l.Add(9)

	l.Add(10)
	l.Add(11)
	l.Add(12)

	l.Add(13)
	l.Add(14)
	l.Add(15)

	l.Add(16)
	l.Add(17)
	l.Add(18)

	l.Add(19)
	l.Add(20)
	l.Add(21)

	assert.Equal(t, l.Size(), uint32(21))
}

func TestLinkedList_Find(t *testing.T) {
	var empty *Node

	l := New()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	n := l.Find(1)
	assert.Equal(t, int32(3), n.Value())

	n = l.Find(2)
	assert.Equal(t, int32(2), n.Value())

	n = l.Find(3)
	assert.Equal(t, int32(1), n.Value())

	n = l.Find(4)
	assert.Equal(t, empty, n)
}

func TestLinkedList_Pop(t *testing.T) {
	var (
		l     *LinkedList
		empty *Node
		err   error
		n     *Node
	)

	l = New()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.Equal(t, l.Size(), uint32(3))

	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(3))
	assert.Equal(t, l.Size(), uint32(2))
	assert.Equal(t, l.First().Value(), int32(1))
	assert.Equal(t, l.First().Next().Value(), int32(2))
	assert.Equal(t, l.End().Value(), int32(2))
	assert.Equal(t, l.First().Next().Next(), empty)

	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(2))
	assert.Equal(t, l.Size(), uint32(1))
	assert.Equal(t, l.First().Value(), int32(1))
	assert.Equal(t, l.End().Value(), int32(1))
	assert.Equal(t, l.First().Next(), empty)

	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(1))
	assert.Equal(t, l.Size(), uint32(0))
	assert.Equal(t, l.First(), empty)
	assert.Equal(t, l.End(), empty)

	n, err = l.Pop()
	if err == nil {
		assert.Error(t, errors.New("error should not be equal nil"))
	}
	assert.Equal(t, n, empty)
	assert.Equal(t, l.First(), empty)
	assert.Equal(t, l.End(), empty)

	l.Add(4)
	l.Add(5)
	l.Add(6)

	assert.Equal(t, l.Size(), uint32(3))

	// ------------ test 1 -----------------
	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(6))
	assert.Equal(t, l.Size(), uint32(2))

	// ------------ test 2 -----------------
	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(5))
	assert.Equal(t, l.Size(), uint32(1))

	// ------------ test 3 -----------------
	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(4))
	assert.Equal(t, l.Size(), uint32(0))
	assert.Equal(t, l.First(), empty)
	assert.Equal(t, l.End(), empty)

	// ------------ test 4 -----------------
	n, err = l.Pop()
	if err == nil {
		assert.Error(t, errors.New("error should not be equal nil"))
	}
	assert.Equal(t, n, empty)
	assert.Equal(t, err, EmptyList)
}

func TestLinkedList_Shift(t *testing.T) {
	var (
		l     *LinkedList
		empty *Node
		err   error
		n     *Node
	)

	l = New()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	assert.Equal(t, l.Size(), uint32(3))

	n, err = l.Shift()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(1))
	assert.Equal(t, l.Size(), uint32(2))
	assert.Equal(t, l.First().Value(), int32(2))
	assert.Equal(t, l.First().Next().Value(), int32(3))
	assert.Equal(t, l.End().Value(), int32(3))
	assert.Equal(t, l.First().Next().Next(), empty)

	n, err = l.Shift()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(2))
	assert.Equal(t, l.Size(), uint32(1))
	assert.Equal(t, l.First().Value(), int32(3))
	assert.Equal(t, l.End().Value(), int32(3))
	assert.Equal(t, l.First().Next(), empty)

	n, err = l.Shift()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(3))
	assert.Equal(t, l.Size(), uint32(0))
	assert.Equal(t, l.First(), empty)
	assert.Equal(t, l.End(), empty)

	n, err = l.Shift()
	if err == nil {
		assert.Error(t, errors.New("error should not be equal nil"))
	}
	assert.Equal(t, n, empty)
	assert.Equal(t, l.First(), empty)
	assert.Equal(t, l.End(), empty)

	l.Add(4)
	l.Add(5)
	l.Add(6)

	assert.Equal(t, l.Size(), uint32(3))

	// ------------ test 1 -----------------
	n, err = l.Shift()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(4))
	assert.Equal(t, l.Size(), uint32(2))

	// ------------ test 2 -----------------
	n, err = l.Shift()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(5))
	assert.Equal(t, l.Size(), uint32(1))

	// ------------ test 3 -----------------
	n, err = l.Shift()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, n.Value(), int32(6))
	assert.Equal(t, l.Size(), uint32(0))
	assert.Equal(t, l.First(), empty)
	assert.Equal(t, l.End(), empty)

	// ------------ test 4 -----------------
	n, err = l.Shift()
	if err == nil {
		assert.Error(t, errors.New("error should not be equal nil"))
	}
	assert.Equal(t, n, empty)
	assert.Equal(t, err, EmptyList)
}
