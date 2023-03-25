package ll

import "errors"

var (
	EmptyList = errors.New("list is empty")
)

func New() *LinkedList {
	return &LinkedList{}
}

type (
	LinkedList struct {
		first *Node
		size  uint32
	}

	Node struct {
		val  interface{}
		next *Node
	}
)

func (n *Node) Value() interface{} {
	return n.val
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *LinkedList) Size() uint32 {
	return l.size
}

func (l *LinkedList) First() *Node {
	return l.first
}

// Pop Удаляет и возвращает последний элемент O(1)
func (l *LinkedList) Pop() (*Node, error) {
	if l.size == 0 && l.first == nil {
		return nil, EmptyList
	}

	defer func() {
		l.size -= 1
	}()

	defer func() {
		l.first = l.first.Next()
	}()

	return l.first, nil
}

// Add Добавляет в конец O(1)
func (l *LinkedList) Add(value interface{}) {
	newNode := &Node{val: value}
	newNode.next = l.first
	l.first = newNode
	l.size += 1
}

func (l *LinkedList) reverse() *Node {
	node := l.first
	var (
		current = node
		prev    *Node
		next    *Node
	)

	for current != nil {
		next = current
		current.next = prev
		prev = current
		current = next
	}
	node = prev

	return node
}

// Find Алгоритм для поиска в односвязном списке k-го элемента с конца - сложность O(n)
func (l *LinkedList) Find(k int) *Node {
	if k <= 0 {
		return nil
	}

	p1 := l.first
	p2 := l.first

	// не должны учитывать сдвиг на единицу
	for i := 0; i < k-1; i++ {
		if p2 == nil {
			return nil
		}

		p2 = p2.next
	}

	if p2 == nil {
		return nil
	}

	for p2.next != nil {
		p2 = p2.next
		p1 = p1.next
	}

	return p1
}

// AddKElement Добавить за k-элемент - сложность O(n)
func (l *LinkedList) AddKElement(value int32, k int) bool {
	node := l.first
	kNode := l.Find(k)
	if node == nil {
		return false
	}

	newNode := &Node{val: value}
	newNode.next = kNode.next
	kNode.next = newNode
	l.size += 1

	return true
}
