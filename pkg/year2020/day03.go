package year2020

import (
	"strings"
)

type Day03 struct{}

type Slope struct {
	down  int
	right int
}

func (p Day03) PartA(lines []string) any {
	grid := buildGrid(lines)
	return howManyTrees(grid, 1, 3)
}

func (p Day03) PartB(lines []string) any {
	slopes := []Slope{
		{down: 1, right: 1},
		{down: 1, right: 3},
		{down: 1, right: 5},
		{down: 1, right: 7},
		{down: 2, right: 1},
	}
	grid := buildGrid(lines)
	total := 1
	for _, slope := range slopes {
		total *= howManyTrees(grid, slope.down, slope.right)
	}
	return total
}

func howManyTrees(grid [][]int, down int, right int) (count int) {
	x, y, count := 0, 0, 0
	for y < len(grid) {
		if x >= len(grid[y]) {
			x -= len(grid[y])
		}

		count += grid[y][x]
		y += down
		x += right
	}
	return
}

func buildGrid(lines []string) (grid [][]int) {
	grid = make([][]int, len(lines))
	for i, line := range lines {
		cells := strings.Split(line, "")
		grid[i] = make([]int, len(cells))
		for j, value := range cells {
			if value == "#" {
				grid[i][j] = 1
			}
		}
	}
	return
}
