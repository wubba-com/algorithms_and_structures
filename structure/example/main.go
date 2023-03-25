package main

import (
	"fmt"
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
