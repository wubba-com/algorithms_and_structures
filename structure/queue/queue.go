package queue

import (
	"errors"
	dll "github.com/wubba-com/algorithms_and_structures/structure/do_linked_list"
)

// errors

var (
	EmptyList = errors.New("list is empty")
)

// Очереди не дают доступа к произвольному элементу, но,
// в отличие от стека, элементы кладутся (enqueue) и забираются
// (dequeue) с разных концов. Такой метод называется «первый вошел, первый вышел»
// (First-In-First-Out или FIFO). То есть забирать элементы из очереди
// мы будем в том же порядке, что и клали. Как реальная очередь или конвейер.

func New() *Queue {
	return &Queue{
		list: dll.New(),
	}
}

type Queue struct {
	list *dll.LinkedList
}

// PeekShift
// Поведение: Возвращает первый элемент.
// Сложность: O(n)
func (q *Queue) PeekShift() *dll.Node {
	return q.list.Find(0)
}

// PeekPop
// Поведение: Возвращает последний элемент.
// Сложность: O(n)
func (q *Queue) PeekPop() *dll.Node {
	return q.list.Find(-1)
}

// Unshift
// Поведение: Добавляет элемент в начало очереди.
// Сложность: O(1)
func (q *Queue) Unshift(item int) {
	q.list.AddFirst(item)
}

// Push
// Поведение: Добавляет элемент в конец очереди.
// Сложность: O(1)
func (q *Queue) Push(item int) {
	q.list.AddLast(item)
}

// Shift
// Поведение: Удаляет первый элемент из очереди и возвращает его. Если очередь пустая, кидает ошибку.
// Сложность: O(1).
func (q *Queue) Shift() (*dll.Node, error) {
	if first := q.list.PopFirst(); first != nil {
		return first, nil
	}
	return nil, EmptyList
}

// Pop
// Поведение: Удаляет первый помещенный элемент из очереди и возвращает его. Если очередь пустая, кидает ошибку.
// Сложность: O(1).
func (q *Queue) Pop() (*dll.Node, error) {
	if last := q.list.PopLast(); last != nil {
		return last, nil
	}

	return nil, EmptyList
}

func Print(q *Queue) {
	q.list.Print()
}
