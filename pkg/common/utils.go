package common

import (
	"strconv"
	"strings"
)

func GetInts(lines []string) []int {
	ints := make([]int, len(lines))
	for i, str := range lines {
		n, err := strconv.Atoi(strings.Trim(str, " "))
		if err != nil {
			panic(err)
		}
		ints[i] = n
	}
	return ints
}

func MaxInt(numbers []int) int {
	max := 0
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

func Without(ints []int, target int) []int {
	newInts := make([]int, 0, len(ints)-1)
	for _, i := range ints {
		if i != target {
			newInts = append(newInts, i)
		}
	}
	return newInts
}

func Unique[T comparable](a []T) []T {
	newA := make([]T, 0)
	for _, y := range a {
		if !Contains(newA, y) {
			newA = append(newA, y)
		}
	}
	return newA
}

func Keys[T comparable, V any](haystack map[T]V) []T {
	keys := make([]T, 0, len(haystack))
	for k := range haystack {
		keys = append(keys, k)
	}
	return keys
}

func Contains[T comparable](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func Remove[T comparable](haystack []T, needle T) []T {
	newArr := make([]T, 0)
	for _, v := range haystack {
		if v != needle {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

func Column(haystack [][]int, idx int) []int {
	out := make([]int, 0, len(haystack))
	for _, v := range haystack {
		out = append(out, v[idx])
	}
	return out
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Sum(ints []int) int {
	total := 0
	for _, i := range ints {
		total += i
	}
	return total
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Bound(a int, low int, high int) int {
	if a > high {
		return high
	}
	if a < low {
		return low
	}
	return a
}

func Pop(stack *[]any) any {
	if len(*stack) == 0 {
		return 0
	}
	item := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return item
}
