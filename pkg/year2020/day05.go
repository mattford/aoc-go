package year2020

import (
	"math"
	"sort"
	"strings"
)

type Day05 struct{}

func (p Day05) PartA(lines []string) any {
	rows := 128
	cols := 8

	highest := 0
	for _, line := range lines {
		_, _, seatId := getSeatId(line, rows, cols)
		if seatId > highest {
			highest = seatId
		}
	}

	return highest
}

func (p Day05) PartB(lines []string) any {
	rows := 128
	cols := 8

	seatIds := make([]int, len(lines))
	for _, line := range lines {
		_, _, seatId := getSeatId(line, rows, cols)
		seatIds = append(seatIds, seatId)
	}
	sort.Sort(sort.IntSlice(seatIds))
	last := 0
	for _, id := range seatIds {
		if id > 0 && last > 0 && id != last+1 {
			return id - 1
		}
		last = id
	}

	return "didn't find an answer"
}

func getSeatId(line string, rows int, cols int) (int, int, int) {
	minRow, maxRow := 0, rows-1
	letters := strings.Split(line, "")

	for _, letter := range letters[:7] {
		switch letter {
		case "F":
			maxRow -= int(math.Ceil(float64(maxRow-minRow) / 2))
		case "B":
			minRow += int(math.Ceil(float64(maxRow-minRow) / 2))
		}
	}

	minCol, maxCol := 0, cols-1
	for _, letter := range letters[7:] {
		switch letter {
		case "L":
			maxCol -= int(math.Ceil(float64(maxCol-minCol) / 2))
		case "R":
			minCol += int(math.Ceil(float64(maxCol-minCol) / 2))
		}
	}
	return minRow, minCol, (minRow * 8) + minCol
}
