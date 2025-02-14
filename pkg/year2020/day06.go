package year2020

import "strings"

type Day06 struct{}

type Group struct {
	members   int
	questions map[string]int
}

func (p Day06) PartA(lines []string) any {
	groups := getGroups(lines)
	count := 0
	for _, group := range groups {
		count += len(group.questions)
	}
	return count
}

func (p Day06) PartB(lines []string) any {
	groups := getGroups(lines)
	count := 0
	for _, group := range groups {
		for _, question := range group.questions {
			if question == group.members {
				count++
			}
		}
	}
	return count
}

func newGroup() Group {
	return Group{
		members:   0,
		questions: make(map[string]int),
	}
}

func getGroups(lines []string) (groups []Group) {
	group := newGroup()
	for _, line := range lines {
		if line == "" {
			groups = append(groups, group)
			group = newGroup()
			continue
		}
		group.members++
		questions := strings.Split(line, "")
		for _, question := range questions {
			group.questions[question]++
		}
	}
	return append(groups, group)
}
