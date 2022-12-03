package year2022

import (
	"strings"
)

type Day03 struct{}

func (p Day03) PartA(lines []string) any {
	total := 0
	for _, line := range lines {
		letters := []rune(line)
		shared := findCommon(items(letters[:len(letters)/2]), items(letters[len(letters)/2:]))
		priority := getPriority(shared)
		total += priority
	}
	return total
}

func (p Day03) PartB(lines []string) any {
	group := make([]map[rune]int, 0, 3)
	total := 0
	for _, line := range lines {
		if len(group) >= 3 {
			common := findCommon(group[0], group[1], group[2])
			group = make([]map[rune]int, 0, 3)
			total += getPriority(common)
		}
		group = append(group, items([]rune(line)))
	}
	common := findCommon(group[0], group[1], group[2])
	total += getPriority(common)
	return total
}

func findCommon(a ...map[rune]int) rune {
main:
	for key := range a[0] {
		for _, b := range a[1:] {
			_, ok := b[key]
			if !ok {
				continue main
			}
		}
		return key
	}
	return 0
}

func items(letters []rune) (items map[rune]int) {
	items = make(map[rune]int)
	for _, letter := range letters {
		items[letter]++
	}
	return
}

func getPriority(item rune) int {
	priority := 0
	codePoint := int(item)
	if codePoint >= 65 && codePoint <= 90 {
		priority += 26
		priority += getPriority([]rune(strings.ToLower(string(item)))[0])
	} else {
		priority += codePoint - 96
	}
	return priority
}
