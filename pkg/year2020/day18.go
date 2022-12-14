package year2020

import (
	"strconv"
	"strings"
)

type Day18 struct{}

func (p Day18) PartA(lines []string) any {
	return total(lines)
}

func (p Day18) PartB(lines []string) any {
	return "implement_me"
}

func total(lines []string) int {
	total := 0
	for _, line := range lines {
		x, _ := resolveExpression(line)
		total += x
	}
	return total
}

func resolveExpression(expression string) (int, int) {
	var operands []int
	var op string
	idx := 0
	chars := strings.Split(expression, "")
	for len(chars) > idx {
		char := chars[idx]
		switch char {
		case " ":
			idx++
			continue
		case "+":
			op = "+"
		case "*":
			op = "*"
		case "(":
			// enter expression
			subTotal, innerIdx := resolveExpression(strings.Join(chars[idx+1:], ""))
			operands = append(operands, subTotal)
			idx += innerIdx + 1
		case ")":
			// exit expression
			return operands[0], idx
		default:
			x, _ := strconv.Atoi(char)
			operands = append(operands, x)
		}
		idx++
		if len(operands) == 2 {
			result := 0
			switch op {
			case "+":
				result = operands[0] + operands[1]
			case "*":
				result = operands[0] * operands[1]
			}
			operands = make([]int, 1, 2)
			operands[0] = result
			op = ""
		}
	}
	return operands[0], idx
}
