package common

type Coordinate4 struct {
	Y int
	X int
	Z int
	W int
}

var Coordinate4Neighbours26 = []Coordinate4{
	{X: 1, Y: 0, Z: 0, W: 0},    // north
	{X: 0, Y: 1, Z: 0, W: 0},    // south
	{X: -1, Y: 0, Z: 0, W: 0},   // west
	{X: 0, Y: -1, Z: 0, W: 0},   // east
	{X: 1, Y: -1, Z: 0, W: 0},   // northeast
	{X: -1, Y: -1, Z: 0, W: 0},  // northwest
	{X: 1, Y: 1, Z: 0, W: 0},    // southeast
	{X: -1, Y: 1, Z: 0, W: 0},   // southwest
	{X: 0, Y: 0, Z: 1, W: 0},    // Forward
	{X: 0, Y: 0, Z: -1, W: 0},   // Back
	{X: 1, Y: 0, Z: 1, W: 0},    // north/forward
	{X: 0, Y: 1, Z: 1, W: 0},    // south/forward
	{X: -1, Y: 0, Z: 1, W: 0},   // west/forward
	{X: 0, Y: -1, Z: 1, W: 0},   // east/forward
	{X: 1, Y: -1, Z: 1, W: 0},   // northeast/forward
	{X: -1, Y: -1, Z: 1, W: 0},  // northwest/forward
	{X: 1, Y: 1, Z: 1, W: 0},    // southeast/forward
	{X: -1, Y: 1, Z: 1, W: 0},   // southwest/forward
	{X: 1, Y: 0, Z: -1, W: 0},   // north/back
	{X: 0, Y: 1, Z: -1, W: 0},   // south/back
	{X: -1, Y: 0, Z: -1, W: 0},  // west/back
	{X: 0, Y: -1, Z: -1, W: 0},  // east/back
	{X: 1, Y: -1, Z: -1, W: 0},  // northeast/back
	{X: -1, Y: -1, Z: -1, W: 0}, // northwest/back
	{X: 1, Y: 1, Z: -1, W: 0},   // southeast/back
	{X: -1, Y: 1, Z: -1, W: 0},  // southwest/back
}

