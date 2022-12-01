package year2022

import (
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
	"strings"
)

type Day01 struct{}

func (p Day01) PartA(lines []string) any {
	elves := getElves(lines)
	_, max := findMinMax(elves)
	return max
}

func (p Day01) PartB(lines []string) any {
	elves := getElves(lines)
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	topThree := elves[0:3]
	var sum int
	for _, i := range topThree {
		sum += i
	}
	return sum
}

func getElves(lines []string) []int {
	elves := make([]int, 0)
	var calorieCount int
	for _, line := range lines {
		if line == "\r" {
			elves = append(elves, calorieCount)
			calorieCount = 0
		} else {

			count, err := strconv.Atoi(strings.Replace(line, "\r", "", 1))
			if err != nil {
				logrus.Error(err)
				continue
			}
			calorieCount += count
		}
	}
	if calorieCount > 0 {
		elves = append(elves, calorieCount)
	}
	return elves
}

func findMinMax(elves []int) (min int, max int) {
	for _, elf := range elves {
		if elf < min {
			min = elf
		}
		if elf > max {
			max = elf
		}
	}
	return min, max
}
