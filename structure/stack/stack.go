package stack

import (
	"fmt"
	ll "github.com/wubba-com/algorithms_and_structures/structure/linked_list"
)

// Стек — это коллекция, элементы которой получают
// по принципу «последний вошел, первый вышел» (Last-In-First-Out или LIFO).
// Это значит, что мы будем иметь доступ только к последнему добавленному элементу.

// Не можем получить доступ к произвольному элементу стека.
// Можем только добавлять или удалять элементы с помощью специальных методов
// У стека нет итератора

func New() *Stack {
	return &Stack{
		arr:  make([]int32, 0),
		list: ll.New(),
	}
}

type (
	Stack struct {
		list *ll.LinkedList
		arr  []int32
	}
)

// ------------- Стек на основе массива s.arr ---------------

// ArrayPush добавляет в конец
func (s *Stack) ArrayPush(value int32) {
	s.arr = append(s.arr, value)
}

// ArrayPeek возвращает последний, но не удаляет
func (s *Stack) ArrayPeek() (int32, bool) {
	if n := len(s.arr); n > 0 {
		v := s.arr[len(s.arr)-1]

		return v, true
	}
	return 0, false
}

// ArrayPop - удаляет и возвращает элемент с конца
func (s *Stack) ArrayPop() (int32, bool) {
	if n := len(s.arr); n > 0 {
		v := s.arr[len(s.arr)-1]
		s.arr = s.arr[:n-1]
		return v, true
	}

	return 0, false
}

// ------------- Стек на основе односвязного списка s.list ---------------

// LLPop - удаляет и возвращает элемент с конца O(n)
func (s *Stack) LLPop() (*ll.Node, error) {
	return s.list.Pop()
}

// LLPush добавляет в конец O(n)
func (s *Stack) LLPush(new int32) bool {
	return s.list.Add(new)
}

// LLPeek возвращает последний, но не удаляет O(n)
func (s *Stack) LLPeek() *ll.Node {
	return s.list.Find(1)
}

// PrintStack - сложность O(n)
func PrintStack(stack *Stack) {
	node := stack.list.First()
	for node != nil {
		fmt.Println(node.Value())
		node = node.Next()
	}
}
