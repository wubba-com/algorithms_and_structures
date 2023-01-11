package stack

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	ll "testing/structure/linked_list"
)

func TestStack_ArrayPeek(t *testing.T) {
	s := New()
	v, ok := s.ArrayPeek()
	assert.Equal(t, uint32(0), v)
	assert.Equal(t, false, ok)
}

func TestStack_ArrayPop(t *testing.T) {
	s := New()
	tests := []int32{
		1,
		2,
		3,
	}

	for _, v := range tests {
		t.Run(fmt.Sprintf("test-%d", v), func(t *testing.T) {
			s.ArrayPush(v)
			val, ok := s.ArrayPop()
			if ok == false {
				assert.Error(t, errors.New("ok should not be equal false"))
			}
			assert.Equal(t, v, val)
		})
	}

	v, ok := s.ArrayPop()
	assert.Equal(t, false, ok)
	assert.Equal(t, int32(0), v)
}

func TestStack_LLPeek(t *testing.T) {
	var empty *ll.Node
	s := New()

	assert.Equal(t, empty, s.LLPeek())
}

func TestStack_LLPush(t *testing.T) {
	s := New()

	s.LLPush(1)
	assert.Equal(t, int32(1), s.LLPeek().Value())

	s.LLPush(2)
	assert.Equal(t, int32(2), s.LLPeek().Value())

	s.LLPush(3)
	assert.Equal(t, int32(3), s.LLPeek().Value())
}

func TestStack_LLPop(t *testing.T) {
	var empty *ll.Node
	s := New()

	// -------------- test 1 ---------------------
	if ok := s.LLPush(1); !ok {
		assert.Error(t, errors.New("not added - 1"))
	}
	n, err := s.LLPop()
	if err != nil {
		assert.Error(t, errors.New("ошибка должна быть равна nil"))
	}
	assert.Equal(t, int32(1), n.Value())

	// -------------- test 2 ---------------------
	if ok := s.LLPush(2); !ok {
		assert.Error(t, errors.New("not added - 2"))
	}
	n, err = s.LLPop()
	if err != nil {
		assert.Error(t, errors.New("ошибка должна быть равна nil"))
	}
	assert.Equal(t, int32(2), n.Value())

	// -------------- test 3 ---------------------
	if ok := s.LLPush(3); !ok {
		assert.Error(t, errors.New("not added - 3"))
	}
	n, err = s.LLPop()
	if err != nil {
		assert.Error(t, errors.New("ошибка должна быть равна nil"))
	}
	assert.Equal(t, int32(3), n.Value())

	// -------------- test 4 ---------------------
	n, err = s.LLPop()
	if err == nil {
		assert.Error(t, errors.New("ошибка не должна быть равна nil"))
	}
	assert.Equal(t, empty, n)
}
