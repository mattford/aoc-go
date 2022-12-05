package year2022

import (
	"regexp"
	"strconv"
	"strings"
)

type Day05 struct{}

func (p Day05) PartA(lines []string) any {
	stacks, instructions := splitInput(lines)
	return runInstructions(stacks, instructions, false)
}

func (p Day05) PartB(lines []string) any {
	stacks, instructions := splitInput(lines)
	return runInstructions(stacks, instructions, true)
}

func runInstructions(stacks [][]string, instructions []string, multi bool) string {
	for _, instr := range instructions {
		expr := regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
		matches := expr.FindStringSubmatch(instr)
		n, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		src, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		dst, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}
		stacks = moveCrates(stacks, n, src, dst, multi)
	}
	return getResult(stacks)
}

func getResult(stacks [][]string) string {
	out := make([]string, len(stacks))
	for i, stack := range stacks {
		out[i] = stack[0]
	}
	return strings.Join(out, "")
}

func moveCrates(stacks [][]string, n int, src int, dst int, multi bool) [][]string {
	// First, sort out the src
	movedCrates := stacks[src-1][:n]
	if multi {
		movedCrates = reverse(movedCrates)
	}
	stacks[src-1] = stacks[src-1][n:]
	stacks[dst-1] = prepend(stacks[dst-1], movedCrates)
	return stacks
}

func prepend(arr []string, items []string) []string {
	for _, item := range items {
		arr = append([]string{item}, arr...)
	}
	return arr
}

func reverse(arr []string) []string {
	out := make([]string, len(arr))
	j := 0
	for i := len(arr) - 1; i >= 0; i-- {
		out[j] = arr[i]
		j++
	}
	return out
}

func splitInput(lines []string) ([][]string, []string) {
	stacks := make([][]string, 0)
	var instructionsIdx int
outer:
	for idx, line := range lines {
		if line == "" {
			instructionsIdx = idx + 1
			break
		}
		chars := []rune(line)
		charIdx := 1
		stackIdx := 0
		for charIdx < len(chars) {
			if len(stacks) <= stackIdx {
				stacks = append(stacks, make([]string, 0))
			}
			char := string(chars[charIdx])
			if char == "1" {
				continue outer
			}
			if char != " " {
				stacks[stackIdx] = append(stacks[stackIdx], char)
			}
			charIdx += 4
			stackIdx++
		}
	}
	instructions := lines[instructionsIdx:]
	return stacks, instructions
}
