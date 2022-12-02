package year2020

import (
	"strconv"
)

type Day01 struct{}

func (p Day01) PartA(lines []string) any {
	for i := 0; i < len(lines); i++ {
		x, err := strconv.Atoi(lines[i])
		if err != nil {
			continue
		}
		for j := 0; j < len(lines); j++ {
			y, err := strconv.Atoi(lines[j])
			if i == j || err != nil {
				continue
			}
			if x+y == 2020 {
				return x * y
			}
		}
	}
	panic("Answer not found :O")
}

func (p Day01) PartB(lines []string) any {
	for i := 0; i < len(lines); i++ {
		x, err := strconv.Atoi(lines[i])
		if err != nil {
			continue
		}
		for j := 0; j < len(lines); j++ {
			y, err := strconv.Atoi(lines[j])
			if i == j || err != nil || x+y >= 2020 {
				continue
			}
			for k := 0; k < len(lines); k++ {
				z, err := strconv.Atoi(lines[k])
				if i == j || i == z || j == z || err != nil {
					continue
				}
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}
	panic("Answer not found :O")
}
