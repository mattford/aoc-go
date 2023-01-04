package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"sort"
	"strings"
)

type Day24 struct{}

type state24 struct {
	grid      map[common.Coordinate2]int
	blizzards map[common.Coordinate2][]rune
	pos       common.Coordinate2
	cost      int
}

func (s state24) hash() string {
	blizMap := ""
	for c, b := range s.blizzards {
		blizMap += fmt.Sprintf("%v,%v:%v/", c.Y, c.X, string(b))
	}
	return fmt.Sprintf("%v,%v,%v", s.pos.Y, s.pos.X, blizMap)
}

func (p Day24) PartA(lines []string) any {
	s, target := parseInput24(lines)
	return findShortest(s, target)
}

func (p Day24) PartB(lines []string) any {
	return "implement_me"
}

func findShortest(initial state24, target common.Coordinate2) int {
	queue := []state24{
		initial,
	}
	v := make([]string, 0)
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if common.Contains(v, s.hash()) {
			continue
		}
		s.progressBlizzards()
		for _, neighbour := range common.Coordinate2Neighbours4 {
			other := common.MoveBy2(s.pos, neighbour)
			if _, ok := s.grid[other]; ok {
				if _, ok2 := s.blizzards[other]; !ok2 {
					nextState := state24{
						blizzards: s.blizzards,
						cost:      s.cost + 1,
						grid:      s.grid,
						pos:       other,
					}
					if other == target {
						return s.cost
					}
					if !common.Contains(v, nextState.hash()) {
						queue = append(queue, nextState)
					}
				}
			}
		}
		queue = append(queue, state24{
			blizzards: s.blizzards,
			cost:      s.cost + 1,
			grid:      s.grid,
			pos:       s.pos,
		})

		sort.Slice(queue, func(i, j int) bool {
			a := queue[i]
			b := queue[j]
			costDiff := a.cost - b.cost
			if costDiff == 0 {
				return common.Manhattan(a.pos, target) < common.Manhattan(b.pos, target)
			}
			return costDiff > 0
		})
	}
	return 0
}

func (s state24) printBliz() {
	for c, b := range s.blizzards {
		fmt.Println(c.Y, c.X, string(b))
	}
}

func (s *state24) progressBlizzards() {
	minX, maxX, minY, maxY := common.Bounds(common.Keys(s.grid))
	minX++
	maxX--
	minY++
	maxY--
	newBlizzards := make(map[common.Coordinate2][]rune)
	for pos, blizzards := range s.blizzards {
		for _, d := range blizzards {
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
	s.blizzards = newBlizzards
}

func parseInput24(lines []string) (state24, common.Coordinate2) {
	var pos, target common.Coordinate2
	grid := make(map[common.Coordinate2]int)
	blizzards := make(map[common.Coordinate2][]rune)
	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			coord := common.Coordinate2{Y: y, X: x}
			if char != "#" {
				grid[coord] = 1
				if char != "." {
					blizzards[coord] = append(blizzards[coord], rune(char[0]))
				}
			}
			if y == 0 && char == "." {
				pos = coord
			} else if y == len(lines)-1 && char == "." {
				target = coord
			}
		}
	}
	return state24{pos: pos, grid: grid, blizzards: blizzards}, target
}
