package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day14 struct{}

func (p Day14) PartA(lines []string) any {
	grid := buildGrid(lines)
	sandEntersFrom := common.Coordinate2{0, 500}
	enteredVoid := false
	for !enteredVoid {
		grid, enteredVoid = simulateSand(grid, sandEntersFrom, 0)
	}
	//printGrid(grid)
	return countGrid(grid, 'o')
}

func (p Day14) PartB(lines []string) any {
	grid := buildGrid(lines)
	maxY := math.MinInt
	for coord := range grid {
		maxY = int(math.Max(float64(maxY), float64(coord.Y)))
	}
	sandEntersFrom := common.Coordinate2{0, 500}
	for {
		grid, _ = simulateSand(grid, sandEntersFrom, maxY+2)
		_, ok := grid[sandEntersFrom]
		if ok {
			break
		}
	}
	//printGrid(grid)
	return countGrid(grid, 'o')
}

func countGrid(grid map[common.Coordinate2]rune, target rune) int {
	total := 0
	for _, r := range grid {
		if r == target {
			total++
		}
	}
	return total
}

func simulateSand(grid map[common.Coordinate2]rune, sandPos common.Coordinate2, floor int) (map[common.Coordinate2]rune, bool) {
	moves := []common.Coordinate2{
		common.South,
		common.SouthWest,
		common.SouthEast,
	}
	voidPos := math.MinInt
	for pos := range grid {
		voidPos = int(math.Max(float64(voidPos), float64(pos.Y)))
	}
outer:
	for {
		if floor == 0 || sandPos.Y < floor-1 {
			for _, move := range moves {
				moved := common.MoveBy2(sandPos, move)
				_, ok := grid[moved]
				if !ok {
					sandPos = moved
					if floor == 0 && sandPos.Y > voidPos {
						return grid, true
					}
					continue outer
				}
			}
		}
		grid[sandPos] = 'o'
		break
	}
	return grid, false
}

func printGrid(grid map[common.Coordinate2]rune) {
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt
	for coord := range grid {
		minX = int(math.Min(float64(minX), float64(coord.X)))
		minY = int(math.Min(float64(minY), float64(coord.Y)))
		maxX = int(math.Max(float64(maxX), float64(coord.X)))
		maxY = int(math.Max(float64(maxY), float64(coord.Y)))
	}
	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			v, ok := grid[common.Coordinate2{y, x}]
			if ok {
				fmt.Print(string(v))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func buildGrid(lines []string) map[common.Coordinate2]rune {
	grid := make(map[common.Coordinate2]rune)
	for _, line := range lines {
		pairs := strings.Split(line, " -> ")
		var previousCoord common.Coordinate2
		for i, pair := range pairs {
			numbers := strings.Split(pair, ",")
			x, _ := strconv.Atoi(numbers[0])
			y, _ := strconv.Atoi(numbers[1])
			thisCoord := common.Coordinate2{y, x}
			grid[thisCoord] = '#'
			if i > 0 {
				grid = drawLine(grid, previousCoord, thisCoord)
			}
			previousCoord = thisCoord
		}
	}
	return grid
}

func drawLine(grid map[common.Coordinate2]rune, from common.Coordinate2, to common.Coordinate2) map[common.Coordinate2]rune {
	if from == to {
		return grid
	}
	if from.Y != to.Y {
		min := int(math.Min(float64(from.Y), float64(to.Y)))
		max := int(math.Max(float64(from.Y), float64(to.Y)))
		for yl := min; yl <= max; yl++ {
			grid[common.Coordinate2{yl, from.X}] = '#'
		}
	}
	if from.X != to.X {
		min := int(math.Min(float64(from.X), float64(to.X)))
		max := int(math.Max(float64(from.X), float64(to.X)))
		for xl := min; xl <= max; xl++ {
			grid[common.Coordinate2{from.Y, xl}] = '#'
		}
	}
	return grid
}
