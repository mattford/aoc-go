package year2022

import (
	"aocgen/pkg/common"
	"strconv"
	"strings"
)

type Day08 struct{}

func (p Day08) PartA(lines []string) any {
	grid := buildTreeGrid(lines)
	total := (len(grid)+len(grid[0]))*2 - 4
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if isTreeVisible(y, x, grid) {
				total++
			}
		}
	}
	return total
}

func (p Day08) PartB(lines []string) any {
	grid := buildTreeGrid(lines)
	bestScore := 0
	for y := 0; y < len(grid)-1; y++ {
		for x := 0; x < len(grid[y])-1; x++ {
			visibilityScore := getVisibilityScore(y, x, grid)
			if visibilityScore > bestScore {
				bestScore = visibilityScore
			}
		}
	}
	return bestScore
}

func getVisibilityScore(y int, x int, grid [][]int) int {
	myTreeHeight := grid[y][x]
	// scan vertical
	col := common.Column(grid, x)
	return getDistance(myTreeHeight, col[y+1:], false) *
		getDistance(myTreeHeight, col[:y], true) *
		getDistance(myTreeHeight, grid[y][x+1:], false) *
		getDistance(myTreeHeight, grid[y][:x], true)

}

func isTreeVisible(y int, x int, grid [][]int) bool {
	myTreeHeight := grid[y][x]
	// scan vertical
	col := common.Column(grid, x)
	return !isBlockedInAxis(myTreeHeight, col[:y]) ||
		!isBlockedInAxis(myTreeHeight, col[y+1:]) ||
		!isBlockedInAxis(myTreeHeight, grid[y][:x]) ||
		!isBlockedInAxis(myTreeHeight, grid[y][x+1:])
}

func isBlockedInAxis(treeHeight int, trees []int) bool {
	for _, otherHeight := range trees {
		if otherHeight >= treeHeight {
			return true
		}
	}
	return false
}

func getDistance(treeHeight int, trees []int, reverse bool) int {
	count := 0
	if reverse {
		for i := len(trees) - 1; i >= 0; i-- {
			count++
			if trees[i] >= treeHeight {
				return count
			}
		}
	} else {
		for i := 0; i < len(trees); i++ {
			count++
			if trees[i] >= treeHeight {
				return count
			}
		}
	}
	return count
}

func buildTreeGrid(lines []string) [][]int {
	rows := make([][]int, 0, len(lines))
	for _, line := range lines {
		chars := strings.Split(line, "")
		cols := make([]int, 0, len(chars))
		for _, char := range chars {
			n, _ := strconv.Atoi(char)
			cols = append(cols, n)
		}
		rows = append(rows, cols)
	}
	return rows
}
