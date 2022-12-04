package year2020

import (
	"regexp"
	"strconv"
	"strings"
)

type Day08 struct{}

type State struct {
	pointer      int
	instructions []string
	accumulator  int
	alreadyRun   []int
}

func (state *State) step() bool {
	if containsInt(state.alreadyRun, state.pointer) {
		return false
	}
	state.alreadyRun = append(state.alreadyRun, state.pointer)
	instruction := string(state.instructions[state.pointer])
	splits := strings.Split(instruction, " ")
	instr, param := splits[0], splits[1]
	expr := regexp.MustCompile("([+-])([0-9]+)")
	matches := expr.FindStringSubmatch(param)
	operator := matches[1]
	numerator, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	switch instr {
	case "acc":
		switch operator {
		case "+":
			state.accumulator += numerator
		case "-":
			state.accumulator -= numerator
		}
		state.pointer++
	case "jmp":
		switch operator {
		case "+":
			state.pointer += numerator
		case "-":
			state.pointer -= numerator
		}
	default:
		state.pointer++
	}
	return state.pointer >= 0 && state.pointer < len(state.instructions)
}

func (state *State) run() int {
	for {
		out := state.step()
		if !out {
			return state.accumulator
		}
	}
}

func (p Day08) PartA(lines []string) any {
	state := &State{
		instructions: lines,
		pointer:      0,
		accumulator:  0,
		alreadyRun:   make([]int, 0, len(lines)),
	}
	return state.run()
}

func (p Day08) PartB(lines []string) any {
	perms := genPermutations(lines)
	for _, perm := range perms {
		state := &State{
			instructions: perm,
			pointer:      0,
			accumulator:  0,
			alreadyRun:   make([]int, 0, len(perm)),
		}
		out := state.run()
		if state.pointer == len(perm) {
			return out
		}
	}
	return "not found"
}

func genPermutations(lines []string) [][]string {
	permutations := make([][]string, 0, len(lines)*2)
	for idx, line := range lines {
		if strings.Index(line, "jmp") == 0 {
			newLines := make([]string, len(lines))
			copy(newLines, lines)
			newLines[idx] = strings.Replace(line, "jmp", "nop", 1)
			permutations = append(permutations, newLines)
		} else if strings.Index(line, "nop") == 0 {
			newLines := make([]string, len(lines))
			copy(newLines, lines)
			newLines[idx] = strings.Replace(line, "nop", "jmp", 1)
			permutations = append(permutations, newLines)
		}
	}
	return permutations
}

func containsInt(haystack []int, needle int) bool {
	for _, i := range haystack {
		if i == needle {
			return true
		}
	}
	return false
}
