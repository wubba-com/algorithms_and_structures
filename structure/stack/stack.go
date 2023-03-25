package stack

import (
	"errors"
	"fmt"
	ll "github.com/wubba-com/algorithms_and_structures/structure/linked_list"
)

type Stack interface {
	Pop() (*Value, error)
	Push(Value)
	Peek() (*Value, bool)
	Print()
	Size() uint32
	Next() bool
}

var (
	ErrEmptyStack = errors.New("stack is empty")
)

// Стек — это коллекция, элементы которой получают
// по принципу «последний вошел, первый вышел» (Last-In-First-Out или LIFO).
// Это значит, что мы будем иметь доступ только к последнему добавленному элементу.

// Не можем получить доступ к произвольному элементу стека.
// Можем только добавлять или удалять элементы с помощью специальных методов
// У стека нет итератора

func NewBasicArray() Stack {
	return &stackArr{
		make([]Value, 0),
	}
}

func NewBasicLinkedList() Stack {
	return &stackLL{
		list: ll.New(),
	}
}

type (
	stackLL struct {
		list *ll.LinkedList
	}

	stackArr struct {
		arr []Value
	}

	Value struct {
		Value interface{}
	}
)

// ------------- Стек на основе массива s.arr ---------------

// Push добавляет в конец
func (s *stackArr) Push(value Value) {
	s.arr = append(s.arr, value)
}

// Peek возвращает последний, но не удаляет
func (s *stackArr) Peek() (*Value, bool) {
	if n := len(s.arr); n > 0 {
		v := s.arr[len(s.arr)-1]

		return &v, true
	}
	return nil, false
}

// Pop - удаляет и возвращает элемент с конца
func (s *stackArr) Pop() (*Value, error) {
	if n := len(s.arr); n > 0 {
		v := s.arr[len(s.arr)-1]
		s.arr = s.arr[:n-1]
		return &v, nil
	}

	return nil, ErrEmptyStack
}

func (s *stackArr) Size() uint32 {
	return uint32(len(s.arr))
}

func (s *stackArr) Next() bool {
	return s.Size() != 0
}

func (s *stackArr) Print() {
	for i := range s.arr {
		if i < len(s.arr)-1 {
			fmt.Printf("%s =>", s.arr[i].Value)
			continue
		}
		fmt.Printf("%s", s.arr[i].Value)
	}
	fmt.Println()
}

// ------------- Стек на основе односвязного списка s.list ---------------

// Pop - удаляет и возвращает элемент с конца O(n)
func (s *stackLL) Pop() (*Value, error) {
	n, err := s.list.Pop()
	if err != nil {
		if errors.Is(err, ll.EmptyList) {
			return nil, ErrEmptyStack
		}
		return nil, err
	}
	return &Value{
		Value: n.Value(),
	}, nil
}

// Push добавляет в конец O(n)
func (s *stackLL) Push(value Value) {
	s.list.Add(value.Value)
}

// Peek возвращает последний, но не удаляет O(1)
func (s *stackLL) Peek() (*Value, bool) {
	el := s.list.First()
	if el == nil {
		return nil, false
	}
	return &Value{
		Value: el.Value(),
	}, true
}

func (s *stackLL) Size() uint32 {
	return s.list.Size()
}

func (s *stackLL) Next() bool {
	return s.Size() != 0
}

// Print - сложность O(n)
func (s *stackLL) Print() {
	node := s.list.First()
	for node != nil {
		if node.Next() == nil {
			fmt.Printf("%v", node.Value())
			node = node.Next()
			continue
		}
		fmt.Printf("%v => ", node.Value())
		node = node.Next()
	}
	fmt.Println()
}
