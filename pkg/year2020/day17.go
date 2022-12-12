package year2020

import (
	"aocgen/pkg/common"
	"strings"
)

type Day17 struct{}

type gridState map[common.Coordinate4]bool

func (p Day17) PartA(lines []string) any {
	grid := parseInputDay17(lines)
	for i := 0; i < 6; i++ {
		grid = simulateCycle(grid, 3)
	}
	return len(grid)
}

func (p Day17) PartB(lines []string) any {
	grid := parseInputDay17(lines)
	for i := 0; i < 6; i++ {
		grid = simulateCycle(grid, 4)
	}
	return len(grid)
}

func simulateCycle(grid gridState, dims int) gridState {
	newGrid := make(gridState)
	visited := make([]common.Coordinate4, 0)
	var moves []common.Coordinate4
	if dims == 4 {
		moves = common.Coordinate4Neighbours80
	} else {
		moves = common.Coordinate4Neighbours26
	}
	for coord, _ := range grid {
		visited = append(visited, coord)
		neighbours := getNeighbours(coord, grid, moves)
		if neighbours == 3 || neighbours == 2 {
			newGrid[coord] = true
		}
		for _, move := range moves {
			neighbour := common.MoveBy4(coord, move)
			if common.Contains(visited, neighbour) {
				continue
			}
			neighbours2 := getNeighbours(neighbour, grid, moves)
			if neighbours2 == 3 {
				newGrid[neighbour] = true
			}
		}
	}
	return newGrid
}

func getNeighbours(coord common.Coordinate4, grid gridState, moves []common.Coordinate4) int {
	neighbours := 0
	for _, move := range moves {
		neighbour := common.MoveBy4(coord, move)
		if _, ok := grid[neighbour]; ok {
			neighbours++
			if neighbours > 3 {
				break
			}
		}
	}
	return neighbours
}

func parseInputDay17(lines []string) gridState {
	out := make(gridState)
	for i, line := range lines {
		chars := strings.Split(line, "")
		for j, char := range chars {
			if char == "#" {
				out[common.Coordinate4{i, j, 0, 0}] = true
			}
		}
	}
	return out
}
