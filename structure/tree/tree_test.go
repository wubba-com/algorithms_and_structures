package tree

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

//           1
//         /  \
//        0    2
//      /       \
//    -2         5
//   /         /
// -1         3
func TestTree_Add(t *testing.T) {
	bt := New(1)
	assert.Equal(t, 1, bt.val)

	err := bt.Add(2)
	if err != nil {
		assert.Error(t, errors.New("err should be equal nil"))
	}
	assert.Equal(t, 2, bt.right.val)

	err = bt.Add(0)
	if err != nil {
		assert.Error(t, errors.New("err should be equal nil"))
	}
	assert.Equal(t, 0, bt.left.val)

	err = bt.Add(-2)
	if err != nil {
		assert.Error(t, errors.New("err should be equal nil"))
	}
	assert.Equal(t, -2, bt.left.left.val)

	err = bt.Add(-1)
	if err != nil {
		assert.Error(t, errors.New("err should be equal nil"))
	}
	assert.Equal(t, -1, bt.left.left.right.val)

	err = bt.Add(5)
	if err != nil {
		assert.Error(t, errors.New("err should be equal nil"))
	}
	assert.Equal(t, 5, bt.right.right.val)

	err = bt.Add(3)
	if err != nil {
		assert.Error(t, errors.New("err should be equal nil"))
	}
	assert.Equal(t, 3, bt.right.right.left.val)
}

func TestBinaryTree_Inorder(t *testing.T) {
	bt := New(1)
	if err := bt.Add(2); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(0); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(-2); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(-1); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(5); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(3); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}

	// получить все значения дерева
	fn := func(c chan int) func(val int) {
		return func(val int) {
			c <- val
		}
	}
	var (
		c      = make(chan int)
		values = make([]int, 0)
	)
	go func() {
		for v := range c {
			values = append(values, v)
		}
	}()
	call := fn(c)

	bt.Inorder(call)
	close(c)

	assert.Equal(t, []int{-2, -1, 0, 1, 2, 3, 5}, values)
}

func TestBinaryTree_Preorder(t *testing.T) {
	bt := New(1)
	if err := bt.Add(2); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(0); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(-2); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(-1); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(5); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(3); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}

	fn := func(c chan int) func(val int) {
		return func(val int) {
			c <- val
		}
	}
	var (
		c      = make(chan int)
		values = make([]int, 0)
	)
	go func() {
		for v := range c {
			values = append(values, v)
		}
	}()
	call := fn(c)

	bt.Preorder(call)
	close(c)

	assert.Equal(t, []int{1, 0, -2, -1, 2, 5, 3}, values)
}

func TestBinaryTree_Postorder(t *testing.T) {
	bt := New(1)
	if err := bt.Add(2); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(0); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(-2); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(-1); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(5); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}
	if err := bt.Add(3); err != nil {
		assert.Error(t, errors.New("err should be == nil"))
	}

	fn := func(c chan int) func(val int) {
		return func(val int) {
			c <- val
		}
	}
	var (
		c      = make(chan int)
		values = make([]int, 0)
	)
	go func() {
		for v := range c {
			values = append(values, v)
		}
	}()
	call := fn(c)

	bt.Postorder(call)
	close(c)

	assert.Equal(t, []int{-1, -2, 0, 3, 5, 2, 1}, values)
}
