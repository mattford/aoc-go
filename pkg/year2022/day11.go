package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Day11 struct{}

type monkey struct {
	items       []int
	op          string
	test        int
	ifTrue      int
	ifFalse     int
	inspections int
}

func (p Day11) PartA(lines []string) any {
	monkeys := parseInput11(lines)
	for i := 0; i < 5; i++ {
		monkeys = doRound(monkeys)
		fmt.Println(monkeys)
	}
	fmt.Println(monkeys)
	return "implement_me"
}

func (p Day11) PartB(lines []string) any {
	return "implement_me"
}

func doRound(monkeys map[int]monkey) map[int]monkey {
	for idx, monk := range monkeys {

		for _, item := range monk.items {
			monk.inspections++
			item = doOp(monk, item)
			item = int(math.Floor(float64(item / 3)))
			if item%monk.test == 0 {
				otherMonkey := monkeys[monk.ifTrue]
				otherMonkey.items = append(otherMonkey.items, item)
				monkeys[monk.ifTrue] = otherMonkey
			} else {
				otherMonkey := monkeys[monk.ifFalse]
				otherMonkey.items = append(otherMonkey.items, item)
				monkeys[monk.ifFalse] = otherMonkey
			}
		}
		monk.items = make([]int, 0)
		monkeys[idx] = monk
	}
	return monkeys
}

func doOp(monk monkey, item int) int {
	op := monk.op
	expr := regexp.MustCompile("new = old ([*+]+) ([a-z0-9]+)")
	matches := expr.FindStringSubmatch(op)
	var operand int
	if matches[2] == "old" {
		operand = item
	} else {
		x, _ := strconv.Atoi(strings.Trim(matches[2], " "))
		operand = x
	}
	switch matches[1] {
	case "+":
		return item + operand
	case "*":
		return item * operand
	}
	return item
}

func parseInput11(lines []string) map[int]monkey {
	thisMonkey := monkey{}
	monkeys := make(map[int]monkey)
	monkeyIdx := 0
	for _, line := range lines {
		if line == "" {
			monkeys[monkeyIdx] = thisMonkey
			monkeyIdx++
			thisMonkey = monkey{}
		}
		parts := strings.Split(strings.Trim(line, " "), ":")
		switch parts[0] {
		case "Starting items":
			thisMonkey.items = common.GetInts(strings.Split(strings.Trim(parts[1], " "), ","))
		case "Operation":
			thisMonkey.op = parts[1]
		case "Test":
			split := strings.Split(parts[1], " by ")
			x, _ := strconv.Atoi(strings.Trim(split[1], " "))
			thisMonkey.test = x
		case "If true":
			split := strings.Split(parts[1], " monkey ")
			x, _ := strconv.Atoi(strings.Trim(split[1], " "))
			thisMonkey.ifTrue = x
		case "If false":
			split := strings.Split(parts[1], " monkey ")
			x, _ := strconv.Atoi(strings.Trim(split[1], " "))
			thisMonkey.ifFalse = x
		}
	}
	monkeys[monkeyIdx] = thisMonkey
	return monkeys
}
