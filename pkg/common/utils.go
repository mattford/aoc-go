package common

import (
	"strconv"
)

func GetInts(lines []string) []int {
	ints := make([]int, len(lines))
	for i, str := range lines {
		n, err := strconv.Atoi(str)
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

func Contains[T comparable](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
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

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
