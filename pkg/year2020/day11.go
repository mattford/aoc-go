package year2020

const EMPTY = 'L'
const FLOOR = '.'
const TAKEN = '#'

type Day11 struct{}

func (p Day11) PartA(lines []string) any {
	state := getGrid(lines)
	for {
		nextState := progressState(state, false)
		if same(state, nextState) {
			break
		}
		state = nextState
	}
	return countType(state, TAKEN)
}

func (p Day11) PartB(lines []string) any {
	state := getGrid(lines)
	for {
		nextState := progressState(state, true)
		if same(state, nextState) {
			break
		}
		state = nextState
	}
	return countType(state, TAKEN)
}

func countType(state [][]rune, target rune) int {
	count := 0
	for _, row := range state {
		for _, col := range row {
			if col == target {
				count++
			}
		}
	}
	return count
}

func progressState(state [][]rune, enhancedVision bool) [][]rune {
	nextState := make([][]rune, len(state))
	for i, row := range state {
		nextState[i] = make([]rune, len(row))
		copy(nextState[i], row)
	}
	requiredToBecomeEmpty := 4
	if enhancedVision {
		requiredToBecomeEmpty = 5
	}
	for y := 0; y < len(state); y++ {
		for x := 0; x < len(state[y]); x++ {
			thisSeat := state[y][x]
			adjacents := getAdjacentOccupied(y, x, state, enhancedVision)
			if thisSeat == EMPTY && adjacents == 0 {
				nextState[y][x] = TAKEN
			} else if thisSeat == TAKEN && adjacents >= requiredToBecomeEmpty {
				nextState[y][x] = EMPTY
			}
		}
	}
	return nextState
}

func getAdjacentOccupied(y int, x int, grid [][]rune, enhancedVision bool) int {
	if enhancedVision {
		return getOccupiedWithEnhancedVision(y, x, grid)
	}
	count := 0
	coords := [][]int{
		{y - 1, x},
		{y + 1, x},
		{y, x - 1},
		{y, x + 1},
		{y - 1, x - 1},
		{y - 1, x + 1},
		{y + 1, x - 1},
		{y + 1, x + 1},
	}
	for _, coord := range coords {
		y2 := coord[0]
		x2 := coord[1]
		if y2 < 0 || y2 >= len(grid) || x2 < 0 || x2 >= len(grid[y2]) {
			continue
		} else if grid[y2][x2] == TAKEN {
			count++
		}
	}
	return count
}

func getOccupiedWithEnhancedVision(y int, x int, grid [][]rune) (count int) {
	slopes := [][]int{
		{0, 1},
		{-1, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
	}
	for _, slope := range slopes {
		firstAtSlope := getFirstAtSlope(y, x, grid, slope)
		if firstAtSlope == TAKEN {
			count++
		}
	}
	return
}

func getFirstAtSlope(y int, x int, grid [][]rune, slope []int) rune {
	yDiff := slope[0]
	xDiff := slope[1]
	for {
		y += yDiff
		x += xDiff
		if y < 0 || x < 0 || y >= len(grid) || x >= len(grid[y]) {
			break
		}
		nextTile := grid[y][x]
		if nextTile == TAKEN || nextTile == EMPTY {
			return nextTile
		}
	}
	return FLOOR
}

func same(stateA [][]rune, stateB [][]rune) bool {
	for y := 0; y < len(stateA); y++ {
		for x := 0; x < len(stateA[y]); x++ {
			if stateA[y][x] != stateB[y][x] {
				return false
			}
		}
	}
	return true
}

func getGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}
