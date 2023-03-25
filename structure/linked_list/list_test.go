package ll

import (
	"errors"
	"fmt"
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
	assert.Equal(t, 1, l.First().Value().(int))

	l.Add(2)
	assert.Equal(t, 2, l.First().Value().(int))

	l.Add(3)
	assert.Equal(t, 3, l.First().Value().(int))

	assert.Equal(t, int(3), l.First().Value().(int))
	assert.Equal(t, int(2), l.First().Next().Value().(int))
	assert.Equal(t, int(1), l.First().Next().Next().Value().(int))
	assert.Equal(t, uint32(3), l.Size())

	for i := 4; i <= 20; i++ {
		l.Add(int32(i))
	}
	fmt.Println(l.Size())
	assert.Equal(t, uint32(20), l.Size())
}

func TestLinkedList_Find(t *testing.T) {
	var empty *Node

	l := New()
	for i := 1; i <= 3; i++ {
		l.Add(int32(i))
	}

	n := l.Find(3)
	assert.Equal(t, int32(3), n.Value().(int32))

	n = l.Find(2)
	assert.Equal(t, int32(2), n.Value().(int32))

	n = l.Find(1)
	assert.Equal(t, int32(1), n.Value().(int32))

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
	for i := 1; i <= 3; i++ {
		l.Add(int32(i))
	}

	assert.Equal(t, uint32(3), l.Size())

	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, int32(3), n.Value().(int32))
	assert.Equal(t, uint32(2), l.Size())
	assert.Equal(t, int32(2), l.First().Value().(int32))
	assert.Equal(t, int32(1), l.First().Next().Value().(int32))
	assert.Equal(t, empty, l.First().Next().Next())

	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, int32(2), n.Value().(int32))
	assert.Equal(t, uint32(1), l.Size())
	assert.Equal(t, int32(1), l.First().Value().(int32))
	assert.Equal(t, empty, l.First().Next())

	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, int32(1), n.Value().(int32))
	assert.Equal(t, uint32(0), l.Size())
	assert.Equal(t, empty, l.First())

	n, err = l.Pop()
	if err == nil {
		assert.Error(t, errors.New("error should not be equal nil"))
	}
	assert.Equal(t, empty, n)
	assert.Equal(t, empty, l.First())

	for i := 4; i <= 6; i++ {
		l.Add(int32(i))
	}

	assert.Equal(t, uint32(3), l.Size())

	// ------------ test 1 -----------------
	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, int32(6), n.Value().(int32))
	assert.Equal(t, uint32(2), l.Size())

	// ------------ test 2 -----------------
	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, int32(5), n.Value().(int32))
	assert.Equal(t, uint32(1), l.Size())

	// ------------ test 3 -----------------
	n, err = l.Pop()
	if err != nil {
		assert.Error(t, errors.New("error should be equal nil"))
	}
	assert.Equal(t, int32(4), n.Value().(int32))
	assert.Equal(t, uint32(0), l.Size())
	assert.Equal(t, empty, l.First())

	// ------------ test 4 -----------------
	n, err = l.Pop()
	if err == nil {
		assert.Error(t, errors.New("error should not be equal nil"))
	}
	assert.Equal(t, n, empty)
	assert.Equal(t, err, EmptyList)
}
