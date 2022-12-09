package common

type Coordinate2 struct {
	Y int
	X int
}

var Coordinate2Neighbours9 = []Coordinate2{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
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
