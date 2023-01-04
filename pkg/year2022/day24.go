package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"sort"
	"strings"
)

type Day24 struct{}

type state24 struct {
	pos  common.Coordinate2
	cost int
}

func (s state24) hash(length int) string {
	return fmt.Sprintf("%v,%v,%v,%v", s.pos.Y, s.pos.X, s.cost, s.cost%length)
}

func (p Day24) PartA(lines []string) any {
	s, target, grid, blizzards := parseInput24(lines)
	return findShortest(s, target, grid, blizzards)
}

func (p Day24) PartB(lines []string) any {
	s, target, grid, blizzards := parseInput24(lines)
	initialPos := s.pos
	there := findShortest(s, target, grid, blizzards)
	back := findShortest(state24{pos: target, cost: there}, initialPos, grid, blizzards)
	return findShortest(state24{pos: initialPos, cost: back}, target, grid, blizzards)
}

func findShortest(initial state24, target common.Coordinate2, grid map[common.Coordinate2]int, blizzards [][]common.Coordinate2) int {
	queue := []state24{
		initial,
	}
	blizLength := len(blizzards)
	v := make([]string, 0)
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if common.Contains(v, s.hash(blizLength)) {
			continue
		}
		v = append(v, s.hash(blizLength))
		bliz := blizzards[(s.cost+1)%len(blizzards)]
		for _, neighbour := range common.Coordinate2Neighbours4 {
			other := common.MoveBy2(s.pos, neighbour)
			if _, ok := grid[other]; ok && !common.Contains(bliz, other) {
				nextState := state24{
					cost: s.cost + 1,
					pos:  other,
				}
				if other == target {
					return s.cost + 1
				}
				queue = append(queue, nextState)
			}
		}
		if !common.Contains(bliz, s.pos) {
			queue = append(queue, state24{
				cost: s.cost + 1,
				pos:  s.pos,
			})
		}

		sort.Slice(queue, func(i, j int) bool {
			a := queue[i]
			b := queue[j]
			return a.cost < b.cost
		})
	}
	return 0
}

func progressBlizzards(grid map[common.Coordinate2]int, blizzards map[common.Coordinate2][]rune) map[common.Coordinate2][]rune {
	minX, maxX, minY, maxY := common.Bounds(common.Keys(grid))
	minY++
	maxY--
	newBlizzards := make(map[common.Coordinate2][]rune)
	for pos, bliz := range blizzards {
		for _, d := range bliz {
			var nextPos common.Coordinate2
			switch d {
			case '>':
				nextPos = common.MoveBy2(pos, common.East)
			case '<':
				nextPos = common.MoveBy2(pos, common.West)
			case 'v':
				nextPos = common.MoveBy2(pos, common.South)
			case '^':
				nextPos = common.MoveBy2(pos, common.North)
			}
			if nextPos.Y < minY {
				nextPos.Y = maxY
			} else if nextPos.Y > maxY {
				nextPos.Y = minY
			}
			if nextPos.X < minX {
				nextPos.X = maxX
			} else if nextPos.X > maxX {
				nextPos.X = minX
			}
			newBlizzards[nextPos] = append(newBlizzards[nextPos], d)
		}
	}
	return newBlizzards
}

func parseInput24(lines []string) (state24, common.Coordinate2, map[common.Coordinate2]int, [][]common.Coordinate2) {
	var pos, target common.Coordinate2
	grid := make(map[common.Coordinate2]int)
	initialBlizzards := make(map[common.Coordinate2][]rune)
	height := len(lines) - 2
	width := len(lines[1]) - 2
	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			coord := common.Coordinate2{Y: y, X: x}
			if char != "#" {
				grid[coord] = 1
				if char != "." {
					initialBlizzards[coord] = append(initialBlizzards[coord], rune(char[0]))
				}
			}
			if y == 0 && char == "." {
				pos = coord
			} else if y == len(lines)-1 && char == "." {
				target = coord
			}
		}
	}
	lcm := common.LCM(height, width)
	blizzards := make([][]common.Coordinate2, lcm)
	for i := 0; i < lcm; i++ {
		blizzards[i] = common.Keys(initialBlizzards)
		initialBlizzards = progressBlizzards(grid, initialBlizzards)
	}

	return state24{pos: pos}, target, grid, blizzards
}
