package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day10 struct{}

type instruction struct {
	op        string
	numerator int
	cycles    int
}

func (p Day10) PartA(lines []string) any {
	instructions := parseInput10(lines)
	registers := make(map[string]int)
	registers["X"] = 1
	targets := []int{20, 60, 100, 140, 180, 220}
	toAdd := make([]int, 0, len(targets))
	cycle := 0
	for _, instr := range instructions {
		instructionCycle := getCycles(instr)
		for i := 0; i < instructionCycle; i++ {
			cycle++
			if common.Contains(targets, cycle) {
				fmt.Println("Cycle", cycle, "Register value: ", registers["X"])
				toAdd = append(toAdd, registers["X"]*cycle)
			}
		}
		doInstruction(instr, &registers)
	}
	return common.Sum(toAdd)
}

func (p Day10) PartB(lines []string) any {
	instructions := parseInput10(lines)
	registers := make(map[string]int)
	registers["X"] = 1
	buffer := make([][]rune, 6)
	cycle := 0
	for _, instr := range instructions {
		instructionCycle := getCycles(instr)
		for i := 0; i < instructionCycle; i++ {
			cycle++
			writeCrtBuffer(cycle, registers["X"], &buffer)
		}
		doInstruction(instr, &registers)
	}
	printCrtBuffer(buffer)
	return 0
}

func getCycles(instr instruction) int {
	return map[string]int{
		"noop": 1,
		"addx": 2,
	}[instr.op]
}

func doInstruction(instr instruction, registers *map[string]int) {
	switch instr.op {
	case "addx":
		(*registers)["X"] += instr.numerator
	}
}

func parseInput10(lines []string) []instruction {
	out := make([]instruction, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		instr := instruction{op: parts[0]}
		if len(parts) > 1 {
			x, _ := strconv.Atoi(parts[1])
			instr.numerator = x
		}
		out = append(out, instr)
	}
	return out
}

func writeCrtBuffer(cycle int, spritPos int, buffer *[][]rune) {
	// Find the x pos
	pos := (cycle - 1) % 40
	y := int(math.Floor(float64((cycle - 1) / 40)))
	if len((*buffer)[y]) == 0 {
		(*buffer)[y] = make([]rune, 40)
	}
	if pos >= spritPos-1 && pos <= spritPos+1 {
		(*buffer)[y][pos] = '#'
	} else {
		(*buffer)[y][pos] = ' '
	}
}

func printCrtBuffer(buffer [][]rune) {
	for _, line := range buffer {
		for _, col := range line {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
	fmt.Println()
}
