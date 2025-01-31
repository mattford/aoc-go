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

func MinInt(numbers []int) int {
	min := 0
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
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

func Without[T comparable](haystack []T, needle T) []T {
	out := make([]T, 0, len(haystack)-1)
	for _, i := range haystack {
		if i != needle {
			out = append(out, i)
		}
	}
	return out
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

func RemoveIndex[T comparable](haystack []T, needle int) []T {
	newArr := make([]T, 0, len(haystack)-1)
	for k, v := range haystack {
		if k != needle {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

func InsertAtIndex[T comparable](haystack []T, needle int, item T) []T {
	newArr := make([]T, 0, len(haystack)+1)
	newArr = append(newArr, haystack[0:needle]...)
	newArr = append(newArr, item)
	newArr = append(newArr, haystack[needle:]...)
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

func CopyMap[T string, V any](m map[T]V) map[T]V {
	newMap := make(map[T]V)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}
