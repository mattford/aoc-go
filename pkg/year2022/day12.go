package year2022

import (
	"aocgen/pkg/common"
	"strings"
)

type Day12 struct{}

func (p Day12) PartA(lines []string) any {
	grid, start, target := parseInput12(lines)
	dij := common.Dijkstra{
		Grid: grid,
		Pos:  start,
	}
	return dij.GetShortestPathCost(func(coord common.Coordinate2) bool {
		return coord == target
	}, func(coord common.Coordinate2) []common.Coordinate2 {
		neighbours := make([]common.Coordinate2, 0)
		for _, x := range common.Coordinate2Neighbours4 {
			neighbour := common.MoveBy2(coord, x)
			value, found := grid[neighbour]
			if found && grid[coord] >= value-1 {
				neighbours = append(neighbours, neighbour)
			}
		}
		return neighbours
	})
}

func (p Day12) PartB(lines []string) any {
	grid, _, target := parseInput12(lines)
	dij := common.Dijkstra{
		Grid: grid,
		Pos:  target,
	}
	return dij.GetShortestPathCost(func(coord common.Coordinate2) bool {
		return grid[coord] == 1
	}, func(coord common.Coordinate2) []common.Coordinate2 {
		neighbours := make([]common.Coordinate2, 0)
		for _, x := range common.Coordinate2Neighbours4 {
			neighbour := common.MoveBy2(coord, x)
			value, found := grid[neighbour]
			if found && value >= grid[coord]-1 {
				neighbours = append(neighbours, neighbour)
			}
		}
		return neighbours
	})
}

func parseInput12(lines []string) (grid map[common.Coordinate2]int, start common.Coordinate2, target common.Coordinate2) {
	grid = make(map[common.Coordinate2]int)
	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			coord := common.Coordinate2{Y: y, X: x}
			if char == "S" {
				grid[coord] = 1
				start = coord
			} else if char == "E" {
				grid[coord] = 26
				target = coord
			} else {
				grid[coord] = int(rune(char[0])) - 96
			}
		}
	}
	return
}
