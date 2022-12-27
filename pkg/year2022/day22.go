package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Day22 struct{}

func (p Day22) PartA(lines []string) any {
	levelMap, instructions, sideLength := parseInput22(lines)
	fmt.Println(sideLength)
	pos, _ := firstLastInRow(levelMap, 0)
	facing := 0
	for _, instr := range instructions {
		x, err := strconv.Atoi(instr)
		if err == nil {
			pos = move(levelMap, pos, facing, x)
		} else {
			facing = turn(facing, rune(instr[0]))
		}
	}
	return score(pos, facing)
}

func (p Day22) PartB(lines []string) any {
	return "implement_me"
}

func score(pos common.Coordinate2, facing int) (score int) {
	score += 1000 * (pos.Y + 1)
	score += 4 * (pos.X + 1)
	score += facing
	return
}

func firstLastInRow(grid map[common.Coordinate2]int, row int) (first, last common.Coordinate2) {
	minX, maxX, _, _ := common.Bounds(common.Keys(grid))
	gotFirst := false
	for x := minX; x < maxX; x++ {
		_, ok := grid[common.Coordinate2{Y: row, X: x}]
		if ok {
			if !gotFirst {
				gotFirst = true
				first = common.Coordinate2{Y: row, X: x}
			}
			last = common.Coordinate2{Y: row, X: x}
		}
	}
	return
}

func firstLastInCol(grid map[common.Coordinate2]int, col int) (first, last common.Coordinate2) {
	_, _, minY, maxY := common.Bounds(common.Keys(grid))
	gotFirst := false
	for y := minY; y < maxY; y++ {
		_, ok := grid[common.Coordinate2{Y: y, X: col}]
		if ok {
			if !gotFirst {
				gotFirst = true
				first = common.Coordinate2{Y: y, X: col}
			}
			last = common.Coordinate2{Y: y, X: col}
		}
	}
	return
}

func move(grid map[common.Coordinate2]int, pos common.Coordinate2, facing int, amount int) common.Coordinate2 {
	var vector common.Coordinate2
	switch facing {
	case 3:
		vector = common.Coordinate2{Y: -1, X: 0}
	case 1:
		vector = common.Coordinate2{Y: 1, X: 0}
	case 2:
		vector = common.Coordinate2{Y: 0, X: -1}
	case 0:
		vector = common.Coordinate2{Y: 0, X: 1}
	}
	for i := 0; i < amount; i++ {
		nextPos := common.MoveBy2(pos, vector)
		v, ok := grid[nextPos]
		if !ok {
			switch facing {
			case 3:
				_, nextPos = firstLastInCol(grid, nextPos.X)
			case 1:
				nextPos, _ = firstLastInCol(grid, nextPos.X)
			case 2:
				_, nextPos = firstLastInRow(grid, nextPos.Y)
			case 0:
				nextPos, _ = firstLastInRow(grid, nextPos.Y)
			}
			if v2 := grid[nextPos]; v2 == 1 {
				break
			}
		} else if v == 1 {
			break
		}
		pos = nextPos
	}
	return pos
}

func turn(current int, direction rune) int {
	if direction == 'L' {
		return (current - 1) / 4
	}
	return (current + 1) % 4
}

func parseInput22(lines []string) (map[common.Coordinate2]int, []string, int) {
	levelMap := make(map[common.Coordinate2]int)
	instructions := make([]string, 0)
	mapFinished := false
	for y, line := range lines {
		if line == "" {
			mapFinished = true
			continue
		}
		if !mapFinished {
			chars := []rune(line)
			for x, char := range chars {
				switch char {
				case '#':
					levelMap[common.Coordinate2{Y: y, X: x}] = 1
				case '.':
					levelMap[common.Coordinate2{Y: y, X: x}] = 0
				}
			}
		} else {
			expr := regexp.MustCompile("([A-Z]+|[0-9]+)(.*)")
			for len(line) > 0 {
				matches := expr.FindStringSubmatch(line)
				instructions = append(instructions, matches[1])
				line = matches[2]
			}
		}
	}
	_, maxX, _, maxY := common.Bounds(common.Keys(levelMap))
	sideLength := (int(math.Max(float64(maxX), float64(maxY))) + 1) / 4
	return levelMap, instructions, sideLength
}
