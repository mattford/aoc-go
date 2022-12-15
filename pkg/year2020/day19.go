package year2020

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day19 struct{}

func (p Day19) PartA(lines []string) any {
	rules, messages := buildRules19(lines)
	expr := regexp.MustCompile("^" + buildRegex(0, rules) + "$")
	total := 0
	for _, message := range messages {
		if expr.Match([]byte(message)) {
			total++
		}
	}
	return total
}

func (p Day19) PartB(lines []string) any {
	rules, messages := buildRules19(lines)
	rules[8] = [][]any{{"42"}, {"42", "8"}}
	rules[11] = [][]any{{"42", "31"}, {"42", "11", "31"}}
	fmt.Println(buildRegex(11, rules))
	expr := regexp.MustCompile("^" + buildRegex(0, rules) + "$")
	total := 0
	for _, message := range messages {
		if expr.Match([]byte(message)) {
			total++
		}
	}
	return total
}

func buildRegex(idx int, rules map[int][][]any) string {
	rule := rules[idx]
	groups := make(map[int][]string)
	repeatingGroup := 100
	repeatingWithinGroup := 100
	for id, group := range rule {
		for xxx, ruleIdx := range group {
			rIdx, err := strconv.Atoi(ruleIdx.(string))
			if err == nil {
				if _, ok := groups[id]; !ok {
					groups[id] = make([]string, 0, 1)
				}
				if rIdx != idx {
					groups[id] = append(groups[id], buildRegex(rIdx, rules))
				} else {
					repeatingGroup = id
					repeatingWithinGroup = xxx
				}
			} else {
				str, _ := ruleIdx.(string)
				return str
			}
		}
	}
	possibles := make([]string, 0)
	for gid, group := range groups {
		thisOne := ""
		for sid, str := range group {
			if sid == repeatingWithinGroup && gid == repeatingGroup {
				str = "(" + str + ")+"
			}
			thisOne += str
		}
		possibles = append(possibles, thisOne)
	}
	pattern := "(" + strings.Join(possibles, "|") + ")"
	return pattern
}

func buildRules19(lines []string) (map[int][][]any, []string) {
	rules := make(map[int][][]any)
	var messages []string
	for i, line := range lines {
		if line == "" {
			messages = lines[i+1:]
			break
		}
		parts := strings.Split(line, ":")
		idx, _ := strconv.Atoi(parts[0])
		subLists := strings.Split(strings.Trim(parts[1], " "), "|")
		thisRule := make([][]any, 0)
		for _, subList := range subLists {
			expr := regexp.MustCompile("\"([a-z])\"")
			matches := expr.FindStringSubmatch(subList)
			if len(matches) > 0 {
				thisRule = append(thisRule, []any{matches[1]})
			} else {
				subRules := make([]any, 0)
				for _, x := range strings.Split(strings.Trim(subList, " "), " ") {
					subRules = append(subRules, x)
				}
				thisRule = append(thisRule, subRules)
			}
		}
		rules[idx] = thisRule
	}
	return rules, messages
}
