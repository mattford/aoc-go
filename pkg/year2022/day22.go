package year2022

import (
	"math"
	"strconv"
	"strings"
)

/**
This solution is large "borrowed" from https://www.reddit.com/r/adventofcode/comments/zsct8w/comment/j1bzuzm/
but has been adapted to work for both parts of the puzzle
*/

type Day22 struct{}

func (p Day22) PartA(lines []string) any {
	grid := parseInput(lines, false)
	grid.run()
	return grid.p.password()
}

func (p Day22) PartB(lines []string) any {
	cube := parseInput(lines, true)
	cube.run()
	return cube.p.password()
}

type dir int

const right dir = 0
const down dir = 1
const left dir = 2
const up dir = 3
const invalidDir dir = -1

func (d dir) String() string {
	switch d {
	case 0:
		return "right"
	case 1:
		return "down"
	case 2:
		return "left"
	case 3:
		return "up"
	default:
		return "invalid"
	}
}

func (d dir) next() dir    { return (d + 1) % 4 }
func (d dir) prev() dir    { return (d + 3) % 4 }
func (d dir) flipped() dir { return (d + 2) % 4 }
func (d dir) turn(turnDir byte) dir {
	switch turnDir {
	case 'R':
		return d.next()
	case 'L':
		return d.prev()
	default:
		panic("Invalid turn")
	}
}

type player struct {
	cell      *cell
	direction dir
}

func (p *player) move(cube bool) bool {
	next := p.cell.neighbors[p.direction]
	if next.val == '#' { // wall
		return false
	}
	if cube && p.cell.face.id != next.face.id {
		p.direction = (next.face.entrySide(p.cell.face) + 2) % 4
	}
	p.cell = next
	return true
}

func (p *player) password() int {
	return 1000*(p.cell.row+1) + 4*(p.cell.col+1) + int(p.direction)
}

type face struct {
	id, originRow, originCol int
	cells                    [][]*cell
	neighbors                [4]*face
}

func NewFace(id, originRow, originCol, sideLength int) *face {
	f := face{
		id:        id,
		originRow: originRow,
		originCol: originCol,
		cells:     make([][]*cell, sideLength),
	}
	for i := range f.cells {
		f.cells[i] = make([]*cell, sideLength)
	}
	return &f
}

func (f face) sideClockwise(side dir) []*cell {
	cells := make([]*cell, len(f.cells))
	n := len(cells) - 1
	var row, col, rowInc, colInc int
	switch side {
	case right:
		row, col, rowInc, colInc = 0, n, 1, 0
	case down:
		row, col, rowInc, colInc = n, n, 0, -1
	case left:
		row, col, rowInc, colInc = n, 0, -1, 0
	case up:
		row, col, rowInc, colInc = 0, 0, 0, 1
	}
	for i := 0; i <= n; i++ {
		cells[i] = f.cells[row][col]
		row += rowInc
		col += colInc
	}

	return cells
}

func (f face) entrySide(source *face) dir {
	for entry := right; entry <= up; entry++ {
		if f.neighbors[entry] == source {
			return entry
		}
	}
	return invalidDir
}

type cell struct {
	face      *face
	row, col  int
	val       byte
	neighbors [4]*cell
}

type board struct {
	faces  [6]*face
	p      *player
	moves  []int
	turns  []byte
	isCube bool
}

func (b *board) run() {
	for i, m := range b.moves {
		for j := 0; j < m; j++ {
			if !b.p.move(b.isCube) {
				break
			}
		}
		if i < len(b.turns) {
			b.p.direction = b.p.direction.turn(b.turns[i])
		} else {
			break
		}
	}
}

func connectSides(f1, f2 *face, s1, s2 dir) {
	side1 := f1.sideClockwise(s1)
	side2 := f2.sideClockwise(s2)
	n := len(side1) - 1
	for i := range side1 {
		cell1, cell2 := side1[i], side2[n-i]
		cell1.neighbors[s1], cell2.neighbors[s2] = cell2, cell1
	}
	f1.neighbors[s1], f2.neighbors[s2] = f2, f1
}

