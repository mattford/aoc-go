package year2020

import (
	"math"
	"regexp"
	"strconv"
)

type Day12 struct{}

type coords struct {
	y int
	x int
}

type state struct {
	direction rune
	y         int
	x         int
	waypoint  coords
}

type instruction struct {
	operation rune
	numerator int
}

func (p Day12) PartA(lines []string) any {
	instructions := parseInput(lines)
	currentState := state{
		direction: 'E',
		y:         0,
		x:         0,
	}
	for _, instr := range instructions {
		currentState = doInstruction(currentState, instr)
	}

	return int(math.Abs(float64(currentState.x)) + math.Abs(float64(currentState.y)))
}

func (p Day12) PartB(lines []string) any {
	instructions := parseInput(lines)
	currentState := state{
		direction: 'E',
		y:         0,
		x:         0,
		waypoint:  coords{-1, 10},
	}
	for _, instr := range instructions {
		currentState = doInstructionWaypoint(currentState, instr)
	}

	return int(math.Abs(float64(currentState.x)) + math.Abs(float64(currentState.y)))
}

func doInstructionWaypoint(currentState state, instr instruction) state {
	switch instr.operation {
	case 'N':
		currentState.waypoint.y -= instr.numerator
	case 'S':
		currentState.waypoint.y += instr.numerator
	case 'E':
		currentState.waypoint.x += instr.numerator
	case 'W':
		currentState.waypoint.x -= instr.numerator
	case 'L':
		return rotate(currentState, 'L', instr.numerator)
	case 'R':
		return rotate(currentState, 'R', instr.numerator)
	case 'F':
		for i := 0; i < instr.numerator; i++ {
			currentState.y += currentState.waypoint.y
			currentState.x += currentState.waypoint.x
		}
	}
	return currentState
}

func doInstruction(currentState state, instr instruction) state {
	switch instr.operation {
	case 'N':
		currentState.y -= instr.numerator
	case 'S':
		currentState.y += instr.numerator
	case 'E':
		currentState.x += instr.numerator
	case 'W':
		currentState.x -= instr.numerator
	case 'L':
		return turn(currentState, 'L', instr.numerator)
	case 'R':
		return turn(currentState, 'R', instr.numerator)
	case 'F':
		return doInstruction(currentState, instruction{operation: currentState.direction, numerator: instr.numerator})
	}
	return currentState
}

func rotate(currentState state, direction rune, amount int) state {
	times := amount / 90
	for i := 0; i < times; i++ {

		switch direction {
		case 'L':
			// 1, -1 => 1, 1 => -1, 1 => -1, -1
			// 0, 1 => -1, 0 => 0, -1 => 1, 0
			// 5, 1 => -1, 5 => -5, -1 => 1, -5
			// 5, -1 => 1, 5 => -5, 1 => -1, -5
			newY := 0 - currentState.waypoint.x
			newX := currentState.waypoint.y
			currentState.waypoint = coords{newY, newX}
		case 'R':
			// 1, 1 => 1, -1 => -1, -1 => -1, 1
			// 5, 1 => 1, -5 => -5, -1 => -1, 5
			newY := currentState.waypoint.x
			newX := 0 - currentState.waypoint.y
			currentState.waypoint = coords{newY, newX}
		}
	}
	return currentState
}

func turn(currentState state, direction rune, amount int) state {
	rightTurns := map[rune]rune{
		'N': 'E',
		'E': 'S',
		'S': 'W',
		'W': 'N',
	}
	leftTurns := make(map[rune]rune, 4)
	for key, val := range rightTurns {
		leftTurns[val] = key
	}
	if amount == 360 {
		return currentState
	}
	times := amount / 90
	for i := 0; i < times; i++ {
		switch direction {
		case 'L':
			currentState.direction = leftTurns[currentState.direction]
		case 'R':
			currentState.direction = rightTurns[currentState.direction]
		}
	}
	return currentState
}

func parseInput(lines []string) []instruction {
	instructions := make([]instruction, 0, len(lines))
	expr := regexp.MustCompile("([A-Z])([0-9]+)")
	for _, line := range lines {
		matches := expr.FindStringSubmatch(line)
		op := rune(matches[1][0])
		n, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{operation: op, numerator: n})
	}
	return instructions
}
