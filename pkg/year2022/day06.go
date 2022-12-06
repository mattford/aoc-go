package year2022

import (
	"aocgen/pkg/common"
)

type Day06 struct{}

func (p Day06) PartA(lines []string) any {
	chars := []rune(lines[0])
	return findMarkerIndex(chars, 4)
}

func (p Day06) PartB(lines []string) any {
	chars := []rune(lines[0])
	return findMarkerIndex(chars, 14)
}

func findMarkerIndex(chars []rune, markerLength int) int {
	for i := markerLength - 1; i < len(chars); i++ {
		group := chars[i-(markerLength-1) : i+1]
		if isMarker(group) {
			return i + 1
		}
	}
	return 0
}

func isMarker(group []rune) bool {
	seen := make([]rune, 0, len(group))
	for _, r := range group {
		if common.Contains(seen, r) {
			return false
		}
		seen = append(seen, r)
	}
	return true
}