func parseCube(input []string, cube bool) [6]*face {
	// Count total number of cells to derive cube face size and board width
	cellCount := 0
	boardWidth := 0
	for _, line := range input {
		if len(line) > boardWidth {
			boardWidth = len(line)
		}
		for _, r := range line {
			if r != ' ' {
				cellCount++
			}
		}
	}
	sideLength := int(math.Sqrt(float64(cellCount / 6)))

	// Create faces and parse cells
	var faces [6]*face
	faceId := 0
	faceGrid := make([][]*face, len(input)/sideLength)
	for faceRowNum := range faceGrid {
		faceRow := make([]*face, boardWidth/sideLength)
		faceGrid[faceRowNum] = faceRow
		boardRowNum := faceRowNum * sideLength

		for faceColNum := range faceGrid[faceRowNum] {
			boardColNum := faceColNum * sideLength
			if boardColNum >= len(input[boardRowNum]) || input[boardRowNum][boardColNum] == ' ' {
				continue
			}

			newFace := NewFace(faceId, boardRowNum, boardColNum, sideLength)
			for row := 0; row < sideLength; row++ {
				for col := 0; col < sideLength; col++ {
					cellRow, cellCol := newFace.originRow+row, newFace.originCol+col
					newCell := cell{
						face: newFace,
						row:  cellRow,
						col:  cellCol,
						val:  input[cellRow][cellCol],
					}
					newFace.cells[row][col] = &newCell
				}
			}

			faces[faceId] = newFace
			faceId++
			faceGrid[faceRowNum][faceColNum] = newFace
		}
	}

	// Set cell neighbors within faces
	for _, f := range faces {
		for _, row := range f.cells {
			for colNum, rightCell := range row[1:] {
				c := row[colNum]
				if c.neighbors[right] != nil || rightCell.neighbors[left] != nil {
					panic("Overwriting neighbor config")
				}
				c.neighbors[right], rightCell.neighbors[left] = rightCell, c
			}
		}
		for rowNum, downRow := range f.cells[1:] {
			for colNum, downCell := range downRow {
				c := f.cells[rowNum][colNum]
				if c.neighbors[down] != nil || downCell.neighbors[up] != nil {
					panic("Overwriting neighbor config")
				}
				c.neighbors[down], downCell.neighbors[up] = downCell, c
			}
		}
	}

	// Set face neighbors based on board before folding
	for _, fRow := range faceGrid {
		for colNum := range fRow[1:] {
			if fRow[colNum] != nil && fRow[colNum+1] != nil {
				connectSides(fRow[colNum], fRow[colNum+1], right, left)
			}
		}
	}
	for rowNum, fRow := range faceGrid[1:] {
		for colNum, f := range fRow {
			if faceGrid[rowNum][colNum] != nil && f != nil {
				connectSides(faceGrid[rowNum][colNum], f, down, up)
			}
		}
	}

	if !cube {
		// Connect the outside edges to their opposite edge
		for _, fRow := range faceGrid {
			firstIdx := firstIdx(fRow)
			lastIdx := lastIdx(fRow)
			connectSides(fRow[firstIdx], fRow[lastIdx], left, right)
		}
		for i := 0; i < len(faceGrid[0]); i++ {
			col := make([]*face, 0)
			for _, fRow := range faceGrid {
				col = append(col, fRow[i])
			}
			firstIdx := firstIdx(col)
			lastIdx := lastIdx(col)
			connectSides(faceGrid[firstIdx][i], faceGrid[lastIdx][i], up, down)
		}
		return faces
	}
	// Finally, fold the cube and set new neighbors
	// (6 faces)*(4 edges per face) = 24 edges forms 12 pairs
	// 5 pairs are handled in the last section, 7 remain
	// Strategy: Repeatedly check for L-shapes and fold them
	for disconnectedPairs := 7; disconnectedPairs > 0; {
		for _, f := range faces {
			for side := right; side <= up; side++ {
				if f.neighbors[side] != nil {
					continue
				}

				next, prev := side.next(), side.prev()
				if mid := f.neighbors[next]; mid != nil { // clockwise search
					next = mid.entrySide(f).next()
					if fold := mid.neighbors[next]; fold != nil {
						next = fold.entrySide(mid).next()
						if fold.neighbors[next] == nil {
							connectSides(f, fold, side, next)
							disconnectedPairs--
						}
					}
				} else if mid := f.neighbors[prev]; mid != nil { // anticlockwise search
					prev = mid.entrySide(f).prev()
					if fold := mid.neighbors[prev]; fold != nil {
						prev = fold.entrySide(mid).prev()
						if fold.neighbors[prev] == nil {
							connectSides(f, fold, side, prev)
							disconnectedPairs--
						}
					}
				}
			}
		}
	}

	return faces
}

func firstIdx(faces []*face) int {
	for i, f := range faces {
		if f != nil {
			return i
		}
	}
	return -1
}

func lastIdx(faces []*face) int {
	for i := len(faces) - 1; i >= 0; i-- {
		if faces[i] != nil {
			return i
		}
	}
	return -1
}

func parseInstruction(instruction string) ([]int, []byte) {
	moves, turns := make([]int, 0), make([]byte, 0)
	for len(instruction) > 0 {
		nextTurn := strings.IndexAny(instruction, "LR")
		if nextTurn == -1 {
			if tiles, err := strconv.Atoi(instruction); err == nil {
				moves = append(moves, tiles)
				break
			} else {
				panic(err)
			}
		} else {
			if tiles, err := strconv.Atoi(instruction[:nextTurn]); err == nil {
				moves = append(moves, tiles)
			} else {
				panic(err)
			}
			turns = append(turns, instruction[nextTurn])
			instruction = instruction[nextTurn+1:]
		}
	}
	return moves, turns
}

func parseInput(input []string, cube bool) *board {
	var emptyLine int
	for lineNum, line := range input {
		if line == "" {
			emptyLine = lineNum
			break
		}
	}

	b := new(board)
	b.faces = parseCube(input[:emptyLine], cube)
	b.moves, b.turns = parseInstruction(input[emptyLine+1])
	b.isCube = cube

	for colNum := 0; colNum < len(b.faces[0].cells); colNum++ {
		if startingCell := b.faces[0].cells[0][colNum]; startingCell.val == '.' {
			b.p = &player{cell: startingCell, direction: right}
			break
		}
	}

	return b
}
