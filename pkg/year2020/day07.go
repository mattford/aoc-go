package year2020

import (
	"regexp"
	"strconv"
	"strings"
)

type Day07 struct{}

type Bag map[string]int

func (p Day07) PartA(lines []string) any {
	bags := getBags(lines)
	myBag := "shiny gold bag"
	canHold := getCanHold(bags, myBag)
	return len(unique(canHold))
}

func (p Day07) PartB(lines []string) any {
	bags := getBags(lines)
	return countBags(bags, "shiny gold bag")
}

func countBags(bags map[string]Bag, myBag string) int {
	count := 0
	for bag, cnt := range bags[myBag] {
		count += cnt
		count += cnt * countBags(bags, bag)
	}
	return count
}

func contains(a []string, y string) bool {
	for _, x := range a {
		if x == y {
			return true
		}
	}
	return false
}

func unique(a []string) []string {
	newA := make([]string, 0)
	for _, y := range a {
		if !contains(newA, y) {
			newA = append(newA, y)
		}
	}
	return newA
}

func getCanHold(bags map[string]Bag, myBag string) []string {
	canHold := make([]string, 0)
	for name, bag := range bags {
		if bag[myBag] > 0 {
			canHold = append(canHold, name)
			for _, b := range getCanHold(bags, name) {
				canHold = append(canHold, b)
			}
		}
	}
	return canHold
}

func getBags(lines []string) map[string]Bag {
	bags := make(map[string]Bag)
	expr := regexp.MustCompile("([0-9]+) ([a-z ]*)")
	for _, line := range lines {
		bits := strings.Split(line, " contain ")
		if bits[1] == "no other bags." {
			continue
		}
		bagName := strings.TrimRight(bits[0], " s")
		if bags[bagName] == nil {
			bags[bagName] = make(Bag)
		}
		containedBags := strings.Split(bits[1], ",")
		for _, containedBag := range containedBags {
			containedBag = strings.TrimRight(containedBag, " .s")

			matches := expr.FindStringSubmatch(containedBag)
			count, err := strconv.Atoi(matches[1])
			if err != nil {
				continue
			}
			bagType := strings.Trim(matches[2], " ")
			bags[bagName][bagType] = count
		}
	}
	return bags
}
