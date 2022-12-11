package year2020

import (
	"aocgen/pkg/common"
	"fmt"
	"strings"
)

type Day17 struct{}

type gridState map[int]map[int]map[int]bool

func (p Day17) PartA(lines []string) any {
	grid := parseInputDay17(lines)
	fmt.Println(grid)
	newGrid := simulateCycle(grid)
	fmt.Println(newGrid)
	return "implement_me"
}

func (p Day17) PartB(lines []string) any {
	return "implement_me"
}

func simulateCycle(grid gridState) gridState {
	newGrid := make(gridState, len(grid))
	for y, row := range grid {
		newGrid[y] = make(map[int]map[int]bool)
		for x, col := range row {
			newGrid[y][x] = make(map[int]bool)
			for z, cell := range col {
				neighbours := getActiveNeighbours(y, x, z, grid)
				fmt.Println(y, x, z, cell, neighbours)
				newGrid[y][x][z] = neighbours == 3 || (cell && neighbours == 2)
			}
		}
	}
	return newGrid
}

func getActiveNeighbours(y int, x int, z int, grid gridState) int {
	total := 0
	checked := 0
	keys := common.Keys(grid)
	for y2 := keys[0] - 1; y2 <= keys[len(keys)-1]+1; y2++ {
		keysX := common.Keys(grid[y2])
		for x2 := keysX[0] - 1; x2 <= keysX[len(keysX)-1]+1; x2++ {
			keysZ := common.Keys(grid[y2][x2])
			for z2 := keysZ[0] - 1; z2 <= keysZ[len(keysZ)-1]+1; z2++ {
				if x == x2 && y == y2 && z == z2 {
					continue
				}
				checked++
				if grid[y2][x2][z2] {
					fmt.Println(y2, x2, z2, "is active")
					total++
				}
			}
		}
	}
	fmt.Println(checked, "cells checked")
	return total
}

func parseInputDay17(lines []string) gridState {
	out := make(gridState, len(lines))
	for i, line := range lines {
		chars := strings.Split(line, "")
		out[i] = make(map[int]map[int]bool, len(chars))
		for j, char := range chars {
			out[i][j] = map[int]bool{0: char == "#"}
		}
	}
	return out
}
