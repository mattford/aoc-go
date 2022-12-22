package year2022

import (
	"aocgen/pkg/common"
	"math"
	"strings"
)

type Day18 struct{}

func (p Day18) PartA(lines []string) any {
	exposed := 0
	points := parseInput18(lines)
	for _, point := range points {
		for _, n := range common.Coordinate3Neighbours6 {
			other := common.MoveBy3(point, n)
			if !common.Contains(points, other) {
				exposed++
			}
		}
	}
	return exposed
}

func (p Day18) PartB(lines []string) any {
	exposed := 0
	points := parseInput18(lines)
	min := math.MaxInt
	max := math.MinInt
	for _, point := range points {
		min = common.MinInt([]int{min, point.X, point.Y, point.Z})
		max = common.MaxInt([]int{max, point.X, point.Y, point.Z})
	}
	visited := make([]common.Coordinate3, 0)
	queue := []common.Coordinate3{
		{0, 0, 0},
	}
	for len(queue) > 0 {
		coord := queue[0]
		queue = queue[1:]
		if common.Contains(visited, coord) || common.Contains(points, coord) {
			continue
		}
		if coord.Z < min-1 || coord.Y < min-1 || coord.X < min-1 || coord.Z > max+1 || coord.Y > max+1 || coord.X > max+1 {
			continue
		}
		visited = append(visited, coord)

		exposed += countAffectedCubes(points, coord)

		for _, n := range common.Coordinate3Neighbours6 {
			queue = append(queue, common.MoveBy3(coord, n))
		}
	}
	return exposed
}

func countAffectedCubes(points []common.Coordinate3, coord common.Coordinate3) (exposed int) {
	for _, n := range common.Coordinate3Neighbours6 {
		other := common.MoveBy3(coord, n)
		if common.Contains(points, other) {
			exposed++
		}
	}
	return exposed
}

func parseInput18(lines []string) []common.Coordinate3 {
	list := make([]common.Coordinate3, 0, len(lines))
	for _, line := range lines {
		ints := common.GetInts(strings.Split(line, ","))
		list = append(list, common.Coordinate3{X: ints[0], Y: ints[1], Z: ints[2]})
	}
	return list
}
