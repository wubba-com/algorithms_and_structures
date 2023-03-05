package sorting

// Fib классический алгоритм через рекурсию (больше используемой памяти)
func Fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fib(n-2) + Fib(n-1)
}

// FibLine оптимизированный алгоритм (меньше используемой памяти)
func FibLine(n int) int {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	x, y := 1, 1
	for i := 2; i < n; i++ {
		x, y = y, x+y
	}

	return y
}

func Includes(sl []int, include int) bool {
	for _, v := range sl {
		if v == include {
			return true
		}
	}

	return false
}

// BinarySearch - O(log n)
func BinarySearch(sl []int, value int) int {
	if sl == nil {
		return -1
	}

	start := 0
	end := len(sl) - 1

	for start <= end {
		mid := (start + end) / 2
		switch {
		case sl[mid] == value:
			return mid
		case sl[mid] < value:
			start = mid + 1
		default:
			end = mid - 1
		}
	}

	return -1
}

// SortSelected Сложность O(n**2)
func SortSelected(sl []int) {
	if len(sl) < 2 {
		return
	}

	for i := 0; i < len(sl); i++ {
		min := i
		for j := i + 1; j < len(sl); j++ {
			if sl[j] < sl[min] {
				min = j
			}
		}
		if sl[min] != sl[i] {
			sl[i], sl[min] = sl[min], sl[i]
		}
	}
}

// SortInsert Сортировка вставками
// Сложность 	Наилучший случай 	В среднем 	Наихудший случай
// Время 		O(n) 				O(n2) 		O(n2)
// Память 		O(1) 				O(1) 		O(1)
func SortInsert(sl []int) {
	for index := 1; index < len(sl); index++ {
		currentValue := sl[index]
		currentIndex := index
		for ; currentIndex >= 1 && sl[currentIndex-1] > currentValue; currentIndex-- {
			sl[currentIndex] = sl[currentIndex-1]
		}
		sl[currentIndex] = currentValue
	}
}

// MergeSort Сортировка слиянием
// Сложность 	Наилучший случай 	В среднем 	Наихудший случай
// Время 		O(n·logN) 		O(n·logN) 	O(n·logN)
// Память 		O(n) 				O(n) 		O(n)
// Метод сортировки слиянием, потому что он использует рекурсию, поэтому вам нужно передать начальную позицию, конец и буфер, этот буфер имеет тот же размер, что и данные
func MergeSort(input []int) []int {
	l := len(input)
	if l == 1 {
		// валидация длинны
		return input
	}
	// разделяем набор на две части
	middleIdx := l / 2
	left := input[:middleIdx]
	right := input[middleIdx:]

	// рекурсивное слияние
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
	// создаем новый слайс путем сравнения по ключам
	for (len(left) > 0) && (len(right) > 0) {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
	}
	return result
}

// Quicksort - Быстрая сортировка
// Сложность 	Наилучший случай 	В среднем 	Наихудший случай
// Время 		O(n·logN) 			O(n·logN) 	O(n2)
// Память 		O(1) 				O(1) 		O(1)

// Sort Быстрая сортировка
// Данная реализация взята из книги "Грокаем Алгоритмы"
func Sort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	less, bigger, pivot := make([]int, 0), make([]int, 0), input[0]

	for _, v := range input[1:] {
		if v > pivot {
			bigger = append(bigger, v)
		} else {
			less = append(less, v)
		}
	}
	input = append(Sort(less), pivot)
	input = append(input, Sort(bigger)...)

	return input
}

// Quicksort Быстрая сортировка — это алгоритм типа «разделяй и властвуй»
func Quicksort(sl []int) {
	quickSort(sl, 0, len(sl)-1)
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

// MoveElementToEnd 1. Реализуйте алгоритм, который принимает массив и перемещает все нули в конец, сохраняя порядок остальных элементов.
func MoveElementToEnd(sl []int, num int) {
	if len(sl) < 2 {
		return
	}
	nums := make([]int, 0)

	for i, v := range sl {
		if v == num {
			if i+1 <= len(sl)-1 {
				nums = append(nums, v)
				copy(sl[i:], sl[i+1:])
				sl = sl[:len(sl)-1]
			}
		}
	}
	sl = append(sl, nums...)
}

// Посчитайте сумму ряда н-го xпирамида нечетных числе (начало с 1)
