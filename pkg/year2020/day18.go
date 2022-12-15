package year2020

import (
	"aocgen/pkg/common"
	"strconv"
	"strings"
)

type Day18 struct{}

func (p Day18) PartA(lines []string) any {
	return total(lines, true)
}

func (p Day18) PartB(lines []string) any {
	return total(lines, false)
}

func total(lines []string, unorderedOps bool) int {
	total := 0
	for _, line := range lines {
		x, _ := resolveExpression(line, unorderedOps)
		total += x
	}
	return total
}

func resolveExpression(expression string, unorderedOps bool) (int, int) {
	idx := 0
	chars := strings.Split(strings.ReplaceAll(expression, " ", ""), "")
	stack := make([]any, 0)
	for len(chars) > idx {
		char := chars[idx]
		switch char {
		case "+":
			stack = append(stack, "+")
		case "*":
			stack = append(stack, "*")
		case "(":
			// enter expression
			subTotal, innerIdx := resolveExpression(strings.Join(chars[idx+1:], ""), unorderedOps)
			stack = append(stack, subTotal)
			idx += innerIdx + 1
		case ")":
			// exit expression
			return calculateStack(stack, unorderedOps), idx
		default:
			x, _ := strconv.Atoi(char)
			stack = append(stack, x)
		}
		idx++
	}
	return calculateStack(stack, unorderedOps), idx
}

func calculateStack(expr []any, unorderedOps bool) int {
	ops := []string{"+", "*"}
	for _, op := range ops {
		stack := make([]any, 0)
		for i := 0; i < len(expr); i++ {
			v, _ := expr[i].(string)
			_, isInt := expr[i].(int)
			if isInt || (!unorderedOps && v != op) {
				stack = append(stack, expr[i])
				continue
			}
			switch v {
			case "*":
				l, _ := common.Pop(&stack).(int)
				r, _ := expr[i+1].(int)
				stack = append(stack, l*r)
				i++
			case "+":
				l, _ := common.Pop(&stack).(int)
				r, _ := expr[i+1].(int)
				stack = append(stack, l+r)
				i++
			default:
				stack = append(stack, v)
			}
		}
		expr = stack
	}
	x, _ := expr[0].(int)
	return x
}
