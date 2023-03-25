package heap

import "fmt"

// NewMax
// Куча (англ. heap) — абстрактная структура данных, поддерживающая следующие операции:
//    Нахождение минимума.
//    Удаление минимума.
//    Добавление нового элемента в кучу.
// 1. Значение в любой вершине не больше, чем значения её потомков.
// 2. У любой вершины не более двух сыновей.
// 3. Слои заполняются последовательно сверху вниз и слева направо, без «дырок».
func NewMax() *Heap {
	return &Heap{
		sl: make([]*node, 0),
	}
}

type Heap struct {
	sl       []*node
	heapSize int
}

type node struct {
	priority int
	value    []byte
}

func (n *node) Value() []byte {
	return n.value
}

func (h *Heap) Size() int {
	return h.size()
}

func (h *Heap) size() int {
	return len(h.sl)
}

func (h *Heap) Pop() *node {
	if h.size() < 1 {
		return nil
	}
	n := h.sl[0]
	if h.size() < 2 {
		h.sl = make([]*node, 0, cap(h.sl))
		return n
	}
	h.sl[0], h.sl[h.heapSize-1] = h.sl[h.heapSize-1], h.sl[0]
	h.sl = h.sl[:h.heapSize-1]
	h.setSizeHeap()
	h.siftDown(0)

	return n
}

func (h *Heap) Top() *node {
	if h.heapSize != 0 {
		return h.sl[0]
	}

	return nil
}

func (h *Heap) Push(pr int, bs []byte) {
	h.sl = append(h.sl, &node{
		priority: pr,
		value:    bs,
	})
	h.setSizeHeap()
	h.siftUp(h.heapSize - 1)
}

// Работа процедуры: если элемент меньше своего отца,
// условие 1 соблюдено для всего дерева, и больше
// ничего делать не нужно. Иначе, мы меняем местами его с отцом.
// После чего выполняем siftUp для этого отца. Иными словами,
// большой элемент всплывает наверх.
// Процедура выполняется за время O(logn).
func (h *Heap) siftUp(i int) {
	parent := h.parent(i)
	for i > 0 && h.sl[parent].priority < h.sl[i].priority {
		h.sl[i], h.sl[parent] = h.sl[parent], h.sl[i]
		i = parent
		parent = h.parent(i)
	}
}

// Работа процедуры: если i-й элемент меньше, чем его сыновья,
// всё поддерево уже является кучей, и делать ничего не надо.
// В противном случае меняем местами i-й элемент с наибольшим
// из его сыновей, после чего выполняем siftDown для этого сына.
// Процедура выполняется за время O(logn)
func (h *Heap) siftDown(i int) {
	for h.left(i) < h.heapSize {
		l := h.left(i)
		r := h.right(i)
		u := l
		if r < h.heapSize && h.sl[r].priority > h.sl[l].priority {
			u = r
		}
		if h.sl[i].priority >= h.sl[u].priority {
			break
		}
		h.sl[i], h.sl[u] = h.sl[u], h.sl[i]
		i = u
	}
}

func (h *Heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) left(i int) int {
	return i*2 + 1
}

func (h *Heap) right(i int) int {
	return i*2 + 2
}

func (h *Heap) setSizeHeap() {
	h.heapSize = h.size()
}

func (h *Heap) Print() {
	for i := range h.sl {
		fmt.Print(h.sl[i].priority, " ")
	}
	fmt.Println()
}
