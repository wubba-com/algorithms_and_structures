package main

import (
	"fmt"
	ll "github.com/wubba-com/algorithms_and_structures/structure/linked_list"
	"github.com/wubba-com/algorithms_and_structures/structure/queue"
	"github.com/wubba-com/algorithms_and_structures/structure/stack"
	"github.com/wubba-com/algorithms_and_structures/structure/tree"
)

func main() {
	// -------- односвязный список + стек ---------
	var list = ll.New()
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	// last element
	fmt.Println(list.Find(1))

	s := stack.New()
	s.LLPush(1)

	lastNodeForRead := s.LLPeek()                // get last element
	lastNodeDeleteAndReturning, err := s.LLPop() // delete and get last element
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(lastNodeDeleteAndReturning.Value() == lastNodeForRead.Value())
	stack.PrintStack(s)

	// -------- двусвязный список + очередь ---------
	q := queue.New()

	for i := 1; i < 3; i++ {
		q.Unshift(i)
		q.Push(i * 10)
	}
	for i := 3; i < 5; i++ {
		q.Unshift(i)
		q.Push(i * 10)
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
}

func Do(value int) {
	fmt.Print(value)
}
