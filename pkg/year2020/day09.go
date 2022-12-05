package year2020

import (
	"aocgen/pkg/common"
	"sort"
)

type Day09 struct{}

var preamble int = 25

func (p Day09) PartA(lines []string) any {
	return findInvalid(common.GetInts(lines))
}

func (p Day09) PartB(lines []string) any {
	numbers := common.GetInts(lines)
	invalid := findInvalid(numbers)
	for i := range numbers {
		check, score := check(numbers, i, invalid)
		if check {
			return score
		}
	}
	panic("Failed to find answer")
}

func check(numbers []int, idx int, target int) (bool, int) {
	acc := numbers[idx]
	group := make([]int, 1)
	group[0] = numbers[idx]
	idx2 := idx
	for acc < target {
		idx2++
		acc += numbers[idx2]
		group = append(group, numbers[idx2])
		if acc == target {
			return true, getScore(group)
		}
	}
	return false, 0
}

func getScore(group []int) int {
	sort.Sort(sort.IntSlice(group))
	return group[0] + group[len(group)-1]
}

func findInvalid(numbers []int) int {
	for idx := preamble; idx < len(numbers)-1; idx++ {
		prev := numbers[idx-preamble : idx]
		if !valid(prev, numbers[idx]) {
			return numbers[idx]
		}
	}
	panic("Can't find an invalid number")
}

func valid(prev []int, current int) bool {
	for i := 0; i < len(prev)-1; i++ {
		n1 := prev[i]
		for _, n2 := range prev[i+1:] {
			if n1+n2 == current {
				return true
			}
		}
	}
	return false
}
