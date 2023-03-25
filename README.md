######  Реализация базовых структур данных и алгоритмов на Go :rocket:

### Cтруктуры данных :open_book:: 
stack, queue, linked list, double linked list, binary tree 

### Aлгоритмы :blue_book:: 
binary search, quick sort, sort insert . 

Подходит для ознакомления новичкам в Go. 

Пример структуры данных stack:
```go

	s := stack.NewBasicLinkedList()

    	for i := 1; i < 10; i++ {
		s.Push(stack.Value{Value: i})
	}
	
	// пока список не пустой
	for s.Next() {
		var v *stack.Value
		v, err = s.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v ", v.Value)
	}
 ```

 Пример структуры данных queue:
```go
	q := queue.New()

	for i := 1; i < 3; i++ {
		q.PushFront(i)
		q.PushBack(i * 10)
	}
	// вывести все элементы
	queue.Print(q) // 4 3 2 1 10 20 30 40

	el, err := q.PopBack()
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Printf("%v", el.Value()) // 40

	el, err = q.PopFront()
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Printf("%v", el.Value()) // 4

	queue.Print(q) // 3 2 1 10 20 30
 ```
P.S Проект будет дополнятся и более сложными структурами и алгоритмами
