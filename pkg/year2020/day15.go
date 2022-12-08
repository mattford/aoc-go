package year2020

import (
	"strconv"
	"strings"
)

type Day15 struct{}

func (p Day15) PartA(lines []string) any {
	return getNthNumber(lines, 2020)
}

func (p Day15) PartB(lines []string) any {
	return getNthNumber(lines, 30000000)
}

func getNthNumber(lines []string, n int) int {
	ints := getNumbers(lines)
	lastSaid := make(map[int]int)
	var previousNumber int
	for i, v := range ints {
		previousNumber = v
		lastSaid[previousNumber] = i + 1
	}
	delete(lastSaid, previousNumber)

	for turn := len(ints) + 1; turn <= n; turn++ {
		lastTurn, ok := lastSaid[previousNumber]
		lastSaid[previousNumber] = turn - 1
		if !ok || lastTurn == turn-1 {
			previousNumber = 0
		} else {
			previousNumber = (turn - 1) - lastTurn
		}
	}
	return previousNumber
}

func getNumbers(lines []string) []int {
	strs := strings.Split(lines[0], ",")
	ints := make([]int, 0, len(strs))
	for _, s := range strs {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}
