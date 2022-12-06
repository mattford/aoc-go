package year2020

import (
	"aocgen/pkg/common"
	"math"
	"strconv"
	"strings"
)

type Day13 struct{}

func (p Day13) PartA(lines []string) any {
	startTime, intervals := getInputParts13(lines, false)
	lowest := math.MaxInt
	lowestInterval := 0
	for _, interval := range intervals {
		x := int(math.Ceil(float64(startTime) / float64(interval)))
		y := x*interval - startTime
		if y < lowest {
			lowest = y
			lowestInterval = interval
		}
	}
	return lowest * lowestInterval
}

func (p Day13) PartB(lines []string) any {
	_, intervals := getInputParts13(lines, true)
	return findWhenBusesAlign(intervals)
}

func findWhenBusesAlign(intervals []int) int {
	first := intervals[0]

	timestamp := first
	period := first

	for i, interval := range intervals {
		if interval == 0 {
			continue
		}
		for (timestamp+i)%interval != 0 {
			timestamp += period
		}
		period = common.LCM(period, interval)
	}
	return timestamp
}

func getInputParts13(lines []string, withBlanks bool) (int, []int) {
	startTime, _ := strconv.Atoi(lines[0])
	parts := strings.Split(lines[1], ",")
	ints := make([]int, 0, len(parts)/2)
	for _, s := range parts {
		if s != "x" {
			n, _ := strconv.Atoi(s)
			ints = append(ints, n)
		} else if withBlanks {
			ints = append(ints, 0)
		}
	}
	return startTime, ints
}
