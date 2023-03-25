package dll

import "fmt"

type LinkedList struct {
	count int
	first *Node
	last  *Node
}

func New() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) First() *Node {
	return l.first
}

func (l *LinkedList) Last() *Node {
	return l.last
}

// PopFirst Удаление первого элемента из списка. Оценка сложности: O(1)
func (l *LinkedList) PopFirst() *Node {
	if l.first != nil {
		node := l.first
		if l.first.next != nil {
			l.first = l.first.next
			l.first.prev = nil

			return node
		} else {
			l.first = nil
			l.last = nil

			return node
		}
	}

	return nil
}

// PopLast Удаление последнего элемента из списка. Оценка сложности: O(1)
func (l *LinkedList) PopLast() *Node {
	if l.last != nil {
		n := l.last
		// если не последний элемент в очереди
		if l.last.prev != nil {
			l.last = l.last.prev
			l.last.setNext(nil)

			return n
			// если последний элемент
		} else {
			l.first = nil
			l.last = nil

			return n
		}
	}

	return nil
}

// AddFirst - в методе обращаемся сразу первому элементу - сложность O(1), что гораздо быстрее 100000 элементов - 3.209ms
func (l *LinkedList) AddFirst(value interface{}) {
	n := &Node{val: value}

	if l.first == nil {
		l.first = n
		l.last = n

		return
	}

	bind(n, l.first)
	//node.setNext(l.first)
	//l.first.setPrev(node)
	l.first = n
}

// AddLast добавление в конец списка. Оценка сложности: O(1)
func (l *LinkedList) AddLast(value interface{}) {
	n := &Node{val: value}

	if l.first == nil {
		l.first = n
		l.last = n

		return
	}

	bind(l.last, n)
	//node.setPrev(l.last)
	//l.last.setNext(node)
	l.last = n
}

// Find поиск по k в списке. Может передать отрицательное значение, тогда элементы будут браться с конца - сложность O(n)
// Если использовать до половины с разных сторон, то оценка сложности O(n/2)
func (l *LinkedList) Find(k int) *Node {
	if k < 0 {
		k = -k
		return l.findReverse(k)
	}

	return l.find(k)
}

// Метод find - поиск сначала списка Оценка сложности: O(n)
func (l *LinkedList) find(k int) *Node {
	var i int
	n := l.first
	for i != k {
		if n == nil {
			return nil
		}
		n = n.next
		i++
	}

	return n
}

// Метод findReverse - поиск с конца списка. Оценка сложности: O(n)
func (l *LinkedList) findReverse(k int) *Node {
	var i = 1
	n := l.last
	for i != k {
		if n == nil {
			return nil
		}

		n = n.prev
		i++
	}

	return n
}

// -------------- Медленная ------------

// AddFirst - в функции придется идти наверх - сложность O(n). 100000 элементов - 7s
func AddFirst(node *Node, item *Node) {
	for node.prev != nil {
		node = node.prev
	}

	item.next = node
	node.prev = item
}

// -------------- Вспомогательные методы для просмотра связей ---------------

// Print - сверху вниз O(n)
func (l *LinkedList) Print() {
	n := l.first

	for n != nil {
		fmt.Printf("%v ", n.val)
		n = n.next
	}
	fmt.Println()
}

// PrintList идет снизу на вверх и печатает сверху вниз O(2*N)
func PrintList(node *Node) {
	if node == nil {
		return
	}
	for node.prev != nil {
		if node == nil {
			return
		}

		node = node.prev
	}

	for node != nil {
		if node == nil {
			return
		}
		fmt.Println(node.val)
		node = node.next
	}
}

func bind(node, with *Node) {
	node.setNext(with)
	with.setPrev(node)
}
