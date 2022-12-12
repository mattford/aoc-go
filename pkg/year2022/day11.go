package year2022

import (
	"aocgen/pkg/common"
	"math"
	"regexp"
	"sort"
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
	for i := 0; i < 20; i++ {
		monkeys = doRound(monkeys, false, 0)
	}
	sort.Slice(monkeys, func(i, j int) bool {
		a := monkeys[i]
		b := monkeys[j]
		return a.inspections > b.inspections
	})
	return monkeys[0].inspections * monkeys[1].inspections
}

func (p Day11) PartB(lines []string) any {
	monkeys := parseInput11(lines)
	mod := getModulo(monkeys)
	for i := 0; i < 10000; i++ {
		monkeys = doRound(monkeys, true, mod)
	}
	sort.Slice(monkeys, func(i, j int) bool {
		a := monkeys[i]
		b := monkeys[j]
		return a.inspections > b.inspections
	})
	return monkeys[0].inspections * monkeys[1].inspections
}

func doRound(monkeys []monkey, extremelyWorried bool, mod int) []monkey {
	for idx, monk := range monkeys {
		for _, item := range monk.items {
			monk.inspections++
			item = doOp(monk, item)
			if !extremelyWorried {
				item = int(math.Floor(float64(item / 3)))
			} else {
				item = item % mod
			}
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

func getModulo(monkeys []monkey) int {
	total := 1
	for _, m := range monkeys {
		total *= m.test
	}
	return total
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

func parseInput11(lines []string) []monkey {
	thisMonkey := monkey{}
	monkeys := make([]monkey, 0)
	for _, line := range lines {
		if line == "" {
			monkeys = append(monkeys, thisMonkey)
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
	monkeys = append(monkeys, thisMonkey)
	return monkeys
}
