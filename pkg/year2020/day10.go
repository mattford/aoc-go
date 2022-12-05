package year2020

import (
	"aocgen/pkg/common"
	"math"
	"sort"
)

type Day10 struct{}

func (p Day10) PartA(lines []string) any {
	joltages := common.GetInts(lines)
	deviceRating := common.MaxInt(joltages) + 3
	sort.Sort(sort.IntSlice(joltages))
	chain, _ := findChain(joltages, deviceRating, make([]int, 0))
	diffs := getDiffs(chain)
	return diffs[1] * diffs[3]
}

func (p Day10) PartB(lines []string) any {
	inputJolts := common.GetInts(lines)
	deviceRating := common.MaxInt(inputJolts) + 3
	joltages := []int{0}
	joltages = append(joltages, inputJolts...)
	joltages = append(joltages, deviceRating)
	sort.Sort(sort.IntSlice(joltages))
	return poss(deviceRating, joltages)
}

var cache = make(map[int]int)

func poss(item int, ints []int) int {
	total := 0
	if item == 0 {
		return 1
	}
	for i := item - 3; i < item; i++ {
		_, ok := cache[i]
		if !ok {
			cache[i] = 0
			if containsInt(ints, i) {
				cache[i] += poss(i, ints)
			}
		}
		total += cache[i]
	}
	return total
}

func getDiffs(chain []int) map[int]int {
	out := make(map[int]int)
	last := 0
	for _, n := range chain {
		diff := math.Abs(float64(n - last))
		out[int(diff)]++
		last = n
	}
	return out
}

func findChain(joltages []int, deviceRating int, chain []int) ([]int, bool) {
	// careful pooh, that's recursion
	for _, joltage := range joltages {
		lastJoltage := 0
		if len(chain) > 0 {
			lastJoltage = chain[len(chain)-1]
		}
		if lastJoltage < joltage-3 || lastJoltage > joltage {
			break
		}
		thisChain := make([]int, len(chain))
		copy(thisChain, chain)
		thisJoltages := common.Without(joltages, joltage)
		thisChain = append(thisChain, joltage)
		if len(thisJoltages) == 0 && joltage >= deviceRating-3 && joltage <= deviceRating {
			thisChain = append(thisChain, deviceRating)
			return thisChain, true
		}
		subChain, valid := findChain(thisJoltages, deviceRating, thisChain)
		if valid {
			return subChain, true
		}
	}
	return chain, false
}
