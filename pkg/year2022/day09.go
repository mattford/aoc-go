package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day09 struct{}

type ropeMove struct {
	dir       rune
	numerator int
}

func (p Day09) PartA(lines []string) any {
	moves := getMoves(lines)
	state := map[int]common.Coordinate2{
		0: {0, 0},
	}
	visited := make([]common.Coordinate2, 1)
	visited[0] = common.Coordinate2{}
	for _, move := range moves {
		state, visited = doMove(state, move, 2, visited)
	}
	return len(visited)
}

func (p Day09) PartB(lines []string) any {
	moves := getMoves(lines)
	state := map[int]common.Coordinate2{
		0: {0, 0},
	}
	visited := make([]common.Coordinate2, 1)
	visited[0] = common.Coordinate2{}
	for _, move := range moves {
		state, visited = doMove(state, move, 10, visited)
	}
	//printState(state)
	return len(visited)
}

func doMove(state map[int]common.Coordinate2, move ropeMove, knots int, visited []common.Coordinate2) (map[int]common.Coordinate2, []common.Coordinate2) {
	var vector common.Coordinate2
	switch move.dir {
	case 'L':
		vector = common.West
	case 'R':
		vector = common.East
	case 'U':
		vector = common.North
	case 'D':
		vector = common.South
	}
	for i := 1; i <= move.numerator; i++ {
		state[0] = common.MoveBy2(state[0], vector)
		for k := 1; k < knots; k++ {
			leader := state[k-1]
			location := state[k]
			if !common.IsAdjacent2(leader, location) && leader != location {
				vy := common.Bound(leader.Y-location.Y, -1, 1)
				vx := common.Bound(leader.X-location.X, -1, 1)
				state[k] = common.MoveBy2(state[k], common.Coordinate2{Y: vy, X: vx})
			}
			if k == knots-1 && !common.Contains(visited, state[k]) {
				visited = append(visited, state[k])
			}
		}
	}
	return state, visited
}

func getMoves(lines []string) []ropeMove {
	list := make([]ropeMove, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		n, _ := strconv.Atoi(parts[1])
		list[i] = ropeMove{rune(parts[0][0]), n}
	}
	return list
}

func printState(state map[int]common.Coordinate2) {
	var minY, maxY, minX, maxX int
	flipped := make(map[common.Coordinate2]int)
	for idx, coord := range state {
		minY = int(math.Min(float64(minY), float64(coord.Y)))
		maxY = int(math.Max(float64(maxY), float64(coord.Y)))
		minX = int(math.Min(float64(minX), float64(coord.X)))
		maxX = int(math.Max(float64(maxX), float64(coord.X)))
		flipped[coord] = idx
	}
	fmt.Println("State:")
	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			thing, ok := flipped[common.Coordinate2{Y: y, X: x}]
			if !ok {
				if y == 0 && x == 0 {
					fmt.Print("s")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(thing)
			}
		}
		fmt.Println()
	}
}
