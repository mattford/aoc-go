package year2022

import (
	"regexp"
	"strconv"
)

type Day21 struct{}

type monkey21 struct {
	value    float64
	op       string
	left     string
	right    string
	resolved bool
}

func (p Day21) PartA(lines []string) any {
	monkeys := parseInput21(lines)
	for !monkeys["root"].resolved {
		monkeys = resolveMonkeys(monkeys)
	}

	return int(monkeys["root"].value)
}

func (p Day21) PartB(lines []string) any {
	monkeys := parseInput21(lines)
	l, r, monkeys := attempt(monkeys, 0)
	var other string
	var otherValue float64
	if l == 0 {
		other = monkeys["root"].left
		otherValue = r
	} else if r == 0 {
		other = monkeys["root"].right
		otherValue = l
	}
	return int(traceBack(monkeys, other, otherValue))
}

func traceBack(monkeys map[string]monkey21, m string, target float64) float64 {
	monk := monkeys[m]
	if monk.left == "" && monk.right == "" {
		return target
	}
	leftM := monkeys[monk.left]
	rightM := monkeys[monk.right]
	var prevMonkey string
	var newTarget, otherValue float64
	isLeft := false
	if leftM.value == 0 {
		prevMonkey = monk.left
		otherValue = rightM.value
		isLeft = true
	} else if rightM.value == 0 {
		prevMonkey = monk.right
		otherValue = leftM.value
	}
	switch monk.op {
	case "+":
		newTarget = target - otherValue
	case "-":
		if isLeft {
			newTarget = target + otherValue
		} else {
			newTarget = otherValue - target
		}
	case "*":
		newTarget = target / otherValue
	case "/":
		if isLeft {
			newTarget = target * otherValue
		} else {
			newTarget = otherValue / target
		}
	}
	return traceBack(monkeys, prevMonkey, newTarget)
}

func attempt(monkeys map[string]monkey21, myNumber float64) (float64, float64, map[string]monkey21) {
	monkeys["humn"] = monkey21{value: myNumber, resolved: true}
	left := monkeys["root"].left
	right := monkeys["root"].right
	for !monkeys["root"].resolved {
		monkeys = resolveMonkeys(monkeys)
	}
	return monkeys[left].value, monkeys[right].value, monkeys
}

func resolveMonkeys(monkeys map[string]monkey21) map[string]monkey21 {
	for i, m := range monkeys {
		if !m.resolved {
			leftM := monkeys[m.left]
			rightM := monkeys[m.right]
			if leftM.resolved && rightM.resolved {
				m.resolved = true
				switch m.op {
				case "+":
					m.value = leftM.value + rightM.value
				case "-":
					m.value = leftM.value - rightM.value
				case "*":
					m.value = leftM.value * rightM.value
				case "/":
					m.value = leftM.value / rightM.value
				}
				if leftM.value == 0 || rightM.value == 0 {
					m.value = 0
				}
			}
			monkeys[i] = m
		}
	}
	return monkeys
}

func parseInput21(lines []string) map[string]monkey21 {
	monkeys := make(map[string]monkey21)
	expr1 := regexp.MustCompile("([a-z]+): ([0-9]+)")
	expr2 := regexp.MustCompile("([a-z]+): ([a-z]+) ([-+*/]) ([a-z]+)")
	for _, line := range lines {
		if expr1.Match([]byte(line)) {
			matches := expr1.FindStringSubmatch(line)
			i, _ := strconv.Atoi(matches[2])
			monkeys[matches[1]] = monkey21{
				value:    float64(i),
				resolved: true,
			}
		} else if expr2.Match([]byte(line)) {
			matches := expr2.FindStringSubmatch(line)
			monkeys[matches[1]] = monkey21{
				value: -1,
				left:  matches[2],
				op:    matches[3],
				right: matches[4],
			}
		}
	}
	return monkeys
}
