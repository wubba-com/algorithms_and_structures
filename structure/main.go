package main

import (
	"fmt"
	"github.com/wubba-com/algorithms_and_structures/structure/graph"
	"github.com/wubba-com/algorithms_and_structures/structure/heap"
	ll "github.com/wubba-com/algorithms_and_structures/structure/linked_list"
	"github.com/wubba-com/algorithms_and_structures/structure/queue"
	"github.com/wubba-com/algorithms_and_structures/structure/stack"
	"github.com/wubba-com/algorithms_and_structures/structure/tree"
)

func main() {
	// -------- односвязный список + стек ---------
	var list = ll.New()
	for i := 1; i <= 5; i++ {
		list.Add(i)
	}

	// last element
	fmt.Println(list.Find(1))

	s := stack.NewBasicLinkedList()
	s.Push(stack.Value{Value: 1})

	lastNodeForRead, ok := s.Peek() // get last element
	if ok {
		// element is
	}
	lastNodeDeleteAndReturning, err := s.Pop() // delete and get last element
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(lastNodeDeleteAndReturning.Value.(int) == lastNodeForRead.Value.(int))

	for i := 1; i < 10; i++ {
		s.Push(stack.Value{Value: i})
	}
	s.Print()
	_, _ = s.Pop()
	s.Print()

	for s.Next() {
		var v *stack.Value
		v, err = s.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v ", v.Value)
	}
	fmt.Println()
	// -------- двусвязный список + очередь ---------
	q := queue.New()

	for i := 1; i < 3; i++ {
		q.PushFront(i)
		q.PushBack(i * 10)
	}
	for i := 3; i < 5; i++ {
		q.PushFront(i)
		q.PushBack(i * 10)
	}
	// вывести все элементы
	queue.Print(q)

	// бинарное дерево
	t := tree.New(4)

	if err = t.Add(2); err != nil {
		fmt.Println(err)
	}

	if err = t.Add(5); err != nil {
		fmt.Println(err)
	}

	if err = t.Add(3); err != nil {
		fmt.Println(err)
	}

	if err = t.Add(1); err != nil {
		fmt.Println(err)
	}

	if err = t.Add(7); err != nil {
		fmt.Println(err)
	}

	if err = t.Add(6); err != nil {
		fmt.Println(err)
	}

	if err = t.Add(8); err != nil {
		fmt.Println(err)
	}

	t.Inorder(Do)
	fmt.Println()
	t.Preorder(Do)
	fmt.Println()
	t.Postorder(Do)
	fmt.Println()
	t.Delete(3)
	fmt.Println(t.Find(3))

	// граф              (7)             (1)
	//         B (1) --------- F (5) ---------- G (6)
	//       /               /    |  \
	//     /  (2)          / (2)  |   \ (4)
	//   /               /        |    \
	// A (0)           D (3)      | (1)  S (7)
	//  \            /            |    /
	//   \ (1)     / (5)          |  / (6)
	//    \      /     (2)        |/
	//     C (2) ------------- E (4)
	// матрица смежности графа:

	//   0 1 2 3 4 5 6
	// 0 x 1 1 x x x x
	// 1 x x x x 1 x x
	// 2 x x x 1 x 1 x
	// 3 x x x x 1 x x
	// 4 x x x x x x 1
	// 5 x x x x x x 1
	// 6 x x x x x x x
	g := graph.New()
	idA := g.AddNode()
	idB := g.AddNode()
	idC := g.AddNode()
	idD := g.AddNode()
	idE := g.AddNode()
	idF := g.AddNode()
	idG := g.AddNode()
	idS := g.AddNode()

	err = g.AddEdge(idA, idB, 2)
	err = g.AddEdge(idB, idA, 2)
	err = g.AddEdge(idA, idC, 1)
	err = g.AddEdge(idB, idF, 7)
	err = g.AddEdge(idC, idD, 5)
	err = g.AddEdge(idC, idE, 2)
	err = g.AddEdge(idD, idF, 2)
	err = g.AddEdge(idE, idF, 1)
	err = g.AddEdge(idF, idG, 1)
	err = g.AddEdge(idF, idS, 4)
	err = g.AddEdge(idE, idS, 6)

	exist := g.BreadthSearch(idA, idS)
	fmt.Println("found:", exist)

	costs, path := g.Short(idB, idS)
	fmt.Println("costs:", costs)
	fmt.Println("path:", path)

	h := heap.NewMax()
	h.Push(10, nil)
	h.Push(20, nil)
	h.Push(30, nil)
	h.Push(40, nil)
	h.Push(50, nil)
	h.Push(60, nil)
	//h.Push(1, nil)
	//h.Push(3, nil)
	h.Print()
	h.Pop()
	h.Print()
}

func Do(value int) {
	fmt.Print(value)
}
