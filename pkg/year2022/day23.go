package year2022

import (
	"aocgen/pkg/common"
	"strings"
)

type Day23 struct{}

type dir23 int

const north dir23 = 0
const east dir23 = 1
const south dir23 = 2
const west dir23 = 3

func (p Day23) PartA(lines []string) any {
	moves := []dir23{north, south, west, east}
	elves := getElves23(lines)
	for i := 0; i < 10; i++ {
		elves, moves, _ = doRound23(elves, moves)
	}
	return countEmpty(elves)
}

func (p Day23) PartB(lines []string) any {
	moves := []dir23{north, south, west, east}
	elves := getElves23(lines)
	movedElves := 1
	round := 0
	for movedElves > 0 {
		round++
		elves, moves, movedElves = doRound23(elves, moves)
	}
	return round
}

func countEmpty(elves map[common.Coordinate2]int) int {
	minX, maxX, minY, maxY := common.Bounds(common.Keys(elves))
	total := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, ok := elves[common.Coordinate2{Y: y, X: x}]; !ok {
				total++
			}
		}
	}
	return total
}

func canMove(grid map[common.Coordinate2]int, elf common.Coordinate2, move dir23) bool {
	var checkPos []common.Coordinate2
	switch move {
	case north:
		checkPos = []common.Coordinate2{
			common.MoveBy2(elf, common.NorthWest),
			common.MoveBy2(elf, common.North),
			common.MoveBy2(elf, common.NorthEast),
		}
	case east:
		checkPos = []common.Coordinate2{
			common.MoveBy2(elf, common.NorthEast),
			common.MoveBy2(elf, common.East),
			common.MoveBy2(elf, common.SouthEast),
		}
	case south:
		checkPos = []common.Coordinate2{
			common.MoveBy2(elf, common.SouthWest),
			common.MoveBy2(elf, common.South),
			common.MoveBy2(elf, common.SouthEast),
		}
	case west:
		checkPos = []common.Coordinate2{
			common.MoveBy2(elf, common.NorthWest),
			common.MoveBy2(elf, common.West),
			common.MoveBy2(elf, common.SouthWest),
		}
	}
	for _, pos := range checkPos {
		if _, ok := grid[pos]; ok {
			return false
		}
	}
	return true
}

func move23(elf common.Coordinate2, move dir23) common.Coordinate2 {
	switch move {
	case north:
		return common.MoveBy2(elf, common.North)
	case east:
		return common.MoveBy2(elf, common.East)
	case south:
		return common.MoveBy2(elf, common.South)
	case west:
		return common.MoveBy2(elf, common.West)
	}
	return elf
}

func rotateMoves(moves []dir23) []dir23 {
	newSlice := moves[1:]
	newSlice = append(newSlice, moves[0])
	return newSlice
}

func doRound23(elves map[common.Coordinate2]int, moves []dir23) (map[common.Coordinate2]int, []dir23, int) {
	proposedMove := make(map[common.Coordinate2]common.Coordinate2)
	proposedMoveCount := make(map[common.Coordinate2]int)
	newElves := make(map[common.Coordinate2]int)
elfLoop:
	for elf := range elves {
		neighbours := common.Coordinate2Neighbours9
		shouldMove := false
		for _, n := range neighbours {
			other := common.MoveBy2(elf, n)
			if _, ok := elves[other]; ok {
				shouldMove = true
				break
			}
		}
		if shouldMove {
			for _, move := range moves {
				if canMove(elves, elf, move) {
					newPos := move23(elf, move)
					proposedMove[elf] = newPos
					proposedMoveCount[newPos]++
					continue elfLoop
				}
			}
		}
		newElves[elf] = 1
	}
	movedElves := 0
	for source, dest := range proposedMove {
		if proposedMoveCount[dest] > 1 {
			newElves[source] = 1
		} else {
			newElves[dest] = 1
			movedElves++
		}
	}
	return newElves, rotateMoves(moves), movedElves
}

func getElves23(lines []string) map[common.Coordinate2]int {
	elves := make(map[common.Coordinate2]int)
	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			if char == "#" {
				elves[common.Coordinate2{Y: y, X: x}] = 1
			}
		}
	}
	return elves
}
