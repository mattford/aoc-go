package year2020

import (
	"aocgen/pkg/common"
	"regexp"
	"strconv"
	"strings"
)

type Day16 struct{}

func (p Day16) PartA(lines []string) any {
	rules, _, otherTickets := parseTickets(lines)
	total := 0
	for _, ticket := range otherTickets {
		valid, invalidValue := validTicket(ticket, rules)
		if !valid {
			total += invalidValue
		}
	}
	return total
}

func (p Day16) PartB(lines []string) any {
	rules, myTicket, otherTickets := parseTickets(lines)
	validTickets := make([][]int, 0)
	for _, ticket := range otherTickets {
		valid, _ := validTicket(ticket, rules)
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}
	ruleList := common.Keys(rules)
	validRules := make([][]string, len(ruleList))
	for i := 0; i < len(ruleList); i++ {
		validRules[i] = make([]string, 0)
		col := common.Column(validTickets, i)

		for ruleName, ranges := range rules {
			validForRule := true
		valueLoop:
			for _, ticketValue := range col {
				for _, rng := range ranges {
					if ticketValue >= rng[0] && ticketValue <= rng[1] {
						continue valueLoop
					}
				}
				validForRule = false
				break
			}
			if validForRule {
				validRules[i] = append(validRules[i], ruleName)
			}
		}
	}
	used := make([]string, len(ruleList))
	found := 0
	for found < len(used) {
		for idx, validRule := range validRules {
			for _, x := range used {
				validRule = common.Remove(validRule, x)
			}
			if len(validRule) == 1 {
				used[idx] = validRule[0]
				found++
			}
		}
	}
	total := 1
	for i, x := range used {
		if strings.Index(x, "departure") == 0 {
			total *= myTicket[i]
		}
	}
	return total
}

func validTicket(ticket []int, rules map[string][][]int) (bool, int) {
	invalidValue := 0
	valid := true
	for _, ticketValue := range ticket {
		validThisValue := false
	ruleLoop:
		for _, ranges := range rules {
			for _, rng := range ranges {
				if ticketValue >= rng[0] && ticketValue <= rng[1] {
					validThisValue = true
					continue ruleLoop
				}
			}
		}
		if !validThisValue {
			valid = false
			invalidValue = ticketValue
		}
	}
	return valid, invalidValue
}

func parseTickets(lines []string) (map[string][][]int, []int, [][]int) {
	rules := make(map[string][][]int)
	var myTicket []int
	otherTickets := make([][]int, 0)

	mode := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		if line == "your ticket:" {
			mode = 1
			continue
		} else if line == "nearby tickets:" {
			mode = 2
			continue
		}
		switch mode {
		case 0:
			// rule
			parseRule(line, &rules)
		case 1:
			// my ticket
			myTicket = parseTicket(line)
		case 2:
			// other tickets
			otherTickets = append(otherTickets, parseTicket(line))
		}
	}
	return rules, myTicket, otherTickets
}

func parseRule(line string, rules *map[string][][]int) {
	expr := regexp.MustCompile("([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)")
	matches := expr.FindStringSubmatch(line)
	class := matches[1]
	a, _ := strconv.Atoi(matches[2])
	b, _ := strconv.Atoi(matches[3])
	c, _ := strconv.Atoi(matches[4])
	d, _ := strconv.Atoi(matches[5])
	(*rules)[class] = [][]int{
		{a, b},
		{c, d},
	}
}

func parseTicket(line string) []int {
	strs := strings.Split(line, ",")
	ints := make([]int, 0, len(strs))
	for _, s := range strs {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}
