package queue

import (
	"errors"
	dll "github.com/wubba-com/algorithms_and_structures/structure/do_linked_list"
)

// errors

var (
	EmptyList = errors.New("queue is empty")
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
	len  int
	list *dll.LinkedList
}

func (q *Queue) Len() int {
	return q.len
}

// PeekFront
// Поведение: Возвращает первый элемент.
// Сложность: O(n)
func (q *Queue) PeekFront() *dll.Node {
	return q.list.Find(0)
}

// PeekBack
// Поведение: Возвращает последний элемент.
// Сложность: O(n)
func (q *Queue) PeekBack() *dll.Node {
	return q.list.Find(-1)
}

// PushFront
// Поведение: Добавляет элемент в начало очереди.
// Сложность: O(1)
func (q *Queue) PushFront(item interface{}) {
	defer func() {
		q.len += 1
	}()

	q.list.AddFirst(item)
}

// PushBack
// Поведение: Добавляет элемент в конец очереди.
// Сложность: O(1)
func (q *Queue) PushBack(item interface{}) {
	defer func() {
		q.len += 1
	}()

	q.list.AddLast(item)
}

// PopFront
// Поведение: Удаляет первый элемент из очереди и возвращает его. Если очередь пустая, кидает ошибку.
// Сложность: O(1).
func (q *Queue) PopFront() (*dll.Node, error) {
	if first := q.list.PopFirst(); first != nil {
		defer func() {
			q.len -= 1
		}()
		return first, nil
	}
	return nil, EmptyList
}

// PopBack
// Поведение: Удаляет первый помещенный элемент из очереди и возвращает его. Если очередь пустая, кидает ошибку.
// Сложность: O(1).
func (q *Queue) PopBack() (*dll.Node, error) {
	if last := q.list.PopLast(); last != nil {
		defer func() {
			q.len -= 1
		}()

		return last, nil
	}

	return nil, EmptyList
}

func Print(q *Queue) {
	q.list.Print()
}