var Coordinate4Neighbours80 = []Coordinate4{
	{X: 1, Y: 0, Z: 0, W: 0},    // north
	{X: 0, Y: 1, Z: 0, W: 0},    // south
	{X: -1, Y: 0, Z: 0, W: 0},   // west
	{X: 0, Y: -1, Z: 0, W: 0},   // east
	{X: 1, Y: -1, Z: 0, W: 0},   // northeast
	{X: -1, Y: -1, Z: 0, W: 0},  // northwest
	{X: 1, Y: 1, Z: 0, W: 0},    // southeast
	{X: -1, Y: 1, Z: 0, W: 0},   // southwest
	{X: 0, Y: 0, Z: 1, W: 0},    // Forward
	{X: 0, Y: 0, Z: -1, W: 0},   // Back
	{X: 1, Y: 0, Z: 1, W: 0},    // north/forward
	{X: 0, Y: 1, Z: 1, W: 0},    // south/forward
	{X: -1, Y: 0, Z: 1, W: 0},   // west/forward
	{X: 0, Y: -1, Z: 1, W: 0},   // east/forward
	{X: 1, Y: -1, Z: 1, W: 0},   // northeast/forward
	{X: -1, Y: -1, Z: 1, W: 0},  // northwest/forward
	{X: 1, Y: 1, Z: 1, W: 0},    // southeast/forward
	{X: -1, Y: 1, Z: 1, W: 0},   // southwest/forward
	{X: 1, Y: 0, Z: -1, W: 0},   // north/back
	{X: 0, Y: 1, Z: -1, W: 0},   // south/back
	{X: -1, Y: 0, Z: -1, W: 0},  // west/back
	{X: 0, Y: -1, Z: -1, W: 0},  // east/back
	{X: 1, Y: -1, Z: -1, W: 0},  // northeast/back
	{X: -1, Y: -1, Z: -1, W: 0}, // northwest/back
	{X: 1, Y: 1, Z: -1, W: 0},   // southeast/back
	{X: -1, Y: 1, Z: -1, W: 0},  // southwest/back

	{X: 1, Y: 0, Z: 0, W: 1},    // north
	{X: 0, Y: 1, Z: 0, W: 1},    // south
	{X: -1, Y: 0, Z: 0, W: 1},   // west
	{X: 0, Y: -1, Z: 0, W: 1},   // east
	{X: 1, Y: -1, Z: 0, W: 1},   // northeast
	{X: -1, Y: -1, Z: 0, W: 1},  // northwest
	{X: 1, Y: 1, Z: 0, W: 1},    // southeast
	{X: -1, Y: 1, Z: 0, W: 1},   // southwest
	{X: 0, Y: 0, Z: 1, W: 1},    // Forward
	{X: 0, Y: 0, Z: -1, W: 1},   // Back
	{X: 1, Y: 0, Z: 1, W: 1},    // north/forward
	{X: 0, Y: 1, Z: 1, W: 1},    // south/forward
	{X: -1, Y: 0, Z: 1, W: 1},   // west/forward
	{X: 0, Y: -1, Z: 1, W: 1},   // east/forward
	{X: 1, Y: -1, Z: 1, W: 1},   // northeast/forward
	{X: -1, Y: -1, Z: 1, W: 1},  // northwest/forward
	{X: 1, Y: 1, Z: 1, W: 1},    // southeast/forward
	{X: -1, Y: 1, Z: 1, W: 1},   // southwest/forward
	{X: 1, Y: 0, Z: -1, W: 1},   // north/back
	{X: 0, Y: 1, Z: -1, W: 1},   // south/back
	{X: -1, Y: 0, Z: -1, W: 1},  // west/back
	{X: 0, Y: -1, Z: -1, W: 1},  // east/back
	{X: 1, Y: -1, Z: -1, W: 1},  // northeast/back
	{X: -1, Y: -1, Z: -1, W: 1}, // northwest/back
	{X: 1, Y: 1, Z: -1, W: 1},   // southeast/back
	{X: -1, Y: 1, Z: -1, W: 1},  // southwest/back

	{X: 1, Y: 0, Z: 0, W: -1},    // north
	{X: 0, Y: 1, Z: 0, W: -1},    // south
	{X: -1, Y: 0, Z: 0, W: -1},   // west
	{X: 0, Y: -1, Z: 0, W: -1},   // east
	{X: 1, Y: -1, Z: 0, W: -1},   // northeast
	{X: -1, Y: -1, Z: 0, W: -1},  // northwest
	{X: 1, Y: 1, Z: 0, W: -1},    // southeast
	{X: -1, Y: 1, Z: 0, W: -1},   // southwest
	{X: 0, Y: 0, Z: 1, W: -1},    // Forward
	{X: 0, Y: 0, Z: -1, W: -1},   // Back
	{X: 1, Y: 0, Z: 1, W: -1},    // north/forward
	{X: 0, Y: 1, Z: 1, W: -1},    // south/forward
	{X: -1, Y: 0, Z: 1, W: -1},   // west/forward
	{X: 0, Y: -1, Z: 1, W: -1},   // east/forward
	{X: 1, Y: -1, Z: 1, W: -1},   // northeast/forward
	{X: -1, Y: -1, Z: 1, W: -1},  // northwest/forward
	{X: 1, Y: 1, Z: 1, W: -1},    // southeast/forward
	{X: -1, Y: 1, Z: 1, W: -1},   // southwest/forward
	{X: 1, Y: 0, Z: -1, W: -1},   // north/back
	{X: 0, Y: 1, Z: -1, W: -1},   // south/back
	{X: -1, Y: 0, Z: -1, W: -1},  // west/back
	{X: 0, Y: -1, Z: -1, W: -1},  // east/back
	{X: 1, Y: -1, Z: -1, W: -1},  // northeast/back
	{X: -1, Y: -1, Z: -1, W: -1}, // northwest/back
	{X: 1, Y: 1, Z: -1, W: -1},   // southeast/back
	{X: -1, Y: 1, Z: -1, W: -1},  // southwest/back

	{X: 0, Y: 0, Z: 0, W: 1},  // worward
	{X: 0, Y: 0, Z: 0, W: -1}, // wack
}

func IsAdjacent4(coord1 Coordinate4, coord2 Coordinate4) bool {
	for _, neighbour := range Coordinate4Neighbours80 {
		moved := MoveBy4(coord1, neighbour)
		if moved == coord2 {
			return true
		}
	}
	return false
}

func MoveBy4(coord Coordinate4, move Coordinate4) Coordinate4 {
	return Coordinate4{
		coord.Y + move.Y,
		coord.X + move.X,
		coord.Z + move.Z,
		coord.W + move.W,
	}
}
