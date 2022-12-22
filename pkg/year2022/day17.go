package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"strings"
)

type Day17 struct{}

func (p Day17) PartA(lines []string) any {
	moves := parseInput17(lines)
	grid := make(map[common.Coordinate2]int)
	var floor, iter int
	for i := 0; i < 2022; i++ {
		iter, grid = simulateRock(moves, iter, i, floor, grid)
		_, _, minY, _ := common.Bounds(common.Keys(grid))
		floor = minY
	}
	//printGrid17(grid)
	return math.Abs(float64(floor))
}

type previousHeight struct {
	floorLevel int
	rocks      int
}

func (p Day17) PartB(lines []string) any {
	moves := parseInput17(lines)
	grid := make(map[common.Coordinate2]int)
	var floor, iter int
	seenCombos := make(map[string]previousHeight)
	target := 1000000000000
	for i := 0; i < target; i++ {
		rIdx := i % 5
		key := hashState(rIdx, iter)
		if prevHeight, ok := seenCombos[key]; ok {
			if n, d := target-i, i-prevHeight.rocks; n%d == 0 {
				diff := floor - prevHeight.floorLevel
				diffRocks := i - prevHeight.rocks
				absDiff := int(math.Abs(float64(diff)))
				return int(math.Abs(float64(floor - (target-i)/diffRocks*(absDiff))))
			}
		}
		seenCombos[key] = previousHeight{
			rocks: i, floorLevel: floor,
		}
		iter, grid = simulateRock(moves, iter, i, floor, grid)
		_, _, minY, _ := common.Bounds(common.Keys(grid))
		floor = minY
	}
	return int(math.Abs(float64(floor)))
}

func hashState(rIdx, iter int) string {
	return fmt.Sprintf("%v/%v", rIdx, iter)
}

func printGrid17(grid map[common.Coordinate2]int) {
	coords := common.Keys(grid)
	_, _, minY, _ := common.Bounds(coords)
	for y := minY - 1; y < 0; y++ {
		fmt.Print("|")
		for x := 0; x <= 6; x++ {
			if _, ok := grid[common.Coordinate2{Y: y, X: x}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|")
		fmt.Println()
	}
	fmt.Print("+")
	for x := 0; x <= 6; x++ {
		fmt.Print("-")
	}
	fmt.Print("+")
	fmt.Println()
}

func simulateRock(moves []common.Coordinate2, iter, iteration, floor int, grid map[common.Coordinate2]int) (int, map[common.Coordinate2]int) {
	rock := getRock(iteration, floor)
	for {
		move := moves[iter]
		iter = (iter + 1) % len(moves)
		newRock := moveRock(rock, move)
		minX, maxX, _, _ := common.Bounds(newRock)
		if minX >= 0 && maxX <= 6 && !rockIntersects(newRock, grid) {
			rock = newRock
		}
		newRock = moveRock(rock, common.South)
		_, _, _, maxY := common.Bounds(newRock)
		if maxY >= 0 || rockIntersects(newRock, grid) {
			for _, point := range rock {
				grid[point] = 1
			}
			break
		}
		rock = newRock
	}
	return iter, grid
}

func rockIntersects(rock []common.Coordinate2, grid map[common.Coordinate2]int) bool {
	for _, point := range rock {
		if _, ok := grid[point]; ok {
			return true
		}
	}
	return false
}

func moveRock(rock []common.Coordinate2, move common.Coordinate2) []common.Coordinate2 {
	newRock := make([]common.Coordinate2, 0, len(rock))
	for _, point := range rock {
		newRock = append(newRock, common.MoveBy2(point, move))
	}
	return newRock
}

func getRock(i, floor int) []common.Coordinate2 {
	i = i % 5
	y := floor - 4
	switch i {
	case 0:
		return []common.Coordinate2{
			{Y: y, X: 2},
			{Y: y, X: 3},
			{Y: y, X: 4},
			{Y: y, X: 5},
		}
	case 1:
		return []common.Coordinate2{
			{Y: y - 2, X: 3},
			{Y: y - 1, X: 2},
			{Y: y - 1, X: 3},
			{Y: y - 1, X: 4},
			{Y: y, X: 3},
		}
	case 2:
		return []common.Coordinate2{
			{Y: y - 2, X: 4},
			{Y: y - 1, X: 4},
			{Y: y, X: 2},
			{Y: y, X: 3},
			{Y: y, X: 4},
		}
	case 3:
		return []common.Coordinate2{
			{Y: y - 3, X: 2},
			{Y: y - 2, X: 2},
			{Y: y - 1, X: 2},
			{Y: y, X: 2},
		}
	case 4:
		return []common.Coordinate2{
			{Y: y - 1, X: 2},
			{Y: y - 1, X: 3},
			{Y: y, X: 2},
			{Y: y, X: 3},
		}
	}
	return []common.Coordinate2{}
}

func parseInput17(lines []string) []common.Coordinate2 {
	moves := make([]common.Coordinate2, 0, len(lines[0]))
	for _, x := range strings.Split(lines[0], "") {
		switch x {
		case "<":
			moves = append(moves, common.West)
		case ">":
			moves = append(moves, common.East)
		}
	}
	return moves
}
