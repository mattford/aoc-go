package year2020

import (
	"regexp"
	"strconv"
	"strings"
)

type Day02 struct{}

func (p Day02) PartA(lines []string) any {
	var validCount int
	for _, line := range lines {
		min, max, targetLetter, letters, err := splitLine(line)
		if err != nil {
			panic(err)
		}

		letterCount := make(map[string]int)
		for idx := range letters {
			letter := letters[idx]
			letterCount[letter]++
		}

		valid := letterCount[targetLetter] >= min && letterCount[targetLetter] <= max
		if valid {
			validCount++
		}
	}
	return validCount
}

func (p Day02) PartB(lines []string) any {
	var validCount int
	for _, line := range lines {
		min, max, targetLetter, letters, err := splitLine(line)
		if err != nil {
			panic(err)
		}
		valid := (letters[min-1] == targetLetter || letters[max-1] == targetLetter) && (letters[min-1] != targetLetter || letters[max-1] != targetLetter)

		if valid {
			validCount++
		}
	}
	return validCount
}

func splitLine(line string) (min, max int, targetLetter string, letters []string, err any) {
	expr, err := regexp.Compile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")
	if err != nil {
		return
	}
	match := expr.FindStringSubmatch(line)
	min, err = strconv.Atoi(match[1])
	if err != nil {
		return
	}
	max, err = strconv.Atoi(match[2])
	if err != nil {
		return
	}
	targetLetter = match[3]
	str := match[4]
	letters = strings.Split(str, "")
	return
}
