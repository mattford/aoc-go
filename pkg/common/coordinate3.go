package common

type Coordinate3 struct {
	Y int
	X int
	Z int
}

var Coordinate3Neighbours6 = []Coordinate3{
	{X: 1, Y: 0, Z: 0},  // north
	{X: 0, Y: 1, Z: 0},  // south
	{X: -1, Y: 0, Z: 0}, // west
	{X: 0, Y: -1, Z: 0}, // east
	{X: 0, Y: 0, Z: 1},  // Forward
	{X: 0, Y: 0, Z: -1}, // Back
}

var Coordinate3Neighbours26 = []Coordinate3{
	{X: 1, Y: 0, Z: 0},    // north
	{X: 0, Y: 1, Z: 0},    // south
	{X: -1, Y: 0, Z: 0},   // west
	{X: 0, Y: -1, Z: 0},   // east
	{X: 1, Y: -1, Z: 0},   // northeast
	{X: -1, Y: -1, Z: 0},  // northwest
	{X: 1, Y: 1, Z: 0},    // southeast
	{X: -1, Y: 1, Z: 0},   // southwest
	{X: 0, Y: 0, Z: 1},    // Forward
	{X: 0, Y: 0, Z: -1},   // Back
	{X: 1, Y: 0, Z: 1},    // north/forward
	{X: 0, Y: 1, Z: 1},    // south/forward
	{X: -1, Y: 0, Z: 1},   // west/forward
	{X: 0, Y: -1, Z: 1},   // east/forward
	{X: 1, Y: -1, Z: 1},   // northeast/forward
	{X: -1, Y: -1, Z: 1},  // northwest/forward
	{X: 1, Y: 1, Z: 1},    // southeast/forward
	{X: -1, Y: 1, Z: 1},   // southwest/forward
	{X: 1, Y: 0, Z: -1},   // north/back
	{X: 0, Y: 1, Z: -1},   // south/back
	{X: -1, Y: 0, Z: -1},  // west/back
	{X: 0, Y: -1, Z: -1},  // east/back
	{X: 1, Y: -1, Z: -1},  // northeast/back
	{X: -1, Y: -1, Z: -1}, // northwest/back
	{X: 1, Y: 1, Z: -1},   // southeast/back
	{X: -1, Y: 1, Z: -1},  // southwest/back
}

func IsAdjacent3(coord1 Coordinate3, coord2 Coordinate3) bool {
	for _, neighbour := range Coordinate3Neighbours26 {
		moved := MoveBy3(coord1, neighbour)
		if moved == coord2 {
			return true
		}
	}
	return false
}

func MoveBy3(coord Coordinate3, move Coordinate3) Coordinate3 {
	return Coordinate3{
		coord.Y + move.Y,
		coord.X + move.X,
		coord.Z + move.Z,
	}
}
