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

	fmt.Println(sorting.Fib(0))
	fmt.Println(sorting.FibLine(0))
}
