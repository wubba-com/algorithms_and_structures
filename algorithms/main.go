package main

import (
	"fmt"
	"github.com/wubba-com/algorithms_and_structures/algorithms/sorting"
)

func main() {
	// -------- сортировка + поиск ---------
	sl := []int{2, 1, 4, 56, 2, 98, 10}
	sl = sorting.Sort(sl)
	fmt.Println(sl)
	i := sorting.BinarySearch(sl, 56)
	fmt.Println(sl[i] == 56)

	fmt.Println(sorting.Fib(4))
	fmt.Println(sorting.FibLine(4))

	sl = []int{2, 0, 4, 0, 6, 0, 3, 0, 1, 0, 1, 0, 1}
	fmt.Println(len(sl))
	sorting.MoveElementToEnd(sl, 0)
	fmt.Println(len(sl), sl)

	sorting.SortSelected(sl)
	fmt.Println(sl)

	fmt.Println("ciphertext size is less than nonceSize")
	p("ciphertext size is less than nonceSize")

}

func p(v any) {
	panic(v)
}
