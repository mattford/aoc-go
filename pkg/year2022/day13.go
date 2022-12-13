package year2022

import (
	"aocgen/pkg/common"
	"fmt"
	"sort"
	"strconv"
)

type Day13 struct{}

func (p Day13) PartA(lines []string) any {
	pairs := parseInput13(lines)
	valid := 0
	for i, pair := range pairs {
		if sub := less(pair[0], pair[1]); sub == 0 || sub == 1 {
			valid += i + 1
		}
	}
	return valid
}

func (p Day13) PartB(lines []string) any {
	lines = append(lines, []string{"[[2]]", "[[6]]"}...)
	pairs := parseInput13(lines)
	allPackets := make([]any, 0, len(pairs)*2)
	for _, pair := range pairs {
		allPackets = append(allPackets, pair...)
	}
	sort.Slice(allPackets, func(i int, j int) bool {
		a := allPackets[i]
		b := allPackets[j]
		return less(a, b) == 1
	})
	out := 1
	for i, packet := range allPackets {
		asString := fmt.Sprintf("%v", packet)
		if asString == "[[[2]]]" || asString == "[[[6]]]" {
			out *= i + 1
		}
	}
	return out
}

func parseInput13(lines []string) [][]any {
	pairs := make([][]any, 0)
	pair := make([]any, 0, 2)
	for _, line := range lines {
		if line == "" {
			continue
		}
		expr, _ := parseExpression(line)
		pair = append(pair, expr)
		if len(pair) == 2 {
			pairs = append(pairs, pair)
			pair = make([]any, 0, 2)
		}
	}
	return pairs
}

func less(a any, b any) int {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt {
		if aInt > bInt {
			return -1
		} else if bInt > aInt {
			return 1
		}
		return 0
	}
	aList, aIsList := a.([]any)
	bList, bIsList := b.([]any)
	if !aIsList {
		aList = []any{aInt}
	}
	if !bIsList {
		bList = []any{bInt}
	}
	max := common.MaxInt([]int{len(aList), len(bList)})
	for i := 0; i < max; i++ {
		if i >= len(aList) {
			return 1
		}
		if i >= len(bList) {
			return -1
		}
		if sub := less(aList[i], bList[i]); sub != 0 {
			return sub
		}
	}
	return 0
}

func parseExpression(expr string) (any, int) {
	chars := []rune(expr)
	idx := 0
	out := make([]any, 0)
	nChars := make([]rune, 0)
	for idx < len(chars) {
		char := chars[idx]
		switch char {
		case '[':
			x, i := parseExpression(string(chars[idx+1:]))
			out = append(out, x)
			idx += i + 1
		case ']':
			if len(nChars) > 0 {
				n, _ := strconv.Atoi(string(nChars))
				out = append(out, n)
				nChars = make([]rune, 0)
			}
			idx++
			return out, idx
		case ',':
			if len(nChars) > 0 {
				n, _ := strconv.Atoi(string(nChars))
				out = append(out, n)
				nChars = make([]rune, 0)
			}
			idx++
		default:
			nChars = append(nChars, char)
			idx++
		}
	}
	return out, idx
}
