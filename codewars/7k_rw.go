package main

import (
	"math"
	"math/bits"
	"sort"
	"strconv"
	"strings"
)

func ReverseWords(str string) string {
	var s string
	for _, v := range strings.Split(str, " ") {
		bs := []byte(v)
		for k, j := 0, len(bs)-1; k < j; k, j = k+1, j-1 {
			bs[k], bs[j] = bs[j], bs[k]
		}
		s += string(bs) + " "
	}

	return strings.Trim(s, " ") // reverse those words
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	var i float64 = 2
	for i <= math.Floor(math.Sqrt(float64(n))) {
		if n%int(i) == 0 {
			return false
		}
		i++
	}
	return true
}

func duplicateCount(s1 string) int {
	m := make(map[rune]int)
	var count int
	for _, r := range strings.ToLower(s1) {
		m[r]++
		if m[r] == 2 {
			count++
		}
	}

	return count
}

func FindOdd(seq []int) (odd int) {
	if len(seq) < 2 {
		return 1
	}
	sort.Ints(seq)
	var count = 1
	for i, _ := range seq {
		if i != len(seq)-1 {
			if seq[i] != seq[i+1] {
				if count%2 != 0 {
					return seq[i]
				}
				count = 1
				continue
			}
			count++
		} else {
			if count == 1 {
				return seq[i]
			}
		}
	}

	return
}

func CountBits(x uint) int {
	var bits int
	for x != 0 {
		if x&1 != 0 {
			bits++
		}
		x >>= 1
	}

	return bits
}

func CountBitsV2(n uint) int {
	return bits.OnesCount(n)
}

func TwoSum(numbers []int, target int) [2]int {
	sum := func(a, b int) int {
		return a + b
	}
	for i := range numbers {
		for j := range numbers {
			if i != j {
				if sum(numbers[i], numbers[j]) == target {
					return [2]int{i, j}
				}
			}
		}
	}
	return [2]int{}
}

func DigPow(n, p int) int {
	nums := make([]int, 0)
	tmp := n
	for tmp != 0 {
		nums = append(nums, tmp%10)
		tmp /= 10
	}
	var sum int
	for i := len(nums) - 1; i >= 0; i-- {
		sum += int(math.Pow(float64(nums[i]), float64(p)))
		p++
	}

	if sum%n == 0 {
		return sum / n
	}
	return -1
}

func OrderWeight(strng string) string {
	blocks := strings.Fields(strng)

	var (
		matrix = make([][2]int, 0)

		sum = func(sl []int) int {
			var s int
			for i := range sl {
				s += sl[i]
			}

			return s
		}
	)
	for i := range blocks {
		n, err := strconv.Atoi(blocks[i])
		if err != nil {
			continue
		}
		tmp := n
		row := make([]int, 0)
		for tmp != 0 {
			row = append(row, tmp%10)
			tmp /= 10
		}
		matrix = append(matrix, [2]int{n, sum(row)})
	}

	sort.Slice(matrix, func(i, j int) bool {
		if matrix[i][1] == matrix[j][1] {
			if strconv.Itoa(matrix[i][0]) < strconv.Itoa(matrix[j][0]) {
				return true
			}
			return false
		}
		return matrix[i][1] < matrix[j][1]
	})
	var nums = make([]string, len(matrix))
	for i := range matrix {
		nums[i] = strconv.Itoa(matrix[i][0])
	}

	return strings.Join(nums, " ")
}

func FindUniq(arr []float32) float32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	if arr[0] == arr[1] {
		return arr[len(arr)-1]
	}
	return arr[0]
}
