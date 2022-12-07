package year2020

import (
	"aocgen/pkg/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day14 struct{}

type write struct {
	idx   int
	value int64
}

func (p Day14) PartA(lines []string) any {
	var mask map[int]int
	memory := make(map[int]int64)
	for _, line := range lines {
		if strings.Index(line, "mask") == 0 {
			mask = getMask(line)
		} else {
			write := parseWrite(line)
			memory = applyWrite(write, mask, memory)
		}
	}
	total := int64(0)
	for _, x := range memory {
		total += x
	}
	return total
}

func (p Day14) PartB(lines []string) any {
	var mask map[int]int
	memory := make(map[int64]int64)
	for _, line := range lines {
		if strings.Index(line, "mask") == 0 {
			mask = getMask(line)
		} else {
			write := parseWrite(line)
			for _, idx := range decodeMemoryAddress(write.idx, mask) {
				memory[idx] = write.value
			}
		}
	}
	total := int64(0)
	for _, x := range memory {
		total += x
	}
	return total
}

func decodeMemoryAddress(addr int, mask map[int]int) []int64 {
	newVal := int64(addr)

	floatings := make([]int64, 0)
	for idx, val := range mask {
		m := int64(math.Pow(2, float64(idx)))
		hasVal := m&newVal == m
		if val == 1 && !hasVal {
			newVal += m
		} else if val == -1 {
			floatings = append(floatings, m)
		}
	}

	if len(floatings) > 0 {
		return generatePerms(newVal, floatings)
	}

	return []int64{newVal}
}

func generatePerms(base int64, floaties []int64) []int64 {
	outs := []int64{base}
	for i, m := range floaties {
		hasVal := m&base == m
		var opp int64
		if hasVal {
			opp = base - m
		} else {
			opp = base + m
		}
		outs = append(outs, opp)
		outs = append(outs, generatePerms(base, floaties[i+1:])...)
		outs = append(outs, generatePerms(opp, floaties[i+1:])...)
	}
	return common.Unique(outs)
}

func applyWrite(wr write, mask map[int]int, memory map[int]int64) map[int]int64 {
	newVal := wr.value
	for idx, val := range mask {
		m := int64(math.Pow(2, float64(idx)))
		hasVal := m&newVal == m
		if val == 1 && !hasVal {
			newVal += m
		} else if val == 0 && hasVal {
			newVal -= m
		}
	}
	memory[wr.idx] = newVal
	return memory
}

func getMask(line string) map[int]int {
	var mask string
	_, _ = fmt.Sscanf(line, "mask = %s", &mask)
	bitMask := make(map[int]int)
	for i, char := range []rune(mask) {
		if char == 'X' {
			bitMask[len(mask)-i-1] = -1
			continue
		}
		v, _ := strconv.Atoi(string(char))
		bitMask[len(mask)-i-1] = v
	}
	return bitMask
}

func parseWrite(line string) write {
	var idx int
	var value int64
	_, _ = fmt.Sscanf(line, "mem[%d] = %d", &idx, &value)
	return write{
		idx:   idx,
		value: value,
	}
}
