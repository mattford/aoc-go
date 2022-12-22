package common

import "math"

type Coordinate2 struct {
	Y int
	X int
}

var (
	East      = Coordinate2{X: 1, Y: 0}
	South     = Coordinate2{X: 0, Y: 1}
	West      = Coordinate2{X: -1, Y: 0}
	North     = Coordinate2{X: 0, Y: -1}
	NorthEast = Coordinate2{X: 1, Y: -1}
	NorthWest = Coordinate2{X: -1, Y: -1}
	SouthEast = Coordinate2{X: 1, Y: 1}
	SouthWest = Coordinate2{X: -1, Y: 1}
)

var Coordinate2Neighbours9 = []Coordinate2{
	North,
	East,
	South,
	West,
	NorthEast,
	NorthWest,
	SouthEast,
	SouthWest,
}

var Coordinate2Neighbours4 = []Coordinate2{
	North,
	East,
	South,
	West,
}

func IsAdjacent2(coord1 Coordinate2, coord2 Coordinate2) bool {
	for _, neighbour := range Coordinate2Neighbours9 {
		moved := MoveBy2(coord1, neighbour)
		if moved == coord2 {
			return true
		}
	}
	return false
}

func MoveBy2(coord Coordinate2, move Coordinate2) Coordinate2 {
	return Coordinate2{
		coord.Y + move.Y,
		coord.X + move.X,
	}
}

func Manhattan(coord Coordinate2, coord2 Coordinate2) int {
	return int(math.Abs(float64(coord.Y-coord2.Y)) + math.Abs(float64(coord.X-coord2.X)))
}

func Bounds(rock []Coordinate2) (int, int, int, int) {
	var minX, minY, maxX, maxY int
	minX = math.MaxInt
	minY = math.MaxInt
	maxY = math.MinInt
	maxX = math.MinInt
	for _, point := range rock {
		minX = int(math.Min(float64(point.X), float64(minX)))
		minY = int(math.Min(float64(point.Y), float64(minY)))
		maxX = int(math.Max(float64(point.X), float64(maxX)))
		maxY = int(math.Max(float64(point.Y), float64(maxY)))
	}
	return minX, maxX, minY, maxY
}
