package year2022

import (
	"strconv"
	"strings"
)

type Day04 struct{}

func (p Day04) PartA(lines []string) any {
	elfPairs := getPairs(lines)
	return findCompleteOverlaps(elfPairs)
}

func (p Day04) PartB(lines []string) any {
	elfPairs := getPairs(lines)
	return findPartialOverlaps(elfPairs)
}

func findPartialOverlaps(elfPairs [][][]int) int {
	total := 0
	for _, pair := range elfPairs {
		minA, maxA, minB, maxB := pair[0][0], pair[0][1], pair[1][0], pair[1][1]
		if (minA >= minB && minA <= maxB) || (maxA <= maxB && maxA >= minB) ||
			(minB >= minA && minB <= maxA) || (maxB <= maxA && maxB >= minA) {
			total++
		}
	}
	return total
}

func findCompleteOverlaps(elfPairs [][][]int) int {
	total := 0
	for _, pair := range elfPairs {
		minA, maxA, minB, maxB := pair[0][0], pair[0][1], pair[1][0], pair[1][1]
		if (minA <= minB && maxA >= maxB) || (minB <= minA && maxB >= maxA) {
			total++
		}
	}
	return total
}

func getPairs(lines []string) [][][]int {
	elfPairs := make([][][]int, 0, len(lines))
	for _, line := range lines {
		elfPairs = append(elfPairs, getRanges(line))
	}
	return elfPairs
}

func getRanges(line string) [][]int {
	ranges := make([][]int, 2)
	elves := strings.Split(line, ",")
	for i, elf := range elves {
		parts := strings.Split(elf, "-")
		min, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		ranges[i] = []int{min, max}
	}
	return ranges
}
