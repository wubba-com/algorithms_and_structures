package dll

type Node struct {
	prev *Node
	next *Node
	val  interface{}
}

func (n *Node) setNext(next *Node) {
	n.next = next
}

func (n *Node) setPrev(prev *Node) {
	n.prev = prev
}

func (n Node) Prev() *Node {
	return n.prev
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Value() interface{} {
	return n.val
}
