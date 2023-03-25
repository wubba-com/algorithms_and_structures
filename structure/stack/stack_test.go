package stack

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_ArrayPeek(t *testing.T) {
	var empty *Value
	s := NewBasicArray()
	v, ok := s.Peek()
	assert.Equal(t, empty, v)
	assert.Equal(t, false, ok)
}

func TestStack_ArrayPop(t *testing.T) {
	s := NewBasicArray()
	tests := []int32{
		1,
		2,
		3,
	}

	for _, v := range tests {
		t.Run(fmt.Sprintf("test-%d", v), func(t *testing.T) {
			s.Push(Value{Value: v})
			val, err := s.Pop()
			if err != nil {
				assert.Error(t, errors.New("err should not be equal error"))
			}
			assert.Equal(t, v, val.Value.(int32))
		})
	}
	var empty *Value
	v, err := s.Pop()
	assert.ErrorIs(t, ErrEmptyStack, err)
	assert.Equal(t, empty, v)
}

func TestStack_LLPeek(t *testing.T) {
	var empty *Value
	s := NewBasicLinkedList()
	v, ok := s.Peek()
	assert.Equal(t, false, ok)
	assert.Equal(t, empty, v)

}

func TestStack_LLPush(t *testing.T) {
	s := NewBasicLinkedList()

	s.Push(Value{Value: int32(1)})
	v, ok := s.Peek()
	assert.Equal(t, true, ok)
	assert.Equal(t, int32(1), v.Value.(int32))

	s.Push(Value{Value: int32(2)})
	v, ok = s.Peek()
	assert.Equal(t, true, ok)
	assert.Equal(t, int32(2), v.Value.(int32))

	s.Push(Value{Value: int32(3)})
	v, ok = s.Peek()
	assert.Equal(t, true, ok)
	assert.Equal(t, int32(3), v.Value.(int32))
}

func TestStack_LLPop(t *testing.T) {
	var (
		empty *Value
		tests = []int32{
			1, 2, 3,
		}
	)
	s := NewBasicLinkedList()
	for i := range tests {
		s.Push(Value{Value: tests[i]})
		n, err := s.Pop()
		if err != nil {
			assert.Error(t, errors.New("ошибка должна быть равна nil"))
		}
		assert.Equal(t, tests[i], n.Value.(int32))
	}

	n, err := s.Pop()
	if err == nil {
		assert.Error(t, errors.New("ошибка не должна быть равна nil"))
	}
	assert.ErrorIs(t, ErrEmptyStack, err)
	assert.Equal(t, empty, n)
}
